package sliceutil

import (
	"math/rand"
	"time"
)

// Shuffle returns a shuffled copy of src.
// The original slice is NOT modified.
func Shuffle[T any](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Fisher–Yates shuffle
	for i := len(dst) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		dst[i], dst[j] = dst[j], dst[i]
	}

	return dst
}
