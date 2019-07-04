package thinkCrypto

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestEncoding(t *testing.T) {
	// utf-8
	str := "你好"
	b := []byte(str)
	fmt.Println(b, str)

	// hex
	str = hex.EncodeToString(b)
	fmt.Println(b, str)
}

func TestMd5(t *testing.T) {
	s := "15891467397"
	fmt.Println(len(GetMD5(s)))
}
