package talib_test

import (
	"fmt"
	"math"

	"github.com/d4l3k/talib"
)

func main() {
	fmt.Println(talib.Sin([]float64{0, math.Pi / 2}))
	// => [0 1]
}
