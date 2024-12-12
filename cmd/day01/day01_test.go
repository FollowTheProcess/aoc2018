package main

import (
	"testing"

	"github.com/FollowTheProcess/test"
)

func TestParseFrequencyChange(t *testing.T) {
	tests := []struct {
		input string // The raw frequency change to parse
		want  int    // The positive or negative integer
	}{
		{
			input: "+1",
			want:  1,
		},
		{
			input: "-1",
			want:  -1,
		},
		{
			input: "+267",
			want:  267,
		},
		{
			input: "-1389",
			want:  -1389,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := parseFrequencyChange(tt.input)
			test.Ok(t, err)
			test.Equal(t, got, tt.want)
		})
	}
}

func TestCalculateFrequency(t *testing.T) {
	tests := []struct {
		input    string // The input
		expected int    // The expected frequency
	}{
		{
			input:    "+1\n-2\n+3\n+1",
			expected: 3,
		},
		{
			input:    "+1\n+1\n+1",
			expected: 3,
		},
		{
			input:    "+1\n+1\n-2",
			expected: 0,
		},
		{
			input:    "-1\n-2\n-3",
			expected: -6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			frequency, err := calculateFrequency(tt.input)
			test.Ok(t, err)
			test.Equal(t, frequency, tt.expected)
		})
	}
}
