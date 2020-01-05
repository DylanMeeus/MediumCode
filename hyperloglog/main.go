package main

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"math"
	"math/bits"
	"math/rand"
	"time"
)

func main() {
	bs, is := getRandomData()
	dt := classicCountDistinct(is)
	h := NewHyperLogLog(64)
	for _, b := range bs {
		h.Add(b)
	}
	hd := h.Count()
	fmt.Printf("classic estimate: %v\n", dt)
	fmt.Printf("hyperloglog estimate: %v\n", hd)
	fmt.Printf("percentage missed: %.2f\n", 100. - (float64(hd) / float64(dt)) * 100)
}

// get random uint32s as a [][]byte slice
func getRandomData() (out [][]byte, intout []uint32) {
	rand.Seed(time.Now().Unix())
	for i := 0; i < math.MaxInt16; i++ {
		i := rand.Uint32()
		b := make([]byte, 4)
		binary.LittleEndian.PutUint32(b, i)
		out = append(out, b)
		intout = append(intout, i)
	}
	return
}

func classicCountDistinct(input []uint32) int {
	m := map[uint32]struct{}{}
	for _, i := range input {
		if _, ok := m[i]; !ok {
			m[i] = struct{}{}
		}
	}
	return len(m)
}

type hyperLogLog struct {
	registers []int
	m         uint // number of registers
	b         uint // bits to calculate [4..16]
}

// Initialize hyperLogLog with m registers
func NewHyperLogLog(m uint) hyperLogLog {
	return hyperLogLog{
		registers: make([]int, m),
		m:         m,
		b:         uint(math.Ceil(math.Log2(float64(m)))),
	}
}

func (h hyperLogLog) Add(data []byte) hyperLogLog {
	x := createHash(data)
	k := 32 - h.b // first b bits
	r := leftmostActiveBit(x << h.b)
	j := x >> uint(k)

	if r > h.registers[j] {
		h.registers[j] = r
	}
	return h
}

func (h hyperLogLog) Count() uint64 {
	sum := 0.
	m := float64(h.m)
	for _, v := range h.registers {
		sum += math.Pow(math.Pow(2, float64(v)),-1)
	}
	estimate := .79402 * m * m / sum
	return uint64(estimate)
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
