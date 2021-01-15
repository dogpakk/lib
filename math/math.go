package math

func IntRatio(a, b int) float64 {
	if a == 0 || b == 0 {
		return 0
	}

	return float64(a) / float64(b)
}

func IntRatioPercentage(a, b int) float64 {
	return IntRatio(a, b) * 100
}
