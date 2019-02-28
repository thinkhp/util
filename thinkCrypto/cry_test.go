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
