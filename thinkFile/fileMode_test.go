package thinkFile

import (
	"fmt"
	"os"
	"testing"
)

func TestFileMode(t *testing.T){
	fmt.Println(os.ModePerm.String())
}
