package file

import (
	"testing"
	"math"
	"os"
	"util/think"
	"bytes"
	"strconv"
)

func TestPlay(t *testing.T){
	for i := 0; i < 1; i++ {
		createWav(1000 + i*100)
	}
}

func createWav(hz int){
	file, err := os.OpenFile("./" + strconv.Itoa(hz) + ".wav", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0660)
	think.IsNil(err)
	defer file.Close()

	audio := oneKind(888)
	//
	//audio := someKind(0, 30)
	buf := Header(1, 8000, 32, len(audio)+44)
	buf = append(buf, audio...)
	//
	file.Write(buf)
}

func someKind(startHz, oriL int) []byte{
	audio := bytes.NewBuffer(nil)

	l := 48000

	audioOri := make([][]float32, 0)
	for i := 0; i < oriL; i++ {
		a := productAudio(startHz + i*200)
		audioOri = append(audioOri, a)
	}
	x := l / oriL
	for i := 0; i < l; i++ {
		if i/x >= oriL {
			addLE(audio, audioOri[oriL-1][i])
		} else {
			addLE(audio, audioOri[i/x][i])
		}

	}

	return audio.Bytes()
}

func oneKind(hz int) []byte{
	audio := bytes.NewBuffer(nil)

	audioOri := make([][]float32, 0)
	a := productAudio(hz)
	audioOri = append(audioOri, a)

	return audio.Bytes()
}

// 产生一秒的 48k, 32bit 的 hz,标准响度的原声
func productAudio(hz int) []float32{
	l := 48000 * 1 // 48k * 1s
	audio := make([]float32, 0)

	single := math.Pi * float64(hz) / 48000
	// 产生数字音频
	for i := 0; i < l; i++ {
		a := math.Sin(float64(i) * single)
		//fmt.Println(a, float32(a))
		audio = append(audio, float32(a))
	}

	//fmt.Println(len(audio))
	//for key, value := range audio {
	//	fmt.Print(key, value, ";")
	//
	//}
	return audio
}