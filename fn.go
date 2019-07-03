package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func clearBitUint64(n uint64, pos uint) uint64 {
	mask := ^(uint64(1) << pos)
	n &= mask
	return n
}

func pst64TraverseIndex1(f *os.File, offset uint64) error {
	currOffset, err := f.Seek(0, 1)
	if err != nil {
		return err
	}
	defer f.Seek(currOffset, 0)

	_, err = f.Seek(int64(offset), 0)
	if err != nil {
		return err
	}

	var node pst64IndexNode
	if err := binary.Read(f, binary.LittleEndian, &node); err != nil {
		return err
	}

	leaf := node.NodeLevel == 0

	if !leaf {
		items := make([]pst64Index1NodeItem, node.ItemCount)
		if err := binary.Read(bytes.NewReader(node.Items[:]), binary.LittleEndian, &items); err != nil {
			return err
		}

		for _, item := range items {
			if err := pst64TraverseIndex1(f, item.Offset); err != nil {
				return err
			}
		}
	} else {
		items := make([]pst64Index1LeafNodeItem, node.ItemCount)
		if err := binary.Read(bytes.NewReader(node.Items[:]), binary.LittleEndian, &items); err != nil {
			return err
		}

		for _, item := range items {
			fmt.Println(item.IID)
		}
	}

	return nil
}

func pst64SearchIndex1(f *os.File, offset uint64, iid uint64) (dataOffset uint64, size uint16, err error) {
	currOffset, err := f.Seek(0, 1)
	if err != nil {
		return 0, 0, err
	}
	defer f.Seek(currOffset, 0)

	_, err = f.Seek(int64(offset), 0)
	if err != nil {
		return 0, 0, err
	}

	var node pst64IndexNode
	if err := binary.Read(f, binary.LittleEndian, &node); err != nil {
		return 0, 0, err
	}

	if node.NodeLevel == 0 {
		items := make([]pst64Index1LeafNodeItem, node.ItemCount)
		if err := binary.Read(bytes.NewReader(node.Items[:]), binary.LittleEndian, &items); err != nil {
			return 0, 0, err
		}

		for _, item := range items {
			// log.Println("considering I-ID", item.IID)
			if item.IID == iid {
				return item.Offset, item.Size, nil
			}
		}

		return 0, 0, nil
	}

	items := make([]pst64Index1NodeItem, node.ItemCount)
	if err := binary.Read(bytes.NewReader(node.Items[:]), binary.LittleEndian, &items); err != nil {
		return 0, 0, err
	}

	for i := node.ItemCount - 1; i >= 0; i-- {
		if items[i].IID <= iid {
			// log.Println("searching node", items[i].IID)
			return pst64SearchIndex1(f, items[i].Offset, iid)
		}
	}

	return 0, 0, nil
}

func getMapiItemValue(desc pstAssociatedDescriptorItem, r *bytes.Reader, m pstDataIndex1MAPI) (interface{}, error) {
	currOffset, err := r.Seek(0, 1)
	if err != nil {
		return nil, err
	}
	defer r.Seek(currOffset, 0)

	switch m.ReferenceType {
	case pstInt16:
		b := make([]byte, 4)
		binary.LittleEndian.PutUint32(b, m.Value)
		var i int16
		if err := binary.Read(bytes.NewReader(b), binary.LittleEndian, &i); err != nil {
			return nil, err
		}
		return i, nil
	case pstInt32:
		b := make([]byte, 4)
		binary.LittleEndian.PutUint32(b, m.Value)
		var i int32
		if err := binary.Read(bytes.NewReader(b), binary.LittleEndian, &i); err != nil {
			return nil, err
		}
		return i, nil
	case pstBool:
		return m.Value != 0, nil
	default:
		internal := m.Value|0x0f != m.Value
		m.Value >>= 4
		if internal {
			if _, err := r.Seek(int64(uint32(desc.IndexOffset)+m.Value+2), 0); err != nil {
				return nil, err
			}

			var pair pstDataIndex1Pair
			if err := binary.Read(r, binary.LittleEndian, &pair); err != nil {
				return nil, err
			}
			// spew.Dump(pair)

			if _, err := r.Seek(int64(pair.Start), 0); err != nil {
				return nil, err
			}

			b := make([]byte, pair.End-pair.Start)
			if _, err := io.ReadFull(r, b); err != nil {
				return nil, err
			}

			switch m.ReferenceType {
			case pstString:
				return string(b), nil
			}
		}
	}

	return nil, nil
}
