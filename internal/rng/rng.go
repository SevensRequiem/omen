package rng

import (
	"crypto/rand"
	"encoding/binary"
	"math/big"
)

func Intn(n int) int {
	val, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	if err != nil {
		panic("crypto/rand failed: " + err.Error())
	}
	return int(val.Int64())
}

func Bool() bool {
	var b [1]byte
	_, err := rand.Read(b[:])
	if err != nil {
		panic("crypto/rand failed: " + err.Error())
	}
	return b[0]&1 == 1
}

func Shuffle(n int, swap func(i, j int)) {
	for i := n - 1; i > 0; i-- {
		j := Intn(i + 1)
		swap(i, j)
	}
}

func Uint32() uint32 {
	var b [4]byte
	_, err := rand.Read(b[:])
	if err != nil {
		panic("crypto/rand failed: " + err.Error())
	}
	return binary.LittleEndian.Uint32(b[:])
}
