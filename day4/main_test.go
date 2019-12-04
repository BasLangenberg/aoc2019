package main

import "testing"

/*

   111111 meets these criteria (double 11, never decreases).
   223450 does not meet these criteria (decreasing pair of digits 50).
   123789 does not meet these criteria (no double).
*/

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		input    int
		validate bool
	}{
		{
			desc:     "111111",
			input:    111111,
			validate: true,
		},
		{
			desc:     "223450",
			input:    223450,
			validate: false,
		},
		{
			desc:     "123789",
			input:    123789,
			validate: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if validate(tC.input) != tC.validate {
				t.Errorf("Failed to validate input %d", tC.input)
			}
		})
	}
}

func TestTwo(t *testing.T) {
	testCases := []struct {
		desc     string
		input    int
		validate bool
	}{
		{
			desc:     "112233",
			input:    112233,
			validate: true,
		},
		{
			desc:     "123444",
			input:    123444,
			validate: false,
		},
		{
			desc:     "111122",
			input:    111122,
			validate: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if validate2(tC.input) != tC.validate {
				t.Errorf("Failed to validate input %d", tC.input)
			}
		})
	}
}
