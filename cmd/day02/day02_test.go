package main

import (
	"slices"
	"testing"

	"github.com/FollowTheProcess/test"
)

const testInput = `
abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab
`

func TestOccursExactly(t *testing.T) {
	t.Run("twice", func(t *testing.T) {
		twice, err := occursExactly(testInput, 2)
		test.Ok(t, err)

		want := []string{
			"bababc", // contains two a and three b, so it counts for both.
			"abbcde", // contains two b, but no letter appears exactly three times.
			"aabcdd", // contains two a and two d, but it only counts once.
			"abcdee", // contains two e.
		}

		slices.Sort(twice)
		slices.Sort(want)

		test.EqualFunc(t, twice, want, slices.Equal)
	})

	t.Run("thrice", func(t *testing.T) {
		twice, err := occursExactly(testInput, 3)
		test.Ok(t, err)

		want := []string{
			"bababc", // contains two a and three b, so it counts for both.
			"abcccd", // contains three c, but no letter appears exactly two times.
			"ababab", // contains three a and three b, but it only counts once.
		}

		slices.Sort(twice)
		slices.Sort(want)

		test.EqualFunc(t, twice, want, slices.Equal)
	})
}

func TestChecksum(t *testing.T) {
	sum, err := checksum(testInput)
	test.Ok(t, err)

	test.Equal(t, sum, 12)
}
