package roguelike

type rect struct {
	X1, Y1 int
	X2, Y2 int
}

func newRectangle(x, y, width, height int) rect {
	return rect{
		X1: x,
		Y1: y,
		X2: x + width,
		Y2: y + height,
	}
}

func (r *rect) Center() (int, int) {
	centerX := (r.X1 + r.X2) / 2
	centerY := (r.Y1 + r.Y2) / 2

	return centerX, centerY
}

func (r *rect) Intersect(other rect) bool {
	return r.X1 <= other.X2 && r.X2 >= other.X1 &&
		r.Y1 <= other.Y1 && r.Y2 >= other.Y1
}
