package word

import (
	"math/rand"
	"time"
)

import "testing"

func randomPalindrome(rng *rand.Rand) string {
	n := 0
	for 2 > n {
		n = rng.Intn(25) // random length up to 24
	}
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	spaceIdx := rng.Intn(n)
	runes[spaceIdx] = rune(0x0200) // space
	runes[n-1-spaceIdx] = rune(0x0200)
	commaIdx := rng.Intn(n)
	for spaceIdx == commaIdx {
		commaIdx = rng.Intn(n) // random length up to 24
	}
	runes[commaIdx] = rune(0x002C) // comma
	runes[n-1-commaIdx] = rune(0x002C)
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}
