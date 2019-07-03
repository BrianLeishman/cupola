package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	humanize "github.com/dustin/go-humanize"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var input string
	flag.StringVar(&input, "i", "", "Set an input PST file location")
	flag.Parse()

	f, err := os.Open(input)
	check(err)
	defer f.Close()

	var header pstHeader
	check(binary.Read(f, binary.LittleEndian, &header))

	if header.Signature != pstSignature {
		panic("PST file isn't valid (signature is wrong)")
	}

	if header.IndexType >= 0x10 {
		fmt.Println("Type: 64 bit")
		fmt.Printf("Index Type: 0x%02X\n", header.IndexType)

		var h pst64Header
		check(binary.Read(f, binary.LittleEndian, &h))

		fmt.Println("Size:", humanize.Bytes(h.TotalFileSize))
		fmt.Printf("Encryption: ")
		switch h.EncryptionType {
		case 0x00:
			fmt.Println("No encryption")
		case 0x01:
			fmt.Println("Compressible")
		case 0x02:
			fmt.Println("Strong")
			panic("strong encryption is not supported")
		default:
			panic(fmt.Errorf("unknown encrption type 0x%02X", h.EncryptionType))
		}

		// pst64TraverseIndex1(f, h.OffsetIndex1)

		nodes := 0
		leafs := 0

		var spiderIndex2Node func(offset uint64, backPointer uint64)
		spiderIndex2Node = func(offset uint64, backPointer uint64) {
			_, err := f.Seek(int64(offset), 0)
			check(err)

			var node pst64IndexNode
			check(binary.Read(f, binary.LittleEndian, &node))

			if node.BackPointer != backPointer {
				panic("backPointers don't match, this PST might be corrupt")
			}

			itemsReader := bytes.NewReader(node.Items[:])

			if node.NodeLevel != 0 {
				nodes++

				items := make([]pst64Index2NodeItem, node.ItemCount)
				check(binary.Read(itemsReader, binary.LittleEndian, &items))

				for i := uint8(0); i < node.ItemCount; i++ {
					spiderIndex2Node(items[i].Offset, items[i].BackPointer)
				}
			} else {
				leafs++

				items := make([]pst64Index2LeafNodeItem, node.ItemCount)
				check(binary.Read(itemsReader, binary.LittleEndian, &items))

				for _, item := range items {
					item.DescIID = clearBitUint64(item.DescIID, 0)

					if item.TreeIID != 0 {
						continue
					}
					if item.DescIID != 0 {
						offset, size, err := pst64SearchIndex1(f, h.OffsetIndex1, item.DescIID)
						// fmt.Println("Found offset:", offset, "size:", size)
						check(err)

						_, err = f.Seek(int64(offset), 0)
						check(err)

						b := make([]byte, size)
						_, err = io.ReadFull(f, b)
						check(err)

						decrypt(b)

						var desc pstAssociatedDescriptorItem
						r := bytes.NewReader(b)
						check(binary.Read(r, binary.LittleEndian, &desc))

						// fmt.Printf("Signature: 0x%04x, IndexOffset: 0x%04x, B5Offset: 0x%04x\n", desc.Signature, desc.IndexOffset, desc.B5Offset)

						switch desc.Signature {
						case 0xbcec:
							_, err = r.Seek(int64(desc.IndexOffset), 0)
							check(err)

							// spew.Dump(b)

							internal := desc.B5Offset|0x0f == desc.B5Offset
							desc.B5Offset >>= 4
							if internal {
								log.Println("low order bits are set!!! the data is somewhere else!!1")
							} else {
								// log.Println("low order bits are NOT set! thakn god lol")

								_, err = r.Seek(int64(uint32(desc.IndexOffset)+desc.B5Offset+2), 0)
								check(err)

								var b5Pair pstDataIndex1Pair
								check(binary.Read(r, binary.LittleEndian, &b5Pair))
								// spew.Dump(b5Pair)

								_, err = r.Seek(int64(b5Pair.Start), 0)
								check(err)

								var b5 pstDataIndex1B5
								check(binary.Read(r, binary.LittleEndian, &b5))
								// spew.Dump(b5)

								b5.DescOffset >>= 4
								_, err = r.Seek(int64(uint32(desc.IndexOffset)+b5.DescOffset+2), 0)
								check(err)

								var descPair pstDataIndex1Pair
								check(binary.Read(r, binary.LittleEndian, &descPair))
								// spew.Dump(descPair)

								_, err = r.Seek(int64(descPair.Start), 0)
								check(err)

								switch b5.DataSize {
								case 6:
									mapiEntries := make([]pstDataIndex1MAPI, (descPair.End-descPair.Start)/8)
									check(binary.Read(r, binary.LittleEndian, mapiEntries))
									for _, m := range mapiEntries {
										fmt.Println(m)
										v, err := getMapiItemValue(desc, r, m)
										check(err)
										fmt.Println("value:", v)
									}
								default:
									panic(fmt.Errorf("unhandled b5.DataSize of 0x%x for block with I_ID of 0x%x", b5.DataSize, item.DescIID))
								}
							}
							// spew.Dump(items)
						default:
							panic(fmt.Errorf("unhandled associated descriptor item with signature 0x%x for block with I_ID of 0x%x", desc.Signature, item.DescIID))
						}

						os.Exit(0)
					}
				}
				os.Exit(0)
			}
		}

		spiderIndex2Node(h.OffsetIndex2, h.BackPointer2)

		fmt.Println("Found", nodes, "outer nodes in index2")
		fmt.Println("Found", leafs, "leaf nodes in index2")

	} else {
		panic("32 bit PST files aren't supported (yet)")
	}
}
