package mapserver

import (
// "io"
)

type TileReader interface {
	ReadTile(tile IntPoint) ([]byte, error)
}

type TileWriter interface {
	WriteTile(tile IntPoint, data []byte) error
}

type Map struct {
	Bounds         Rectangle
	SampleLayout   SampleLayout
	BytesPerSample int
	NumChannels    int
	ChannelType    ChannelType
	NumTiles       IntSize
	NumTileSamples IntSize
	TileStride     int

	TileReader TileReader
	TileWriter TileWriter

	tileCache [2][2]Tile
}

func (self *Map) BytesPerTile() int {
	return self.NumTileSamples.Y * self.TileStride
}

func (self *Map) InterpolateSample(p Point, result []float64) error {
	if !self.Bounds.IncludesPoint(p) {
		return PointOutsideBounds{}
	}

	var interpolationParams SampleInterpolationParams
	self.SampleLayout.InterpolationParams(self, p, &interpolationParams)

	return nil
}

func (self *Map) cachedTile(tile, sample IntPoint) (*Tile, error) {
	if !self.NumTiles.IncludesPoint(tile) || !self.NumTileSamples.IncludesPoint(sample) {
		return nil, PointOutsideBounds{}
	}
	t := &self.tileCache[tile.X%2][tile.Y%2]
	if t.Data == nil || t.Pos != tile {
		data, err := self.TileReader.ReadTile(tile)
		if err != nil {
			return nil, err
		}
		t.Map = self
		t.Pos = tile
		t.Data = data
	}
	return t, nil
}

func (self *Map) SampleBytes(tile, sample IntPoint, result []byte) error {
	t, err := self.cachedTile(tile, sample)
	if err != nil {
		return err
	}
	self.ChannelType.SampleBytes(t, sample, result)
	return nil
}

func (self *Map) SampleFloats(tile, sample IntPoint, result []float64) error {
	t, err := self.cachedTile(tile, sample)
	if err != nil {
		return err
	}
	self.ChannelType.SampleFloats(t, sample, result)
	return nil
}
