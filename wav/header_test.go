package wav

//import (
//	"fmt"
//	"github.com/go-audio/riff"
//	"github.com/go-audio/wav"
//)

//// AddLE serializes and adds the passed value using little endian
//func (e *Encoder) AddLE(src interface{}) error {
//	e.WrittenBytes += binary.Size(src)
//	return binary.Write(e.w, binary.LittleEndian, src)
//}
//
//// AddBE serializes and adds the passed value using big endian
//func (e *Encoder) AddBE(src interface{}) error {
//	e.WrittenBytes += binary.Size(src)
//	return binary.Write(e.w, binary.BigEndian, src)
////}

//第三方库
//func addHead(e *wav.Encoder) error {
//	// riff ID
//	if err := e.AddLE(riff.RiffID); err != nil {
//		return err
//	}
//	// file size uint32, to update later on.
//	if err := e.AddLE(uint32(42)); err != nil {
//		return err
//	}
//	// wave headers
//	if err := e.AddLE(riff.WavFormatID); err != nil {
//		return err
//	}
//	// form
//	if err := e.AddLE(riff.FmtID); err != nil {
//		return err
//	}
//	// chunk size
//	if err := e.AddLE(uint32(16)); err != nil {
//		return err
//	}
//	// wave format
//	if err := e.AddLE(uint16(e.WavAudioFormat)); err != nil {
//		return err
//	}
//	// num channels
//	if err := e.AddLE(uint16(e.NumChans)); err != nil {
//		return fmt.Errorf("error encoding the number of channels - %v", err)
//	}
//	// samplerate
//	if err := e.AddLE(uint32(e.SampleRate)); err != nil {
//		return fmt.Errorf("error encoding the sample rate - %v", err)
//	}
//	blockAlign := e.NumChans * e.BitDepth / 8
//	// avg bytes per sec
//	if err := e.AddLE(uint32(e.SampleRate * blockAlign)); err != nil {
//		return fmt.Errorf("error encoding the avg bytes per sec - %v", err)
//	}
//	// block align
//	if err := e.AddLE(uint16(blockAlign)); err != nil {
//		return err
//	}
//	// bits per sample
//	if err := e.AddLE(uint16(e.BitDepth)); err != nil {
//		return fmt.Errorf("error encoding bits per sample - %v", err)
//	}
//
//	return nil
//}
