package mapserver

import (
// "encoding/binary"
)

func isPow2(x int) bool {
	return x&(x-1) == 0
}

// func SampleToFloats(m Map, data []byte) []float64 {
// 	floats := make([]float, m.NumChannels())
// 	binary.Read(data, binary.LittleEndian, floats)
// 	return floats
// }

func InterpolateBilinear(leftTop, rightTop, leftBottom, rightBottom []float64, x, y float64, result []float64) {
	x0 := 1 - x
	y0 := 1 - y
	for i := range result {
		result[i] = leftTop[i]*x0*y0 + rightTop[i]*x*y0 + leftBottom[i]*x0*y + rightBottom[i]*x*y
	}
}

func InterpolateTrilinear(leftTop, rightTop, leftBottom, rightBottom, average []float64, x, y, z float64, result []float64) {
	x0 := 1 - x
	y0 := 1 - y
	for i := range result {
		result[i] = (leftTop[i]*x0*y0+rightTop[i]*x*y0+leftBottom[i]*x0*y+rightBottom[i]*x*y)*(1-z) + average[i]*z
	}
}
