/*
--- Day 3: No Matter How You Slice It ---

The Elves managed to locate the chimney-squeeze prototype fabric for Santa's suit (thanks to someone who helpfully wrote its box
IDs on the wall of the warehouse in the middle of the night).

Unfortunately, anomalies are still affecting them - nobody can even agree on how to cut the fabric.

The whole piece of fabric they're working on is a very large square - at least 1000 inches on each side.

Each Elf has made a claim about which area of fabric would be ideal for Santa's suit. All claims have an ID and consist of a
single rectangle with edges parallel to the edges of the fabric. Each claim's rectangle is defined as follows:

The number of inches between the left edge of the fabric and the left edge of the rectangle.
The number of inches between the top edge of the fabric and the top edge of the rectangle.
The width of the rectangle in inches.
The height of the rectangle in inches.

A claim like #123 @ 3,2: 5x4 means that claim ID 123 specifies a rectangle 3 inches from the left edge, 2 inches from the top edge, 5 inches wide,
and 4 inches tall. Visually, it claims the square inches of fabric represented by # (and ignores the square inches of fabric represented by .)
in the diagram below:

...........
...........
...#####...
...#####...
...#####...
...#####...
...........
...........
...........

The problem is that many of the claims overlap, causing two or more claims to cover part of the same areas. For example, consider the following claims:

#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2

Visually, these claim the following areas:

........
...2222.
...2222.
.11XX22.
.11XX22.
.111133.
.111133.
........

The four square inches marked with X are claimed by both 1 and 2. (Claim 3, while adjacent to the others, does not overlap either of them.)

If the Elves all proceed with their own plans, none of them will have enough fabric. How many square inches of fabric are within two or more claims?
*/

package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/FollowTheProcess/msg"
	"github.com/FollowTheProcess/parser"
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
	fmt.Println(input)
	return nil
}

type Claim struct {
	ID       string // The ID of the claim
	FromLeft int    // The distance in inches from the left edge
	FromTop  int    // The distance in inches from the top edge
	Width    int    // The width of the claim in inches
	Height   int    // The height of the claim in inches
}

// parseClaim parses a Claim from text.
func parseClaim(raw string) (Claim, error) {
	// I literally wrote this parsing library for advent of code, glad
	// it's coming in handy
	values, rest, err := parser.Chain(
		parser.Exact("#"),
		parser.TakeWhile(unicode.IsDigit), // #<id>
		parser.TakeWhile(unicode.IsSpace), // Consume any whitespace
		parser.Exact("@"),
		parser.TakeWhile(unicode.IsSpace), // Consume any whitespace
		parser.TakeWhile(unicode.IsDigit), // Inches from left
		parser.Exact(","),
		parser.TakeWhile(unicode.IsDigit), // Inches from top
		parser.Exact(":"),
		parser.TakeWhile(unicode.IsSpace), // Consume any whitespace
		parser.TakeWhile(unicode.IsDigit), // Width in inches
		parser.Exact("x"),
		parser.TakeWhile(unicode.IsDigit), // Height in inches
	)(raw)
	if err != nil {
		return Claim{}, fmt.Errorf("could not parse Claim from %q: %w", raw, err)
	}

	if rest != "" {
		return Claim{}, fmt.Errorf("unconsumed input: %q", rest)
	}

	if len(values) != 13 { //nolint:mnd
		return Claim{}, fmt.Errorf(
			"wrong number of parsed values, expected 13, got %d",
			len(values),
		)
	}

	// These must be valid integers if we got here so no need to check the err
	fromLeft, _ := strconv.Atoi(values[5]) //nolint:errcheck
	fromTop, _ := strconv.Atoi(values[7])  //nolint:errcheck
	width, _ := strconv.Atoi(values[10])   //nolint:errcheck
	height, _ := strconv.Atoi(values[12])  //nolint:errcheck

	claim := Claim{
		ID:       values[1],
		FromLeft: fromLeft,
		FromTop:  fromTop,
		Width:    width,
		Height:   height,
	}

	return claim, nil
}
