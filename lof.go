package lof

type lof struct {
	nn       map[Point]Points
	k        int
	lofScore map[Point]float32
	lrdScore map[Point]float32
	kScore   map[Point]float32
}

func Lof(points Points, k int) map[Point]float32 {
	nn := NearestNeighbor(points, k)
	l := lof{
		nn:       nn,
		k:        k,
		lrdScore: make(map[Point]float32),
		lofScore: make(map[Point]float32),
		kScore:   make(map[Point]float32),
	}

	l.scanLof()

	return l.lofScore
}

func (l *lof) scanLof() {
	for p := range l.nn {
		l.lof(p)
	}
}

func (l *lof) lof(p Point) {
	neighbors := l.nn[p]
	var sum float32
	for i := 0; i < len(neighbors); i++ {
		neighbor := neighbors[i]
		sum += l.lrd(neighbor)
	}
	l.lofScore[p] = sum / (float32(l.k) * l.lrd(p))
}

func (l *lof) lrd(p Point) float32 {
	score, exists := l.lrdScore[p]
	if !exists {
		score = l.calculateLrd(p)
		l.lrdScore[p] = score
	}
	return score
}

func (l *lof) calculateLrd(p Point) float32 {
	neighbors := l.nn[p]
	var sum float32
	for i := 0; i < len(neighbors); i++ {
		neighbor := neighbors[i]
		sum += l.reachabilityDistance(p, neighbor)
	}

	return float32(l.k) / sum
}

func (l *lof) reachabilityDistance(a Point, b Point) float32 {
	distance := a.Distance(b)
	kDistanceB := l.kDistance(b)
	if distance > kDistanceB {
		return distance
	}
	return kDistanceB
}

func (l *lof) kDistance(p Point) float32 {
	score, exists := l.kScore[p]
	if !exists {
		score = l.calculateKDistance(p)
		l.kScore[p] = score
	}
	return score
}

func (l *lof) calculateKDistance(p Point) float32 {
	neighbors := l.nn[p]
	var max float32
	for i := 0; i < len(neighbors); i++ {
		neighbor := neighbors[i]
		distance := neighbor.Distance(p)
		if distance > max {
			max = distance
		}
	}

	return max
}
