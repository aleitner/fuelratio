package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	"math"
)

func main() {
	js.Global.Set("fuelratio", map[string]interface{}{
		"bruteForceCalculate": bruteForceCalculate,
	})
}

// todo: use a linear equations solver instead of brute forcing it.
func bruteForceCalculate(Wc, Fc, Wt, Ft, W1, W2 float64) (fuel1, fuel2, resultRatio float64) {
	fmt.Println(Wc, Fc, Wt, Ft, W1, W2)
	if Fc > Ft {
		return fuel1, fuel2, resultRatio
	}

	for i := 0.0; i <= Ft-Fc; i += 0.1 {
		f1 := i
		f2 := Ft - Fc - i
		ratio := (f1*W1 +f2*W2 + Wc*Fc) / (Fc + f1 + f2)

		if math.Abs(Wt - ratio) < math.Abs(Wt - resultRatio) {

			resultRatio = ratio
			fuel1 = i
			fuel2 = Ft - Fc - i
		}
	}

	return fuel1, fuel2, resultRatio
}