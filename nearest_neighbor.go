package lof

type nearestNeighbor struct {
	points  []Point
	k       int
	nearest map[Point]Points
}

func NearestNeighbor(points []Point, k int) map[Point]Points {
	nn := nearestNeighbor{
		points:  points,
		k:       k,
		nearest: make(map[Point]Points),
	}
	nn.scan()
	return nn.nearest
}

func (nn *nearestNeighbor) scan() {
	for i := 0; i < len(nn.points); i++ {
		point := nn.points[i]
		nearest := LimitedSizeNeighbors{Limit: nn.k}
		for j := 0; j < len(nn.points); j++ {
			if i != j {
				other := nn.points[j]
				nearest.Add(other, point.Distance(other))
			}
		}
		nn.nearest[point] = nearest.GetNearest()
	}
}
