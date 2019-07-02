package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
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
					// if items[i].TreeIID != 0 {
					// 	// spew.Dump(search(f, h.OffsetIndex1, items[i].TreeIID))
					// 	os.Exit(0)
					// }
					if item.DescIID != 0 {
						fmt.Println("DescIID:", item.DescIID)

						stack := make([]uint64, 0)
						spew.Dump(pst64SearchIndex1(f, h.OffsetIndex1, item.DescIID, stack))

						// _, err := f.Seek(int64(item.DescIID), 0)
						// check(err)

						// var desc pstAssociatedDescriptorItem
						// check(binary.Read(f, binary.LittleEndian, &desc))

						// spew.Dump(desc)

						os.Exit(0)
					}
				}
			}
		}

		spiderIndex2Node(h.OffsetIndex2, h.BackPointer2)

		fmt.Println("Found", nodes, "outer nodes in index2")
		fmt.Println("Found", leafs, "leaf nodes in index2")

	} else {
		panic("32 bit PST files aren't supported (yet)")
	}
}
