package lof

import "testing"

func TestLimitedSize(t *testing.T) {
	limited := LimitedSizeNeighbors{Limit: 2}

	limited.Add(Point{1, 1}, 1)
	checkSize(t, limited, 1)
	checkIncludes(t, limited, Point{1, 1})

	limited.Add(Point{2, 2}, 2)
	checkSize(t, limited, 2)
	checkIncludes(t, limited, Point{1, 1})
	checkIncludes(t, limited, Point{2, 2})

	limited.Add(Point{3, 3}, 3)
	checkSize(t, limited, 2)
	checkIncludes(t, limited, Point{1, 1})
	checkIncludes(t, limited, Point{2, 2})

	limited.Add(Point{0.5, 0.5}, 1)
	checkSize(t, limited, 2)
	checkIncludes(t, limited, Point{1, 1})
	checkIncludes(t, limited, Point{0.5, 0.5})
}

func checkIncludes(t *testing.T, limited LimitedSizeNeighbors, point Point) {
	nearest := limited.GetNearest()
	for i := 0; i < len(nearest); i++ {
		if point.Equals(nearest[i]) {
			return
		}
	}
	t.Fatalf("point %+v not located", point)
}

func checkSize(t *testing.T, limited LimitedSizeNeighbors, expectedSize int) {
	nearest := limited.GetNearest()
	actualSize := len(nearest)
	if actualSize != expectedSize {
		t.Fatalf("expected size %v, actual is %v", expectedSize, actualSize)
	}
}
