package bmp

// bmp head
type header struct {
	// BitmapFileHeader 14字节
	sigBM     [2]byte   // 文件类型
	fileSize  uint32    // 文件大小
	resverved [2]uint16 // 保留 2个2字节
	pixOffset uint32    // 位图数据偏移
	// BitmapInfoHeader 40字节
	dibHeaderSize   uint32 // 信息头大小 40
	width           uint32 // 宽
	height          uint32 // 高
	colorPlane      uint16 //
	bpp             uint16
	compression     uint32
	imageSize       uint32
	xPixelsPerMeter uint32
	yPixelsPerMeter uint32
	colorUse        uint32
	colorImportant  uint32
}
