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
	return IntPoint{int64(self.X), int64(self.Y)}
}

type IntPoint struct {
	X int64
	Y int64
}
