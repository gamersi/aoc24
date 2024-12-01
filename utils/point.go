package utils

// 2-D Point
type Point struct {
	X int
	Y int
}

// +Y direction
var North = Point{0, 1}

// -Y direction
var South = Point{0, -1}

// +X direction
var East = Point{1, 0}

// -X direction
var West = Point{-1, 0}

// Slice of cardinal directions
var Directions = []Point{North, South, East, West}

// Return the sum of two points
func (p Point) Add(p2 Point) Point {
	r := Point{p.X + p2.X, p.Y + p2.Y}

	return r
}

// Multiply point by a scalar
func (p Point) Scale(factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

// Turn point right 90 degrees
func (p Point) Right() Point {
	return Point{p.Y, -p.X}
}

// Turn point left 90 degrees
func (p Point) Left() Point {
	return Point{-p.Y, p.X}
}

// Return manhattan distance (abs(x)+abs(y))
func (p Point) Manhattan() int {
	return Abs(p.X) + Abs(p.Y)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
