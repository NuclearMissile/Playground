package main

func (m Matrix) Move(step Step) bool {
	if m[step.Y1][step.X1] != 0 {
		return false
	}
	m[step.Y1][step.X1], m[step.Y2][step.X2] =
		m[step.Y2][step.X2], m[step.Y1][step.X1]
	return true
}

func (np *NPuzzle) Solve() {
	return
}
