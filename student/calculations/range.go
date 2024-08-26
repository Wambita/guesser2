package calculations

import "math"

func Range(numslice []float64) (int, int) {
	startCalc := len(numslice) - 4
	if startCalc < 0 {
		startCalc = 0
	}

	chunk := numslice[startCalc:]

	// predict the next value
	predicted := PredictNextValue(chunk)

	variance := Variance(chunk)
	stdDev := Stddev(variance)

	adjustedLowRange := predicted - stdDev*(2)
	adjustedHighRange := predicted + stdDev*(2)

	lowRange := int(math.Round(adjustedLowRange))
	highRange := int(math.Round(adjustedHighRange))
	return lowRange, highRange
}
