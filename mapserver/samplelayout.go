package mapserver

import (
// "fmt"
)

type SampleInterpolationParams struct {
	Tiles   IntRectangle
	Samples IntRectangle
	Weight  IntPoint
}

const (
	SampleCentered SampleLayout = iota
	SampleCenteredRepeatingBorder
	SampleOrigin
	SampleOriginRepeatingBorder
	SampleLayoutInvalid
)

type SampleLayout int

func (self SampleLayout) InterpolationParams(m *Map, p Point, result *SampleInterpolationParams) {
	sampleLayoutInterpolationParams[self](m, p, result)
}

func (self SampleLayout) String() string {
	switch self {
	case SampleCentered:
		return "SampleCentered"
	case SampleCenteredRepeatingBorder:
		return "SampleCenteredRepeatingBorder"
	case SampleOrigin:
		return "SampleOrigin"
	case SampleOriginRepeatingBorder:
		return "SampleOriginRepeatingBorder"
	}
	return "SampleLayoutInvalid"
}

func SampleLayoutFromString(s string) SampleLayout {
	switch s {
	case "SampleCentered":
		return SampleCentered
	case "SampleCenteredRepeatingBorder":
		return SampleCenteredRepeatingBorder
	case "SampleOrigin":
		return SampleOrigin
	case "SampleOriginRepeatingBorder":
		return SampleOriginRepeatingBorder
	}
	return SampleLayoutInvalid
}

var sampleLayoutInterpolationParams = [4]func(m *Map, p Point, result *SampleInterpolationParams){
	interpolateSampleCentered,
	interpolateSampleCenteredRepeatingBorder,
	interpolateSampleOrigin,
	interpolateOriginRepeatingBorder,
}

func interpolateSampleCentered(m *Map, p Point, result *SampleInterpolationParams) {

}

func interpolateSampleCenteredRepeatingBorder(m *Map, p Point, result *SampleInterpolationParams) {

}

func interpolateSampleOrigin(m *Map, p Point, result *SampleInterpolationParams) {

}

func interpolateOriginRepeatingBorder(m *Map, p Point, result *SampleInterpolationParams) {

}
