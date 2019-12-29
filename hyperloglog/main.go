package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	fmt.Println("starting up")
}

type hyperLogLog struct {
	registers []uint
	m         uint // number of registers
	b         uint // bits to calculate [4..16]
}

// Initialize hyperLogLog with m registers
func NewHyperLogLog(m uint) hyperLogLog {
	return hyperLogLog{
		registers: make([]uint, m),
		m:         m,
	}
}

// create a 32-bit hash
func createHash(stream []byte) uint32 {
	h := fnv.New32()
	h.Write(stream)
	sum := h.Sum32()
	h.Reset()
	return sum
}
