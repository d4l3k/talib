# talib
A Go(lang) wrapper for TA-Lib

[![Build Status](https://travis-ci.org/d4l3k/talib.svg?branch=master)](https://travis-ci.org/d4l3k/talib)
[![GoDoc](https://godoc.org/github.com/d4l3k/talib?status.svg)](https://godoc.org/github.com/d4l3k/talib)

## Example

```go
package main

import (
	"fmt"
	"math"

	"github.com/d4l3k/talib"
)

func main() {
	fmt.Println(talib.Sin([]float64{0, math.Pi / 2}))
	// => [0 1]
}
```
