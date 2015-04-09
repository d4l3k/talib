# go-talib
A Go(lang) wrapper for TA-Lib(Techinal Analysis Library) which is often used for stock/financial analysis.

http://ta-lib.org/

This wrapper is automatically generated using Ruby. It's sort of sketchy, but works fairly well.

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

## Contributing
To generate the code run:
```sh
ruby generate.rb
```

## License
Copyright (c) 2015 Tristan Rice

talib is licensed under the [MIT License](http://opensource.org/licenses/MIT).
