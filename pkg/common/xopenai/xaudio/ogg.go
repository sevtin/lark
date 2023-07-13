package xaudio

import (
	"bytes"
	"errors"
	"github.com/pion/opus"
	"github.com/pion/opus/pkg/oggreader"
	"io"
	"os"
)

func OggToWavByPath(ogg string, wav string) (err error) {
	var (
		input  *os.File
		output *os.File
	)
	if input, err = os.Open(ogg); err != nil {
		return err
	}
	defer input.Close()

	if output, err = os.Create(wav); err != nil {
		return err
	}
	defer output.Close()
	return OggToWav(input, output)
}

func OggToWav(input io.Reader, output io.WriteSeeker) (err error) {
	var (
		ogg      *oggreader.OggReader
		out      []byte
		decoder  opus.Decoder
		encoder  *Encoder
		segments [][]byte
		i        int
	)
	if ogg, _, err = oggreader.NewWith(input); err != nil {
		return
	}
	out = make([]byte, 1920)
	decoder = opus.NewDecoder()
	encoder = NewEncoder(output, 44100, 16)
	defer encoder.Close()
	for {
		segments, _, err = ogg.ParseNextPage()
		if errors.Is(err, io.EOF) {
			break
		} else if bytes.HasPrefix(segments[0], []byte("OpusTags")) {
			continue
		}
		if err != nil {
			return
		}
		for i = range segments {
			if _, _, err = decoder.Decode(segments[i], out); err != nil {
				return
			}
			encoder.Write(out)
		}
	}
	return
}
