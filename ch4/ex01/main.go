package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%d\n", popCountArray(xor(&c1, &c2)))
}

func xor(c1, c2 *[32]byte) [32]byte {
	var d [32]byte

	for i := 0; i < len(d); i++ {
		d[i] = c1[i] ^ c2[i]
	}

	return d
}

func popCountArray(array [32]byte) (count int) {
	for _, v := range array {
		count += popCount(v)
	}
	return
}

func popCount(v byte) (count int) {
	for v != 0 {
		v &= (v - 1)
		count++
	}
	return
}
