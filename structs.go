package main

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

type pst64Index1LeafNodeItem struct {
	IID    uint64
	Offset uint64
	Size   uint32
	_      uint32
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
	ParentDID uint64
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
