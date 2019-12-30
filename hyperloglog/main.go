package main

import (
	"fmt"
	"hash/fnv"
	"math/bits"
)

func main() {
	fmt.Printf("%v\n", leftmostActiveBit(uint32(16)))
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


func leftmostActiveBit(x uint32) int {
	return 1 + bits.LeadingZeros32(x)
}


// create a 32-bit hash
func createHash(stream []byte) uint32 {
	h := fnv.New32()
	h.Write(stream)
	sum := h.Sum32()
	h.Reset()
	return sum
}

