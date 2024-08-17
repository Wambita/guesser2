package calculations

import "math"

func Range(numslice []float64) (int, int) {
	startCalc := len(numslice) - 4
	if startCalc < 0 {
		startCalc = 0
	}

	chunk := numslice[startCalc:]

	// calculate pearson's correlation coefficient
	correlation := PearsonsCorrelationCoefficient(chunk)

	// predict the next value
	predicted := PredictNextValue(chunk)

	variance := Variance(chunk)
	stdDev := Stddev(variance)

	// adjust the range from coefficient
	confidenceMultiplier := math.Abs(correlation)
	adjustedLowRange := predicted - stdDev*(2.4-confidenceMultiplier) // Use 1.8 as a multiplier for flexibility
	adjustedHighRange := predicted + stdDev*(2.4-confidenceMultiplier)

	lowRange := int(math.Round(adjustedLowRange))
	highRange := int(math.Round(adjustedHighRange))
	return lowRange, highRange
}
