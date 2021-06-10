package scramble

import (
	"math/rand"
	"strings"
	"time"
)

// Scramble will return a string of same length of str
// with the same letters of str in a different order.
func Scramble(str string) (result string) {
	s := strings.Split(str, "")

	for range s {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(len(s))
		// if the random number between 1 and len(s) is the same as length then minus 1
		if r == len(s) {
			r -= 1
		}
		// build result string adding chars
		result += s[r]
		// rebuilt s after removing s[r]
		s = append(s[:r], s[r+1:]...)
	}
	return
}
