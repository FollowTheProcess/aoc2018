package main

import (
	"fmt"
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

func TestOnlySingleDifference(t *testing.T) {
	tests := []struct {
		a, b   string
		common string // Common characters between the two
		want   bool
	}{
		{
			a:      "abcde",
			b:      "axcye",
			want:   false,
			common: "ace",
		},
		{
			a:      "fghij",
			b:      "fguij",
			want:   true,
			common: "fgij",
		},
		{
			a:      "abcde",
			b:      "klmno",
			want:   false,
			common: "",
		},
		{
			a:      "hello",
			b:      "helpo",
			want:   true,
			common: "helo",
		},
		{
			a:      "hello",
			b:      "helpi",
			want:   false,
			common: "hel",
		},
		{
			a:      "hello",
			b:      "hello",
			want:   false,
			common: "", // All common but no diff so no common
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("%s-%s", tt.a, tt.b)
		t.Run(name, func(t *testing.T) {
			onlyOneDiff, common := onlySingleDifference(tt.a, tt.b)

			test.Equal(t, onlyOneDiff, tt.want)
			test.Equal(t, common, tt.common)
		})
	}
}

func TestCorrectBoxIDs(t *testing.T) {
	// Because we use strings.Fields, we don't have to care about indent
	boxes := `abcde
	fghij
	klmno
	pqrst
	fguij
	axcye
	wvxyz`

	box1, box2, common, err := correctBoxIDs(boxes)
	test.Ok(t, err)
	test.Equal(t, box1, "fghij")
	test.Equal(t, box2, "fguij")
	test.Equal(t, common, "fgij")
}
