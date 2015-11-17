# go-talib
A Go(lang) wrapper for TA-Lib(Techinal Analysis Library) which is often used for stock/financial analysis.

http://ta-lib.org/

[![Build Status](https://travis-ci.org/d4l3k/talib.svg?branch=master)](https://travis-ci.org/d4l3k/talib)
[![GoDoc](https://godoc.org/github.com/d4l3k/talib?status.svg)](https://godoc.org/github.com/d4l3k/talib)

To use the library you need TA-Lib installed.

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

## Installing

Install the dependencies then run

```
$ go install github.com/d4l3k/talib
```

### Dependencies

To use TA-Lib for python, you need to have the
[TA-Lib](http://ta-lib.org/hdr_dw.html) already installed:

##### Mac OS X

```
$ brew install ta-lib
```

##### Windows

Download [ta-lib-0.4.0-msvc.zip](http://prdownloads.sourceforge.net/ta-lib/ta-lib-0.4.0-msvc.zip)
and unzip to ``C:\ta-lib``

##### Linux

Install from your package manager or install from source.

On Arch Linux `ta-lib` is available from the AUR.

To compile first download [ta-lib-0.4.0-src.tar.gz](http://prdownloads.sourceforge.net/ta-lib/ta-lib-0.4.0-src.tar.gz) and:
```
$ untar and cd
$ ./configure --prefix=/usr LDFLAGS="-lm"
$ make
$ sudo make install
```

> If you build ``TA-Lib`` using ``make -jX`` it will fail but that's OK!
> Simply rerun ``make -jX`` followed by ``[sudo] make install``.

## Contributing
This wrapper is automatically generated using Ruby. It's sort of sketchy, but works fairly well.

To generate the code run:
```sh
$ ruby generate.rb
```

## License
Copyright (c) 2015 [Tristan Rice](https://fn.lc)

talib is licensed under the [MIT License](http://opensource.org/licenses/MIT).
