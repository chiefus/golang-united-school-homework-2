package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
import "math"

type intCustomType int8

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	switch sidesNum {
	case SidesCircle:
		return calcCircleSquare(sideLen)
	case SidesTriangle:
		return calcTriangleSquare(sideLen)
	case SidesSquare:
		return calcSquareSquare(sideLen)
	default:
		return 0
	}
}

func calcCircleSquare(radius float64) float64 {
	return 2 * math.Pi * radius * radius
}

func calcTriangleSquare(sideLen float64) float64 {
	p := 3 * sideLen / 2
	a := p - sideLen
	return math.Sqrt(p * a * a * a)
}

func calcSquareSquare(sideLen float64) float64 {
	return sideLen * sideLen
}
