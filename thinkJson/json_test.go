package thinkJson

import (
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	i := 0
	fmt.Println(string(MustMarshal(i)))
	s := ""
	fmt.Println(string(MustMarshal(s)))
}
