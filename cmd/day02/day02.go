/*
--- Day 2: Inventory Management System ---

You stop falling through time, catch your breath, and check the screen on the device.

"Destination reached. Current Year: 1518. Current Location: North Pole Utility Closet 83N10." You made it! Now, to find those anomalies.

Outside the utility closet, you hear footsteps and a voice. "...I'm not sure either.
But now that so many people have chimneys, maybe he could sneak in that way?" Another voice responds,
"Actually, we've been working on a new kind of suit that would let him fit through tight spaces like that.
But, I heard that a few days ago, they lost the prototype fabric, the design plans, everything! Nobody on the team can even
seem to remember important details of the project!"

"Wouldn't they have had enough fabric to fill several boxes in the warehouse? They'd be stored together, so the box IDs should be similar.
Too bad it would take forever to search the warehouse for two similar box IDs..." They walk too far away to hear any more.

Late at night, you sneak to the warehouse - who knows what kinds of paradoxes you could cause if you were discovered - and use your fancy
wrist device to quickly scan every box and produce a list of the likely candidates (your puzzle input).

To make sure you didn't miss any, you scan the likely candidate boxes again, counting the number that have an ID containing exactly two of
any letter and then separately counting those with exactly three of any letter. You can multiply those two counts together to get a rudimentary
checksum and compare it to what your device predicts.

For example, if you see the following box IDs:

abcdef contains no letters that appear exactly two or three times.
bababc contains two a and three b, so it counts for both.
abbcde contains two b, but no letter appears exactly three times.
abcccd contains three c, but no letter appears exactly two times.
aabcdd contains two a and two d, but it only counts once.
abcdee contains two e.
ababab contains three a and three b, but it only counts once.

Of these box IDs, four of them contain a letter which appears exactly twice, and three of them contain a letter which appears exactly three times.

Multiplying these together produces a checksum of 4 * 3 = 12.

What is the checksum for your list of box IDs?

--- Part Two ---

Confident that your list of box IDs is complete, you're ready to find the boxes full of prototype fabric.

The boxes will have IDs which differ by exactly one character at the same position in both strings. For example, given the following box IDs:

abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz

The IDs abcde and axcye are close, but they differ by two characters (the second and fourth).
However, the IDs fghij and fguij differ by exactly one character, the third (h and u). Those must be the correct boxes.

What letters are common between the two correct box IDs? (In the example above, this is found by removing
the differing character from either ID, producing fgij.)
*/

package main

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/FollowTheProcess/collections/counter"
	"github.com/FollowTheProcess/msg"
)

//go:embed input.txt
var input string

func main() {
	if err := run(); err != nil {
		msg.Err(err)
		os.Exit(1)
	}
}

func run() error {
	sum, err := checksum(input)
	if err != nil {
		return err
	}
	fmt.Printf("Part 1: %d\n", sum)

	// Thought I might need the box IDs but apparently not
	_, _, common, err := correctBoxIDs(input)
	if err != nil {
		return err
	}

	fmt.Printf("Part 2: %s\n", common)

	return nil
}

// occursExactly returns the lines of input that contain letters that occur exactly n times.
func occursExactly(input string, n int) ([]string, error) {
	input = strings.TrimSpace(input)

	scanner := bufio.NewScanner(strings.NewReader(input))
	var occurs []string
	for scanner.Scan() {
		// Count each character in the line
		line := scanner.Text()
		chars := counter.From([]rune(line))

		for _, count := range chars.Descending() {
			if count == n {
				// Only count things once
				if !slices.Contains(occurs, line) {
					occurs = append(occurs, line)
				}
			}
		}
	}

	return occurs, scanner.Err()
}

// checksum calculates the checksum of the input.
func checksum(input string) (int, error) {
	containsExactlyTwo, err := occursExactly(input, 2) //nolint:mnd
	if err != nil {
		return 0, err
	}
	containsExactlyThree, err := occursExactly(input, 3) //nolint:mnd
	if err != nil {
		return 0, err
	}

	sum := len(containsExactlyTwo) * len(containsExactlyThree)

	return sum, nil
}

// correctBoxIDs returns the ids of the two boxes that only differ by a single character
// at the same index in the ID string.
func correctBoxIDs(input string) (box1, box2, common string, err error) {
	input = strings.TrimSpace(input)

	lines := strings.Fields(input)

	for _, line := range lines {
		for _, otherLine := range lines {
			if ok, commonChars := onlySingleDifference(line, otherLine); ok {
				return line, otherLine, commonChars, nil
			}
		}
	}

	return "", "", "", errors.New("no correct box IDs found")
}

// onlySingleDifference returns whether or not two strings are exactly the same other
// than a single different character at the same index in both, and returns the
// common characters between the two strings.
//
// The strings are assumed to always be the same length, so any difference
// in length returns false.
//
// If a is exactly equal to b, this also returns false, to prevent strings always
// matching against themselves.
func onlySingleDifference(a, b string) (bool, string) {
	if len(a) != len(b) {
		return false, ""
	}

	if a == b {
		return false, ""
	}

	aChars := []rune(a)
	bChars := []rune(b)

	diffs := 0 // Keep track of the number of differences, if this > 1 return false

	var common []rune
	for i, aChar := range aChars {
		bChar := bChars[i]
		if aChar != bChar {
			diffs++
		} else {
			common = append(common, aChar)
		}
	}

	return diffs <= 1, string(common)
}
