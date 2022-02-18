package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type customside int

const SidesTriangle customside = 3
const SidesSquare customside = 4
const SidesCircle customside = 0

func CalcSquare(sideLen float64, sidesnum customside) float64 {

	var result float64 = 0

	if sidesnum == SidesSquare {
		result = calcSquare(sideLen)
	} else if sidesnum == SidesTriangle {
		result = calcTriangle(sideLen)
	} else if sidesnum == SidesCircle {
		result = calcCircle(sideLen)
	} else {
		result = float64(0)
	}
	return result
}

func calcSquare(squareLen float64) float64 {
	square := math.Pow(squareLen, 2)
	return square
}

func calcTriangle(triangleLen float64) float64 {
	triangle := math.Sqrt(3) / 4 * triangleLen
	return triangle
}

func calcCircle(circleLen float64) float64 {
	circle := math.Pi * math.Pow(circleLen, 2)
	return circle
}
