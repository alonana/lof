package lof

import (
	"fmt"
	"testing"
)

func TestLof(t *testing.T) {
	var points Points

	p1 := Point{1, 1}
	p2 := Point{2, 1}
	p3 := Point{1, 2}
	outlier := Point{5, 5}

	points = append(points, p1)
	points = append(points, p2)
	points = append(points, p3)
	points = append(points, outlier)

	lof := Lof(points, 2)
	if len(lof) != 4 {
		t.Fatalf("wrong size")
	}

	for p, score := range lof {
		fmt.Printf("%+v %v\n", p, score)
	}

	if lof[outlier] < lof[p1] {
		t.Fatalf("invalid lof")
	}
}
