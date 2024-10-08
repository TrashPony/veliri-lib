package game_math

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func IntFromBytes(b []byte) int32 {
	var x int32
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.BigEndian, &x)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}

	return x
}

func IntFromBytesFromJS(b []byte) int {
	var value int
	value |= int(b[0])
	value |= int(b[1]) << 8
	value |= int(b[2]) << 16
	value |= int(b[3]) << 24

	return value
}

func GetIntBytes(x int) []byte {

	b := []byte{
		byte(0xff & (x >> 24)),
		byte(0xff & (x >> 16)),
		byte(0xff & (x >> 8)),
		byte(0xff & x),
	}

	return b
}

func GetInt64Bytes(x int64) []byte {

	b := []byte{
		byte(0xff & (x >> 56)),
		byte(0xff & (x >> 48)),
		byte(0xff & (x >> 40)),
		byte(0xff & (x >> 32)),
		byte(0xff & (x >> 24)),
		byte(0xff & (x >> 16)),
		byte(0xff & (x >> 8)),
		byte(0xff & x),
	}

	return b
}

func ReuseByteSlice(s *[]byte, startIndex int, appendSlice []byte) {
	for i, b := range appendSlice {
		if len(*s) > startIndex+i {
			(*s)[startIndex+i] = b
		} else {
			*s = append(*s, b)
		}
	}
}

func BoolToByte(bitSet bool) byte {
	if bitSet {
		return 1
	}

	return 0
}
