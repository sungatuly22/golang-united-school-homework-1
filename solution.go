package solution

import (
	"math"
)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	switch sidesNum {
	case 0:
		return math.Pi * sideLen * sideLen
	case 3:
		return math.Sqrt(3.0) / 4 * sideLen * sideLen
	case 4:
		return sideLen * sideLen
	default:
		return 0
	}
}
