package util

import (
	"bytes"
	"encoding/binary"
	"math"
)

func Byte2Int(b []byte) int {
	b_buf  :=  bytes.NewBuffer(b)
	var x int
	binary.Read(b_buf, binary.BigEndian, &x)
	return x
}
/*func Byte2Int2(b []byte) int {
	x, _ := strconv.ParseInt(string(b), 10, 64)
	return x
}*/
func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}