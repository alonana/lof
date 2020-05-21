package lof

type pointAndDistance struct {
	point    Point
	distance float32
}

type LimitedSizeNeighbors struct {
	Limit   int
	nearest []pointAndDistance
}

func (l *LimitedSizeNeighbors) Add(point Point, distance float32) {
	newPoint := pointAndDistance{
		point:    point,
		distance: distance,
	}

	if len(l.nearest) < l.Limit {
		l.nearest = append(l.nearest, newPoint)
		return
	}

	for i := 0; i < len(l.nearest); i++ {
		existing := l.nearest[i]
		if existing.distance > newPoint.distance {
			l.nearest[i] = newPoint
			return
		}
	}
}

func (l *LimitedSizeNeighbors) GetNearest() Points {
	var points Points
	for i := 0; i < len(l.nearest); i++ {
		points = append(points, l.nearest[i].point)
	}
	return points
}
