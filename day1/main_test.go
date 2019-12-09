package main

import (
	"testing"
)

func Test_fuelforfuel(t *testing.T) {
	i := fuelFuel(100756)
	if i != 50346 {
		t.Errorf("Failed to get correct answer, got: %d, want: %d", i, 50346)
	}

	i = fuelFuel(1969)
	if i != 966 {
		t.Errorf("Failed to get correct answer, got: %d, want: %d", i, 966)
	}
}
