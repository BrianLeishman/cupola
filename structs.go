package main

import "fmt"

const (
	pstSignature uint32 = 0x4e444221
)

type pstHeader struct {
	Signature uint32
	_         [6]byte
	IndexType uint8
}

type pst64Header struct {
	_              [173]byte
	TotalFileSize  uint64
	_              [24]byte
	BackPointer2   uint64
	OffsetIndex2   uint64
	BackPointer1   uint64
	OffsetIndex1   uint64
	_              [265]byte
	EncryptionType uint8
}

type pst64IndexNode struct {
	Items        [488]byte
	ItemCount    uint8
	MaxItemCount uint8
	ItemSize     uint8
	NodeLevel    uint8
	_            [12]byte
	BackPointer  uint64
}

// index 1
type pst64Index1NodeItem struct {
	IID         uint64
	BackPointer uint64
	Offset      uint64
}

func (n pst64Index1NodeItem) String() string {
	return fmt.Sprintf("%d", n.IID)
}

type pst64Index1LeafNodeItem struct {
	IID    uint64
	Offset uint64
	Size   uint16
	_      uint16
	_      [4]byte
}

type pstDataIndex1Pair struct {
	Start uint16
	End   uint16
}

type pstDataIndex1B5 struct {
	Signature  uint16
	DataSize   uint16
	DescOffset uint32
}

type pstDataIndex1MAPI struct {
	ItemType      uint16
	ReferenceType uint16
	Value         uint32
}

func (m pstDataIndex1MAPI) String() string {
	return fmt.Sprintf("(ReferenceType: %s, ItemType: %s, Value: 0x%08x)", referenceName(m.ReferenceType), itemName(m.ItemType), m.Value)
}

// index 2
type pst64Index2NodeItem struct {
	DID         uint64
	BackPointer uint64
	Offset      uint64
}

type pst64Index2LeafNodeItem struct {
	DID       uint64
	DescIID   uint64
	TreeIID   uint64
	ParentDID uint32
	_         uint32
}

type pst64AssociatedTree struct {
	Signature uint16
	Count     uint16
	_         uint32
}

type pst64AssociatedTreeItem struct {
	ID2      uint32
	_        uint16
	_        uint16
	IID      uint64
	ChildIID uint64
}

type pstAssociatedDescriptorItem struct {
	IndexOffset uint16
	Signature   uint16
	B5Offset    uint32
}
