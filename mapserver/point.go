package mapserver

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

func (self *Point) IsInteger() bool {
	return math.Floor(self.X) == self.X && math.Floor(self.Y) == self.Y
}

func (self *Point) IntPoint() IntPoint {
	return IntPoint{int(self.X), int(self.Y)}
}

type IntPoint struct {
	X int
	Y int
}

type IntSize IntPoint

func (self *IntSize) IncludesPoint(p IntPoint) bool {
	return p.X >= 0 && p.X < self.X && p.Y >= 0 && p.Y < self.Y
}
