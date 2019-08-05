package file

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

var (
	RiffID       = [4]byte{'R', 'I', 'F', 'F'}
	WavFormatID  = [4]byte{'W', 'A', 'V', 'E'}
	FmtID        = [4]byte{'f', 'm', 't', ' '}
	DataFormatID = [4]byte{'d', 'a', 't', 'a'}
)

func Header(NumChannels, SampleRates, BitPerSample, fileSize int) []byte {
	var (
		SubChunk1Size = 16 //16 for PCM
		SubChunk2Size = fileSize

		header = bytes.NewBuffer(nil)
	)

	addLE(header, RiffID) //ChunkID

	addLE(header, uint32(4+(8+SubChunk1Size)+(8+SubChunk2Size))) //ChunkSize
	addLE(header, WavFormatID)                                   //Format

	addLE(header, FmtID)                 //SubChunk1ID
	addLE(header, uint32(SubChunk1Size)) //SubChunk1Size
	addLE(header, uint16(1))             //AudioFormat
	addLE(header, uint16(NumChannels))   //NumChannels
	addLE(header, uint32(SampleRates))   //SampleRates

	blockAlign := NumChannels * (BitPerSample / 8)
	addLE(header, uint32(SampleRates*blockAlign)) //ByteRate
	addLE(header, uint16(blockAlign))             //BlockAlign
	addLE(header, uint16(BitPerSample))
	addLE(header, DataFormatID) //SubChunk2ID
	addLE(header, uint32(SubChunk2Size))

	return header.Bytes()
}

func addLE(w io.Writer, src interface{}) error {
	fmt.Println(binary.Size(src), src)
	return binary.Write(w, binary.LittleEndian, src)
}
