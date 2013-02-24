package mapserver

import (
	"math"
)

const (
	Sint8Channel ChannelType = iota
	Uint8Channel
	Sint16Channel
	Uint16Channel
	Sint32Channel
	Uint32Channel
	Float32Channel
	InvalidChannel
)

type ChannelType int

func (self ChannelType) String() string {
	switch self {
	case Sint8Channel:
		return "Sint8Channel"
	case Uint8Channel:
		return "Uint8Channel"
	case Sint16Channel:
		return "Sint16Channel"
	case Uint16Channel:
		return "Uint16Channel"
	case Sint32Channel:
		return "Sint32Channel"
	case Uint32Channel:
		return "Uint32Channel"
	case Float32Channel:
		return "Float32Channel"
	}
	return "InvalidChannel"
}

func ChannelTypeFromString(s string) ChannelType {
	switch s {
	case "Sint8Channel":
		return Sint8Channel
	case "Uint8Channel":
		return Uint8Channel
	case "Sint16Channel":
		return Sint16Channel
	case "Uint16Channel":
		return Uint16Channel
	case "Sint32Channel":
		return Sint32Channel
	case "Uint32Channel":
		return Uint32Channel
	case "Float32Channel":
		return Float32Channel
	}
	return InvalidChannel
}

var channelTypeNumBytes = [7]int{1, 1, 2, 2, 4, 4, 4}

func (self ChannelType) NumBytes() int {
	return channelTypeNumBytes[self]
}

func (self ChannelType) SampleBytes(tile *Tile, p IntPoint, result []byte) {
	bytesPerSample := self.NumBytes() * tile.Map.NumChannels
	offset := p.Y*tile.Map.TileStride + p.X*bytesPerSample
	for i := 0; i < bytesPerSample; i++ {
		result[i] = tile.Data[i+offset]
	}
}

func (self ChannelType) SampleFloats(tile *Tile, p IntPoint, result []float64) {
	bytesPerSample := self.NumBytes() * tile.Map.NumChannels
	offset := p.Y*tile.Map.TileStride + p.X*bytesPerSample
	channelTypeSampleFloats[self](tile.Data, offset, tile.Map.NumChannels, result)
}

type channelTypeSampleFloatsFunc func(data []byte, offset, numChannels int, result []float64)

var channelTypeSampleFloats = [7]channelTypeSampleFloatsFunc{
	sampleFloatsSint8,
	sampleFloatsUint8,
	sampleFloatsSint16,
	sampleFloatsUint16,
	sampleFloatsSint32,
	sampleFloatsUint32,
	sampleFloatsFloat32,
}

func sampleFloatsSint8(data []byte, offset, numChannels int, result []float64) {
	for i := 0; i < numChannels; i++ {
		result[i] = float64(int8(data[offset+i]))
	}
}

func sampleFloatsUint8(data []byte, offset, numChannels int, result []float64) {
	for i := 0; i < numChannels; i++ {
		result[i] = float64(uint8(data[offset+i]))
	}
}

func sampleFloatsSint16(data []byte, offset, numChannels int, result []float64) {
	for i := 0; i < numChannels; i++ {
		result[i] = float64(int16(data[offset+i]) + int16(data[offset+i+1])<<8)
	}
}

func sampleFloatsUint16(data []byte, offset, numChannels int, result []float64) {
	for i := 0; i < numChannels; i++ {
		result[i] = float64(uint16(data[offset+i]) + uint16(data[offset+i+1])<<8)
	}
}

func sampleFloatsSint32(data []byte, offset, numChannels int, result []float64) {
	for i := 0; i < numChannels; i++ {
		result[i] = float64(int32(data[offset+i]) + int32(data[offset+i+1])<<8 + int32(data[offset+i+2])<<16 + int32(data[offset+i+3])<<24)
	}
}

func sampleFloatsUint32(data []byte, offset, numChannels int, result []float64) {
	for i := 0; i < numChannels; i++ {
		result[i] = float64(uint32(data[offset+i]) + uint32(data[offset+i+1])<<8 + uint32(data[offset+i+2])<<16 + uint32(data[offset+i+3])<<24)
	}
}

func sampleFloatsFloat32(data []byte, offset, numChannels int, result []float64) {
	for i := 0; i < numChannels; i++ {
		result[i] = math.Float64frombits(uint64(data[offset+i]) + uint64(data[offset+i+1])<<8 + uint64(data[offset+i+2])<<16 + uint64(data[offset+i+3])<<24)
	}
}
