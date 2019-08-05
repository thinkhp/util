package file

import (
	"testing"
	"os"
	"util/think"
	"fmt"
)

func TestAddHeader(t *testing.T){
	p, err := os.Open("./test.pcm")
	think.IsNil(err)
	defer p.Close()
	info, err := p.Stat()
	think.IsNil(err)

	buf := make([]byte, info.Size())
	l, err := p.Read(buf)
	think.IsNil(err)

	fmt.Println(l, info.Size())


	w, err := os.Create("./test.wav")
	think.IsNil(err)
	defer w.Close()

	w.Write(Header(1,8000, 16, int(info.Size())))
	w.Write(buf)
}



