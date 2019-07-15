package thinkJson

import (
	"testing"
	"fmt"
)

func TestJson(t *testing.T){
	i := 0
	fmt.Println(string(MustMarshal(i)))
	s := ""
	fmt.Println(string(MustMarshal(s)))
}
