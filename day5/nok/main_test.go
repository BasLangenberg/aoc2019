package main

import (
	"testing"
)

const testcaseone = "1,9,10,3,2,3,11,0,99,30,40,50"

func Test_compute(t *testing.T) {
	t.Run("Test case 1", func(t *testing.T) {

		data := getInput(testcaseone)
		i := compute(data)
		if i != 3500 {
			t.Errorf("Did not get correct output. Want: %d, Got: %d", 3500, i)
		}
	})
}
