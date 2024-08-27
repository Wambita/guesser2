package calculations

func LinearRegression(y []float64) (m, c float64) {
	// y dependent variable
	// x indices, independent var
	// m = slope
	// intercept

	// generate x axis values
	x := make([]float64, len(y))
	for i := range y {
		x[i] = float64(i)
	}

	// get means
	meanX := Mean(x)
	meanY := Mean(y)

	// calculation of m and c
	// XYSum sum of products of deviations
	// X2Sum sum of squared deviations

	var stddevX, stddevY float64
	var varianceX, varianceY float64
	var X []float64
	for i := range y {
		xi := float64(i)
		X = append(X, xi)
	}
	varianceX = Variance(X)
	varianceY = Variance(y)
	stddevX = Stddev(varianceX)
	stddevY = Stddev(varianceY)

	// slope (rate at whixh y chnages for each unit of x)
	corelation := PearsonsCorrelationCoefficient(y)

	m = corelation * stddevY / stddevX

	// where the regression line crosses the y axis
	// c = y - mx
	c = meanY - m*meanX
	return
}
