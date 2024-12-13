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

func TestFirstRepeatFrequency(t *testing.T) {
	tests := []struct {
		input    string // Raw input
		expected int    // The expected first repeat frequency
	}{
		{
			input:    "+1\n-2\n+3\n+1",
			expected: 2,
		},
		{
			input:    "+1\n-1",
			expected: 0,
		},
		{
			input:    "+3\n+3\n+4\n-2\n-4",
			expected: 10,
		},
		{
			input:    "-6\n+3\n+8\n+5\n-6",
			expected: 5,
		},
		{
			input:    "+7\n+7\n-2\n-7\n-4",
			expected: 14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := firstRepeatFrequency(tt.input)
			test.Ok(t, err)

			test.Equal(t, got, tt.expected)
		})
	}
}
