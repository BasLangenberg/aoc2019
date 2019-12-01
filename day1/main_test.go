package main

import "testing"

func Test_fuelforfuel(t *testing.T) {
	i := fuelforfuel(100756)
	if i != 50346+100756 {
		t.Errorf("Failed to get correct answer, got: %d, want: %d", i, 100756+50346)
	}
}
