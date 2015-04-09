package talib_test

import (
	"math"

	"github.com/d4l3k/talib"
)

func ExampleTrig() {
	talib.Sin([]float64{0, math.Pi / 2, math.Pi})
	// Output: [0 1 0]
}
