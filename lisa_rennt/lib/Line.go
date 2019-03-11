package lib

type Line struct {
	A Vertex
	B Vertex
}

/*func (l *Line) SortForX() {
	if l.B.X < l.A.X {
		l.A, l.B = l.B, l.A
	}
}

func (l *Line) CalcGradient() (float64, bool) {
	if l.B.X > l.A.X {
		return (l.B.Y - l.A.Y) / (l.B.X - l.A.X), false
	} else if l.A.X > l.B.X {
		return (l.A.Y - l.B.Y) / (l.A.X - l.B.X), false
	}
	return math.MaxFloat64, true
}

func (l *Line) CalcYIntercept() (float64, bool) {
	m, isVert := l.CalcGradient()
	if isVert {
		return 0, true
	}
	return l.A.Y - (l.A.X * m), false
}*/
