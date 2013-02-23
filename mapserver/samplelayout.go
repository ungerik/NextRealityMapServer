package mapserver

type SampleInterpolation struct {
	Tiles   IntRectangle
	Samples IntRectangle
	Weight  IntPoint
}

const (
	SampleCentered SampleLayout = iota
	SampleCenteredRepeatingBorder
	SampleOrigin
	SampleOriginRepeatingBorder
)

type SampleLayout int

func (self SampleLayout) Interpolate(m *Map, x, y float64, result *SampleInterpolation) {
	interpolate[self](m, x, y, result)
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
	panic("")
}

var interpolate = [...]func(m *Map, x, y float64, result *SampleInterpolation){
	interpolateSampleCentered,
	interpolateSampleCenteredRepeatingBorder,
	interpolateSampleOrigin,
	interpolateOriginRepeatingBorder,
}

func interpolateSampleCentered(m *Map, x, y float64, result *SampleInterpolation) {

}

func interpolateSampleCenteredRepeatingBorder(m *Map, x, y float64, result *SampleInterpolation) {

}

func interpolateSampleOrigin(m *Map, x, y float64, result *SampleInterpolation) {

}

func interpolateOriginRepeatingBorder(m *Map, x, y float64, result *SampleInterpolation) {

}
