package main

import "testing"

func Test_ClosestIntersect(t *testing.T) {
	testCases := []struct {
		desc   string
		wire1  string
		wire2  string
		output int
	}{
		{
			desc:   "Should give 159",
			wire1:  "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			wire2:  "U62,R66,U55,R34,D71,R55,D58,R83",
			output: 159,
		},
		{
			desc:   "Should give 135",
			wire1:  "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			wire2:  "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			output: 135,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			crds := intersects(tC.wire1, tC.wire2)
			result := closest(crds)
			if result != tC.output {
				t.Errorf("Output incorrect. Want: %d, Got: %d", tC.output, result)
			}
		})
	}
}

func Test(t *testing.T) {
	testCases := []struct {
		desc   string
		wire1  string
		wire2  string
		output int
	}{
		{
			desc:   "Should give 159",
			wire1:  "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			wire2:  "U62,R66,U55,R34,D71,R55,D58,R83",
			output: 610,
		},
		{
			desc:   "Should give 135",
			wire1:  "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			wire2:  "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			output: 410,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ints := intersects(tC.wire1, tC.wire2)
			wire1 := parseCoords(tC.wire1)
			wire2 := parseCoords(tC.wire2)
			result := getShortestPath(ints, wire1, wire2)
			if result != tC.output {
				t.Errorf("Unable to find shortest path. Got: %d, Want: %d", result, tC.output)
			}
		})
	}
}
