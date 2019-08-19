package main

// MaxBall ...
func MaxBall(v0 int) (rc int) {
	for i := 0.0; ; i++ {
		if getHeight(float64(v0), i) > getHeight(float64(v0), i+1) {
			rc = int(i)
			break
		}
	}
	return
}

func getHeight(v, t float64) float64 {
	const gravity = 9.81
	var (
		km            float64 = 1000
		hour          float64 = 3600
		tenthOfSecond float64 = 10
	)
	v = v * km / hour
	t = t / tenthOfSecond
	return v*t - 0.5*gravity*t*t
}
