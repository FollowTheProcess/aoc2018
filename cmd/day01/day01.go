/*
--- Day 1: Chronal Calibration ---

"We've detected some temporal anomalies," one of Santa's Elves at the Temporal Anomaly Research and Detection Instrument Station tells you.
She sounded pretty worried when she called you down here. "At 500-year intervals into the past, someone has been changing Santa's history!"

"The good news is that the changes won't propagate to our time stream for another 25 days, and we have a device" - she attaches something to your wrist -
"that will let you fix the changes with no such propagation delay. It's configured to send you 500 years further into the past every few days;
that was the best we could do on such short notice."

"The bad news is that we are detecting roughly fifty anomalies throughout time; the device will indicate fixed anomalies with stars.
The other bad news is that we only have one device and you're the best person for the job! Good lu--"
She taps a button on the device and you suddenly feel like you're falling. To save Christmas, you need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first.
Each puzzle grants one star. Good luck!

After feeling like you've been falling for a few minutes, you look at the device's tiny screen.

"Error: Device must be calibrated before first use. Frequency drift detected. Cannot maintain destination lock."

Below the message, the device shows a sequence of changes in frequency (your puzzle input). A value like +6 means the current
frequency increases by 6; a value like -3 means the current frequency decreases by 3.

For example, if the device displays frequency changes of +1, -2, +3, +1, then starting from a frequency of zero, the following changes would occur:

Current frequency  0, change of +1; resulting frequency  1.
Current frequency  1, change of -2; resulting frequency -1.
Current frequency -1, change of +3; resulting frequency  2.
Current frequency  2, change of +1; resulting frequency  3.

In this example, the resulting frequency is 3.

Here are other example situations:

+1, +1, +1 results in  3
+1, +1, -2 results in  0
-1, -2, -3 results in -6

Starting with a frequency of zero, what is the resulting frequency after all of the changes in frequency have been applied?
*/

package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/FollowTheProcess/msg"
	"github.com/FollowTheProcess/parser"
)

//go:embed input.txt
var input string

func main() {
	if err := run(); err != nil {
		msg.Error("%v", err)
		os.Exit(1)
	}
}

func run() error {
	frequency, err := calculateFrequency(input)
	if err != nil {
		return err
	}

	fmt.Printf("Part 1: %d\n", frequency)

	return nil
}

// parseFrequencyChange takes a raw frequency change from the input and parses
// it into something we can actually use in the program.
//
// The result will be a positive or negative integer that can be summed
// to produce the final frequency.
func parseFrequencyChange(input string) (int, error) {
	sign, rest, err := parser.Take(1)(input)
	if err != nil {
		return 0, err
	}

	if sign == "" {
		return 0, fmt.Errorf("input %s has no leading +/- sign", input)
	}

	amount, err := strconv.Atoi(rest)
	if err != nil {
		return 0, fmt.Errorf("input %s has bad amount: %w", input, err)
	}

	switch sign {
	case "+":
		return amount, nil
	case "-":
		return -amount, nil
	default:
		return 0, fmt.Errorf("unexpected sign %s in input %s", sign, input)
	}
}

// calculateFrequency takes the raw input and calculates the final frequency.
func calculateFrequency(input string) (int, error) {
	input = strings.TrimSpace(input)
	scanner := bufio.NewScanner(strings.NewReader(input))

	frequency := 0
	for scanner.Scan() {
		change, err := parseFrequencyChange(scanner.Text())
		if err != nil {
			return 0, err
		}
		frequency += change
	}

	return frequency, nil
}
