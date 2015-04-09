package talib_test

import (
	"fmt"
	"math"

	"github.com/d4l3k/talib"
)

func Example() {
	fmt.Println(talib.Sin([]float64{0, math.Pi / 2, math.Pi}))
	// Output: [0 1 1.2246467991473532e-16]
}
