package lib

import "math"

func LineSegementIntersection(u *Line, v *Line) bool {
	if isBoxCollision(u, v) {
		return isSegmentLineCollision(u, v) && isSegmentLineCollision(v, u)
	}

	return false
}

func isSegmentLineCollision(line *Line, segment *Line) bool {
	transLine := Line{Vertex{}, Vertex{line.B.X - line.A.X, line.B.Y - line.A.Y}}
	transSegment := Line{Vertex{segment.A.X - line.A.X, segment.A.Y - line.A.Y},
		Vertex{segment.B.X - line.A.X, segment.B.Y - line.A.Y}}
	dA := dotProd(transLine.B, transSegment.A)
	dB := dotProd(transLine.B, transSegment.B)

	return (dA >= 0 && dB <= 0) || (dA <= 0 && dB >= 0)
}

func isBoxCollision(u *Line, v *Line) bool {
	if math.Max(u.A.X, u.B.X) >= math.Min(v.A.X, v.B.X) && math.Min(u.A.X, u.B.X) <= math.Max(v.A.X, v.B.X) &&
		math.Max(u.A.Y, u.B.Y) >= math.Min(v.A.Y, v.B.Y) && math.Min(u.A.Y, u.B.Y) <= math.Max(v.A.Y, v.B.Y) {
		return true
	}

	return false
}

func dotProd(u Vertex, v Vertex) float64 {
	return u.X*v.Y - v.X*u.Y
}
