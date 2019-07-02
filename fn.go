package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

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

func pst64SearchIndex1(f *os.File, offset uint64, iid uint64, stack []uint64) (dataOffset uint64, size uint32, err error) {
	log.Println(stack)
	log.Println("searching for I-ID", iid)

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
			log.Println("considering I-ID", item.IID)
			if item.IID == iid {
				os.Exit(0)
				return item.Offset, item.Size, nil
			}
		}

		return 0, 0, nil
	}

	items := make([]pst64Index1NodeItem, node.ItemCount)
	if err := binary.Read(bytes.NewReader(node.Items[:]), binary.LittleEndian, &items); err != nil {
		return 0, 0, err
	}

	log.Println("checking items")
	for _, item := range items {
		log.Print(item.IID)
		newStack := make([]uint64, len(stack), len(stack)+1)
		copy(newStack, stack)
		newStack = append(stack, item.IID)
		// fmt.Println(pst64SearchIndex1(f, item.Offset, iid, newStack))
	}
	return 0, 0, nil
}
