package mapserver

import (
// "io"
)

// type TileReader interface {
// 	ReadTile(x, y int) ([]byte, error)
// }

// type TileWriter interface {
// 	WriteTile(x, y int, data []byte) error
// }

type Map struct {
	Bounds         Rectangle
	HorTiles       int
	VerTiles       int
	SampleLayout   SampleLayout
	SampleSize     int
	Channels       int
	HorTileSamples int
	VerTileSamples int
	TileStride     int

	tileCache [4][]byte
}

func (self *Map) SampleBytes(x, y float64) []byte {
	panic("")
}

func (self *Map) SampleFloats(x, y float64, result []float64) {
	panic("")
}

func (self *Map) TileSize() int {
	return self.VerTileSamples * self.TileStride
}
