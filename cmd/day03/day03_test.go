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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseClaim(tt.input)
			test.Ok(t, err)

			test.Equal(t, got, tt.want)
		})
	}
}
