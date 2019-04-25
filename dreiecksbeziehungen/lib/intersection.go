package lib

import "math"

func LineSegementIntersection(u *Line, v *Line) bool {
	if isBoxCollision(u, v) {
		return isSegmentLineCollision(u, v) && isSegmentLineCollision(v, u)
	}

	return false
}

func isSegmentLineCollision(line *Line, segment *Line) bool {
	transLine := Line{Vector{}, Vector{line.B.X - line.A.X, line.B.Y - line.A.Y}}
	transSegment := Line{Vector{segment.A.X - line.A.X, segment.A.Y - line.A.Y},
		Vector{segment.B.X - line.A.X, segment.B.Y - line.A.Y}}
	//dA := crossProdComp3(transLine.B, transSegment.A)
	//dB := crossProdComp3(transLine.B, transSegment.B)

	//return (dA >= 0 && dB <= 0) || (dA <= 0 && dB >= 0)

	return leftFromLine(transLine.B, transSegment.A) != leftFromLine(transLine.B, transSegment.B)
}

func isBoxCollision(u *Line, v *Line) bool {
	if math.Max(u.A.X, u.B.X) >= math.Min(v.A.X, v.B.X) && math.Min(u.A.X, u.B.X) <= math.Max(v.A.X, v.B.X) &&
		math.Max(u.A.Y, u.B.Y) >= math.Min(v.A.Y, v.B.Y) && math.Min(u.A.Y, u.B.Y) <= math.Max(v.A.Y, v.B.Y) {
		return true
	}

	return false
}

func leftFromLine(v Vector, p Vector) bool {
	return v.X*p.Y-p.X*v.Y > 0
}

func rightFromLine(v Vector, p Vector) bool {
	return v.X*p.Y-p.X*v.Y < 0
}
