package main

import (
	"github.com/gopherjs/gopherjs/js"
	"math"
)

func main() {
	js.Global.Set("fuelratio", map[string]interface{}{
		"bruteForceCalculate": bruteForceCalculate,
	})
}

// todo: use a linear equations solver instead of brute forcing it.
// Wc - Current ethanol Ratio
// Fc - Current gallons of fuel
// Wt - Optimal ethanol Ratio
// Ft - tank number of gallons
// W1 - ethanol Ratio of fuel 1
// W2 - ethanol Ratio of fuel 2
func bruteForceCalculate(Wc, Fc, Wt, Ft, W1, W2 float64) (fuel1, fuel2, resultRatio float64) {
	if Fc > Ft {
		return fuel1, fuel2, resultRatio
	}

	for i := 0.0; i <= Ft-Fc; i += 0.1 {
		f1 := i
		f2 := Ft - Fc - i
		ratio := (f1*W1 +f2*W2 + Wc*Fc) / (Fc + f1 + f2)

		if ratio == Wt {
			return f1, f2, ratio
		}

		if math.Abs(Wt - ratio) < math.Abs(Wt - resultRatio) {
			resultRatio = ratio
			fuel1 = f1
			fuel2 = f2
		}
	}

	return fuel1, fuel2, resultRatio
}