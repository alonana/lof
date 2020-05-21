package lof

import (
	"fmt"
	"testing"
)

func TestNearestNeighbor(t *testing.T) {
	var points Points
	points = append(points, Point{1, 1})
	points = append(points, Point{1, 2})
	points = append(points, Point{2, 1})
	points = append(points, Point{2, 2})
	nearest := NearestNeighbor(points, 2)

	if len(nearest) != 4 {
		t.Fatalf("wrong size")
	}

	for p, neighbors := range nearest {
		fmt.Printf("%+v %+v\n", p, neighbors)
		if len(neighbors) != 2 {
			t.Fatalf("wrong size")
		}
	}
}
