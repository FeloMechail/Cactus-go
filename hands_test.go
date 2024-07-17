package main

import (
	"testing"
)

type testCase struct {
	arg1 []uint32
	want string
}

func TestHandRank(t *testing.T) {
	cases := []testCase{
		{[]uint32{69634, 135427, 529159, 4199953, 8394515}, "Flush"},
		// Straight
		{[]uint32{1053707, 2131213, 4212241, 8398611, 16783383}, "Straight"},
		// Full House
		{[]uint32{134253349, 134236965, 134228773, 67119647, 67115551}, "Full House"},
		// Four of a Kind
		{[]uint32{98306, 81922, 73730, 69634, 135427}, "Four of a Kind"},
		// Three of a Kind
		{[]uint32{98306, 81922, 73730, 533255, 1057803}, "Three of a Kind"},
		// Two Pair
		{[]uint32{98306, 81922, 164099, 147715, 270853}, "Two Pair"},
		// One Pair
		{[]uint32{98306, 81922, 270853, 533255, 1057803}, "One Pair"},
		// High Card
		{[]uint32{268442665, 139523, 270853, 533255, 1057803}, "High Card"},
		//straight flush
		{[]uint32{268471337, 16812055, 33589533, 67144223, 134253349}, "Straight Flush"},
	}

	for _, tc := range cases {
		got := handRank(tc.arg1)
		if got != tc.want {
			t.Errorf("got %s want %s", got, tc.want)
		}
	}

}
