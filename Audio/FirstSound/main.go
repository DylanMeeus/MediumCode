package main

/* program to create a pitch perfect (440Hz) sound */

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

const (
	Duration   = 2
	SampleRate = 44100
	Frequency  = 440
)

var (
	tau = math.Pi * 2
)

func main() {
	fmt.Fprintf(os.Stderr, "generating sine wave..\n")
	generate()
	fmt.Fprintf(os.Stderr, "done")
}

func generate() {
	var (
		start float64 = 1.0
		end   float64 = 1.0e-4
	)
	nsamps := Duration * SampleRate
	var angle float64 = tau / float64(nsamps)
	file := "out.bin"
	f, _ := os.Create(file)
	decayfac := math.Pow(end/start, 1.0/float64(nsamps))
	for i := 0; i < nsamps; i++ {
		sample := math.Sin(angle * Frequency * float64(i))
		sample *= start
		start *= decayfac
		var buf [8]byte
		binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
		bw, err := f.Write(buf[:])
		if err != nil {
			panic(err)
		}
		fmt.Printf("\rWrote: %v bytes to %s", bw, file)
	}
}
