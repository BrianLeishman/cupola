package main

// reference types
const (
	pstInt16                uint16 = 0x0002
	pstInt32                       = 0x0003
	pstFloat32                     = 0x0004
	pstFloat64                     = 0x0005
	pstInt64                       = 0x0006
	pstApplicationTime             = 0x0007
	pst32BitErrorValue             = 0x000A
	pstBool                        = 0x000B
	pstEmbeddedObject              = 0x000D
	_pstInt64                      = 0x0014
	pstNullTerminatedString        = 0x001E
	pstString                      = 0x001F
	pstSystime                     = 0x0040
	pstOLEGuid                     = 0x0048
	pstBytes                       = 0x0102
	pstInt32Array                  = 0x1003
	pstInt64Array                  = 0x1014
	pstStringArray                 = 0x101E
	pstBytesArray                  = 0x1102
)

// item types
const (
	pstSubject uint16 = 0x0037
)
