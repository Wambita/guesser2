package calculations

func PredictNextValue(y []float64) float64 {
	slope, intercept := LinearRegression(y)
	nextIndex := float64(len(y))
	// y=mx+c
	return slope*nextIndex + intercept
}
