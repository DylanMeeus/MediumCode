package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	fmt.Printf("%v\n", leftBitPosition(uint32(10)))
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


func (h hyperLogLog) Add(data []byte) hyperLogLog {
	hash := createHash(data)
	k := 32 - h.b
	_ = k
	_ = hash
	return h
}


func leftBitPosition(x uint32) uint32 {
	// get the previous power of 2
	x = x | (x >> 1)
	x = x | (x >> 2)
	x = x | (x >> 4)
	x = x | (x >> 8)
	x = x | (x >> 16)
	return x - (x >> 1)
}


// create a 32-bit hash
func createHash(stream []byte) uint32 {
	h := fnv.New32()
	h.Write(stream)
	sum := h.Sum32()
	h.Reset()
	return sum
}

