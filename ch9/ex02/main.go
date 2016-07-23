package popcount

import "sync"

var pc [256]byte

func loadPC() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

var loadPCOnce sync.Once

func PopCount(x uint64) int {
	loadPCOnce.Do(loadPC)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
