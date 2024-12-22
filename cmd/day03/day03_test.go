package main

import (
	"testing"

	"github.com/FollowTheProcess/test"
)

func TestParseClaim(t *testing.T) {
	tests := []struct {
		name  string // Name of the test case
		input string // The input text to parse
		want  Claim  // Expected Claim
	}{
		{
			name:  "example 1",
			input: "#123 @ 3,2: 5x4",
			want: Claim{
				ID:       "123",
				FromLeft: 3,
				FromTop:  2,
				Width:    5,
				Height:   4,
			},
		},
		{
			name:  "example 2",
			input: "#1 @ 1,3: 4x4",
			want: Claim{
				ID:       "1",
				FromLeft: 1,
				FromTop:  3,
				Width:    4,
				Height:   4,
			},
		},
		{
			name:  "example 3",
			input: "#2 @ 3,1: 4x4",
			want: Claim{
				ID:       "2",
				FromLeft: 3,
				FromTop:  1,
				Width:    4,
				Height:   4,
			},
		},
		{
			name:  "example 3",
			input: "#3 @ 5,5: 2x2",
			want: Claim{
				ID:       "3",
				FromLeft: 5,
				FromTop:  5,
				Width:    2,
				Height:   2,
			},
		},
		{
			name:  "insane",
			input: "#99999 @ 12735,91725: 127x9725",
			want: Claim{
				ID:       "99999",
				FromLeft: 12735,
				FromTop:  91725,
				Width:    127,
				Height:   9725,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseClaim(tt.input)
			test.Ok(t, err)

			test.Equal(t, got, tt.want)
		})
	}
}
