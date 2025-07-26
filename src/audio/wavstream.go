package audio

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/go-audio/wav"
	"github.com/hajimehoshi/ebiten/v2/audio"
	ebiwav "github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type WavStream struct {
	ctx     *audio.Context
	player  *audio.Player
	played  bool
	options WavStreamOptions
}

type WavStreamOptions struct {
	Loop bool
}

func getWavSampleRate(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	decoder := wav.NewDecoder(file)
	if !decoder.IsValidFile() {
		return 0, fmt.Errorf("invalid WAV file")
	}

	format := decoder.Format()

	return int(format.SampleRate), nil
}

func CreateWavStream(path string) *WavStream {
	return CreateWavStreamWithOptions(path, WavStreamOptions{})
}

func CreateWavStreamWithOptions(path string, options WavStreamOptions) *WavStream {
	sampleRate, err := getWavSampleRate(path)

	if err != nil {
		log.Printf("%v", err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Printf("%v", err)

	}

	d, err := ebiwav.DecodeWithoutResampling(file)
	if err != nil {
		log.Printf("%v", err)
	}

	ctx := audio.CurrentContext()

	if ctx == nil {
		ctx = audio.NewContext(sampleRate)
	}

	p, err := ctx.NewPlayer(d)
	if err != nil {
		log.Printf("%v", err)
	}

	return &WavStream{
		ctx:     ctx,
		player:  p,
		played:  false,
		options: options,
	}
}

func (wavStream *WavStream) Play() {
	if !wavStream.played {
		wavStream.player.Play()
		wavStream.played = true
	} else {
		if wavStream.options.Loop && !wavStream.player.IsPlaying() {
			wavStream.player.Rewind()
			wavStream.played = false
		}
	}
}

func (wavStream *WavStream) Rewind() {
	wavStream.player.Rewind()
	wavStream.played = false
}
