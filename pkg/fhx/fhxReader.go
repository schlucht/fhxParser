package fhx

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"unicode/utf16"
)

func ReadUTF16(path string) (string, error) {
	f, err := os.Open(path)
	var str = ""
	if err != nil {
		return "", err
	}
	defer f.Close()
	buf := make([]byte, 1024)
	i := 0
	for {
		n, err := f.Read(buf)
		i++
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		if n > 0 {
			s, err := DecodeUtf16(buf[:n], binary.LittleEndian)
			if err != nil {
				return "", err
			}
			str = str + s
		}
	}
	return str, nil
}

func DecodeUtf16(b []byte, order binary.ByteOrder) (string, error) {
	ints := make([]uint16, len(b)/2)
	if err := binary.Read(bytes.NewReader(b), order, &ints); err != nil {
		return "", err
	}
	return string(utf16.Decode(ints)), nil
}
