package mapserver

import (
// "github.com/ungerik/json-tree"
)

type Rectangle struct {
	Min Point
	Max Point
}

// func (self *Rectangle) InitFromJSON(tree json.Tree) {
// 	self.Min.X = tree.Select("MinX").GetFloat()
// 	self.Min.Y = tree.Select("MinY").GetFloat()
// 	self.Max.X = tree.Select("MaxX").GetFloat()
// 	self.Max.Y = tree.Select("MaxY").GetFloat()
// }

// func (self *Rectangle) WriteJSON(builder *json.Builder) {
// 	builder.Name("MinX").Value(self.Min.X)
// 	builder.Name("MinY").Value(self.Min.Y)
// 	builder.Name("MaxX").Value(self.Max.X)
// 	builder.Name("MaxY").Value(self.Max.Y)
// }

func (self *Rectangle) Width() float64 {
	return self.Max.X - self.Min.X
}

func (self *Rectangle) Height() float64 {
	return self.Max.Y - self.Min.Y
}

func (self *Rectangle) IncludesPoint(p Point) bool {
	return p.X >= self.Min.X && p.X <= self.Max.X && p.Y >= self.Min.Y && p.Y <= self.Max.Y
}

type IntRectangle struct {
	Min IntPoint
	Max IntPoint
}

func (self *IntRectangle) Width() int {
	return self.Max.X - self.Min.X
}

func (self *IntRectangle) Height() int {
	return self.Max.Y - self.Min.Y
}

func (self *IntRectangle) IncludesPoint(p IntPoint) bool {
	return p.X >= self.Min.X && p.X <= self.Max.X && p.Y >= self.Min.Y && p.Y <= self.Max.Y
}
