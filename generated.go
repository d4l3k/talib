package talib

// #cgo LDFLAGS: -lta_lib
// #include "ta-lib/ta_libc.h"
import "C"

import (
	"fmt"
	"unsafe"
)

func init() {
	n, err := C.TA_Initialize()
	if n != 0 {
		panic(fmt.Sprintf("ta-lib status is %d %s", n, err))
	}
}

/*Acos - Vector Trigonometric ACos

Input = double

Output = double

*/
func Acos(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_ACOS(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Ad - Chaikin A/D Line

Input = High, Low, Close, Volume

Output = double

*/
func Ad(high, low, close, volume []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_AD(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), (*C.double)(unsafe.Pointer(&volume[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Add - Vector Arithmetic Add

Input = double, double

Output = double

*/
func Add(real0, real1 []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real0))
	C.TA_ADD(0, C.int(len(real0)-1), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Adosc - Chaikin A/D Oscillator

Input = High, Low, Close, Volume

Output = double

Optional Parameters

-------------------

optInFastPeriod:(From 2 to 100000)

Number of period for the fast MA

optInSlowPeriod:(From 2 to 100000)

Number of period for the slow MA

*/
func Adosc(high, low, close, volume []float64, fastPeriod, slowPeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_ADOSC(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), (*C.double)(unsafe.Pointer(&volume[0])), C.int(fastPeriod), C.int(slowPeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Adx - Average Directional Movement Index

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Adx(high, low, close []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_ADX(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Adxr - Average Directional Movement Index Rating

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Adxr(high, low, close []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_ADXR(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Apo - Absolute Price Oscillator

Input = double

Output = double

Optional Parameters

-------------------

optInFastPeriod:(From 2 to 100000)

Number of period for the fast MA

optInSlowPeriod:(From 2 to 100000)

Number of period for the slow MA

optInMAType:

Type of Moving Average

*/
func Apo(real []float64, fastPeriod, slowPeriod, mAType int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_APO(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(fastPeriod), C.int(slowPeriod), C.TA_MAType(mAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Aroon - Aroon

Input = High, Low

Output = double, double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Aroon(high, low []float64, timePeriod int) ([]float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	outAroonDown := make([]float64, len(high))
	outAroonUp := make([]float64, len(high))
	C.TA_AROON(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outAroonDown[0])), (*C.double)(unsafe.Pointer(&outAroonUp[0])))
	return outAroonDown, outAroonUp
}

/*Aroonosc - Aroon Oscillator

Input = High, Low

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Aroonosc(high, low []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_AROONOSC(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Asin - Vector Trigonometric ASin

Input = double

Output = double

*/
func Asin(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_ASIN(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Atan - Vector Trigonometric ATan

Input = double

Output = double

*/
func Atan(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_ATAN(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Atr - Average True Range

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Atr(high, low, close []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_ATR(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Avgprice - Average Price

Input = Open, High, Low, Close

Output = double

*/
func Avgprice(open, high, low, close []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(open))
	C.TA_AVGPRICE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Bbands - Bollinger Bands

Input = double

Output = double, double, double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

optInNbDevUp:(From TA_REAL_MIN to TA_REAL_MAX)

Deviation multiplier for upper band

optInNbDevDn:(From TA_REAL_MIN to TA_REAL_MAX)

Deviation multiplier for lower band

optInMAType:

Type of Moving Average

*/
func Bbands(real []float64, timePeriod int, nbDevUp, nbDevDn float64, mAType int) ([]float64, []float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	outRealUpperBand := make([]float64, len(real))
	outRealMiddleBand := make([]float64, len(real))
	outRealLowerBand := make([]float64, len(real))
	C.TA_BBANDS(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), C.double(nbDevUp), C.double(nbDevDn), C.TA_MAType(mAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outRealUpperBand[0])), (*C.double)(unsafe.Pointer(&outRealMiddleBand[0])), (*C.double)(unsafe.Pointer(&outRealLowerBand[0])))
	return outRealUpperBand, outRealMiddleBand, outRealLowerBand
}

/*Beta - Beta

Input = double, double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Beta(real0, real1 []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real0))
	C.TA_BETA(0, C.int(len(real0)-1), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Bop - Balance Of Power

Input = Open, High, Low, Close

Output = double

*/
func Bop(open, high, low, close []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(open))
	C.TA_BOP(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Cci - Commodity Channel Index

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Cci(high, low, close []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_CCI(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Cdl2crows - Two Crows

Input = Open, High, Low, Close

Output = int

*/
func Cdl2crows(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDL2CROWS(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdl3blackcrows - Three Black Crows

Input = Open, High, Low, Close

Output = int

*/
func Cdl3blackcrows(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDL3BLACKCROWS(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdl3inside - Three Inside Up/Down

Input = Open, High, Low, Close

Output = int

*/
func Cdl3inside(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDL3INSIDE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdl3linestrike - Three-Line Strike

Input = Open, High, Low, Close

Output = int

*/
func Cdl3linestrike(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDL3LINESTRIKE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdl3outside - Three Outside Up/Down

Input = Open, High, Low, Close

Output = int

*/
func Cdl3outside(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDL3OUTSIDE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdl3starsinsouth - Three Stars In The South

Input = Open, High, Low, Close

Output = int

*/
func Cdl3starsinsouth(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDL3STARSINSOUTH(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdl3whitesoldiers - Three Advancing White Soldiers

Input = Open, High, Low, Close

Output = int

*/
func Cdl3whitesoldiers(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDL3WHITESOLDIERS(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlabandonedbaby - Abandoned Baby

Input = Open, High, Low, Close

Output = int

Optional Parameters

-------------------

optInPenetration:(From 0 to TA_REAL_MAX)

Percentage of penetration of a candle within another candle

*/
func Cdlabandonedbaby(open, high, low, close []float64, penetration float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLABANDONEDBABY(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.double(penetration), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdladvanceblock - Advance Block

Input = Open, High, Low, Close

Output = int

*/
func Cdladvanceblock(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLADVANCEBLOCK(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlbelthold - Belt-hold

Input = Open, High, Low, Close

Output = int

*/
func Cdlbelthold(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLBELTHOLD(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlbreakaway - Breakaway

Input = Open, High, Low, Close

Output = int

*/
func Cdlbreakaway(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLBREAKAWAY(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlclosingmarubozu - Closing Marubozu

Input = Open, High, Low, Close

Output = int

*/
func Cdlclosingmarubozu(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLCLOSINGMARUBOZU(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlconcealbabyswall - Concealing Baby Swallow

Input = Open, High, Low, Close

Output = int

*/
func Cdlconcealbabyswall(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLCONCEALBABYSWALL(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlcounterattack - Counterattack

Input = Open, High, Low, Close

Output = int

*/
func Cdlcounterattack(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLCOUNTERATTACK(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdldarkcloudcover - Dark Cloud Cover

Input = Open, High, Low, Close

Output = int

Optional Parameters

-------------------

optInPenetration:(From 0 to TA_REAL_MAX)

Percentage of penetration of a candle within another candle

*/
func Cdldarkcloudcover(open, high, low, close []float64, penetration float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLDARKCLOUDCOVER(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.double(penetration), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdldoji - Doji

Input = Open, High, Low, Close

Output = int

*/
func Cdldoji(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLDOJI(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdldojistar - Doji Star

Input = Open, High, Low, Close

Output = int

*/
func Cdldojistar(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLDOJISTAR(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdldragonflydoji - Dragonfly Doji

Input = Open, High, Low, Close

Output = int

*/
func Cdldragonflydoji(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLDRAGONFLYDOJI(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlengulfing - Engulfing Pattern

Input = Open, High, Low, Close

Output = int

*/
func Cdlengulfing(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLENGULFING(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdleveningdojistar - Evening Doji Star

Input = Open, High, Low, Close

Output = int

Optional Parameters

-------------------

optInPenetration:(From 0 to TA_REAL_MAX)

Percentage of penetration of a candle within another candle

*/
func Cdleveningdojistar(open, high, low, close []float64, penetration float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLEVENINGDOJISTAR(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.double(penetration), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdleveningstar - Evening Star

Input = Open, High, Low, Close

Output = int

Optional Parameters

-------------------

optInPenetration:(From 0 to TA_REAL_MAX)

Percentage of penetration of a candle within another candle

*/
func Cdleveningstar(open, high, low, close []float64, penetration float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLEVENINGSTAR(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.double(penetration), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlgapsidesidewhite - Up/Down-gap side-by-side white lines

Input = Open, High, Low, Close

Output = int

*/
func Cdlgapsidesidewhite(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLGAPSIDESIDEWHITE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlgravestonedoji - Gravestone Doji

Input = Open, High, Low, Close

Output = int

*/
func Cdlgravestonedoji(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLGRAVESTONEDOJI(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlhammer - Hammer

Input = Open, High, Low, Close

Output = int

*/
func Cdlhammer(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLHAMMER(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlhangingman - Hanging Man

Input = Open, High, Low, Close

Output = int

*/
func Cdlhangingman(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLHANGINGMAN(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlharami - Harami Pattern

Input = Open, High, Low, Close

Output = int

*/
func Cdlharami(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLHARAMI(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlharamicross - Harami Cross Pattern

Input = Open, High, Low, Close

Output = int

*/
func Cdlharamicross(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLHARAMICROSS(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlhighwave - High-Wave Candle

Input = Open, High, Low, Close

Output = int

*/
func Cdlhighwave(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLHIGHWAVE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlhikkake - Hikkake Pattern

Input = Open, High, Low, Close

Output = int

*/
func Cdlhikkake(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLHIKKAKE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlhikkakemod - Modified Hikkake Pattern

Input = Open, High, Low, Close

Output = int

*/
func Cdlhikkakemod(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLHIKKAKEMOD(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlhomingpigeon - Homing Pigeon

Input = Open, High, Low, Close

Output = int

*/
func Cdlhomingpigeon(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLHOMINGPIGEON(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlidentical3crows - Identical Three Crows

Input = Open, High, Low, Close

Output = int

*/
func Cdlidentical3crows(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLIDENTICAL3CROWS(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlinneck - In-Neck Pattern

Input = Open, High, Low, Close

Output = int

*/
func Cdlinneck(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLINNECK(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlinvertedhammer - Inverted Hammer

Input = Open, High, Low, Close

Output = int

*/
func Cdlinvertedhammer(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLINVERTEDHAMMER(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlkicking - Kicking

Input = Open, High, Low, Close

Output = int

*/
func Cdlkicking(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLKICKING(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlkickingbylength - Kicking - bull/bear determined by the longer marubozu

Input = Open, High, Low, Close

Output = int

*/
func Cdlkickingbylength(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLKICKINGBYLENGTH(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlladderbottom - Ladder Bottom

Input = Open, High, Low, Close

Output = int

*/
func Cdlladderbottom(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLLADDERBOTTOM(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdllongleggeddoji - Long Legged Doji

Input = Open, High, Low, Close

Output = int

*/
func Cdllongleggeddoji(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLLONGLEGGEDDOJI(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdllongline - Long Line Candle

Input = Open, High, Low, Close

Output = int

*/
func Cdllongline(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLLONGLINE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlmarubozu - Marubozu

Input = Open, High, Low, Close

Output = int

*/
func Cdlmarubozu(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLMARUBOZU(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlmatchinglow - Matching Low

Input = Open, High, Low, Close

Output = int

*/
func Cdlmatchinglow(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLMATCHINGLOW(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlmathold - Mat Hold

Input = Open, High, Low, Close

Output = int

Optional Parameters

-------------------

optInPenetration:(From 0 to TA_REAL_MAX)

Percentage of penetration of a candle within another candle

*/
func Cdlmathold(open, high, low, close []float64, penetration float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLMATHOLD(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.double(penetration), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlmorningdojistar - Morning Doji Star

Input = Open, High, Low, Close

Output = int

Optional Parameters

-------------------

optInPenetration:(From 0 to TA_REAL_MAX)

Percentage of penetration of a candle within another candle

*/
func Cdlmorningdojistar(open, high, low, close []float64, penetration float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLMORNINGDOJISTAR(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.double(penetration), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlmorningstar - Morning Star

Input = Open, High, Low, Close

Output = int

Optional Parameters

-------------------

optInPenetration:(From 0 to TA_REAL_MAX)

Percentage of penetration of a candle within another candle

*/
func Cdlmorningstar(open, high, low, close []float64, penetration float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLMORNINGSTAR(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.double(penetration), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlonneck - On-Neck Pattern

Input = Open, High, Low, Close

Output = int

*/
func Cdlonneck(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLONNECK(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlpiercing - Piercing Pattern

Input = Open, High, Low, Close

Output = int

*/
func Cdlpiercing(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLPIERCING(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlrickshawman - Rickshaw Man

Input = Open, High, Low, Close

Output = int

*/
func Cdlrickshawman(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLRICKSHAWMAN(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlrisefall3methods - Rising/Falling Three Methods

Input = Open, High, Low, Close

Output = int

*/
func Cdlrisefall3methods(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLRISEFALL3METHODS(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlseparatinglines - Separating Lines

Input = Open, High, Low, Close

Output = int

*/
func Cdlseparatinglines(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLSEPARATINGLINES(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlshootingstar - Shooting Star

Input = Open, High, Low, Close

Output = int

*/
func Cdlshootingstar(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLSHOOTINGSTAR(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlshortline - Short Line Candle

Input = Open, High, Low, Close

Output = int

*/
func Cdlshortline(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLSHORTLINE(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlspinningtop - Spinning Top

Input = Open, High, Low, Close

Output = int

*/
func Cdlspinningtop(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLSPINNINGTOP(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlstalledpattern - Stalled Pattern

Input = Open, High, Low, Close

Output = int

*/
func Cdlstalledpattern(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLSTALLEDPATTERN(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlsticksandwich - Stick Sandwich

Input = Open, High, Low, Close

Output = int

*/
func Cdlsticksandwich(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLSTICKSANDWICH(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdltakuri - Takuri (Dragonfly Doji with very long lower shadow)

Input = Open, High, Low, Close

Output = int

*/
func Cdltakuri(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLTAKURI(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdltasukigap - Tasuki Gap

Input = Open, High, Low, Close

Output = int

*/
func Cdltasukigap(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLTASUKIGAP(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlthrusting - Thrusting Pattern

Input = Open, High, Low, Close

Output = int

*/
func Cdlthrusting(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLTHRUSTING(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdltristar - Tristar Pattern

Input = Open, High, Low, Close

Output = int

*/
func Cdltristar(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLTRISTAR(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlunique3river - Unique 3 River

Input = Open, High, Low, Close

Output = int

*/
func Cdlunique3river(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLUNIQUE3RIVER(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlupsidegap2crows - Upside Gap Two Crows

Input = Open, High, Low, Close

Output = int

*/
func Cdlupsidegap2crows(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLUPSIDEGAP2CROWS(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Cdlxsidegap3methods - Upside/Downside Gap Three Methods

Input = Open, High, Low, Close

Output = int

*/
func Cdlxsidegap3methods(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(open))
	C.TA_CDLXSIDEGAP3METHODS(0, C.int(len(open)-1), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Ceil - Vector Ceil

Input = double

Output = double

*/
func Ceil(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_CEIL(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Cmo - Chande Momentum Oscillator

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Cmo(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_CMO(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Correl - Pearson's Correlation Coefficient (r)

Input = double, double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Correl(real0, real1 []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real0))
	C.TA_CORREL(0, C.int(len(real0)-1), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Cos - Vector Trigonometric Cos

Input = double

Output = double

*/
func Cos(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_COS(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Cosh - Vector Trigonometric Cosh

Input = double

Output = double

*/
func Cosh(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_COSH(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Dema - Double Exponential Moving Average

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Dema(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_DEMA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Div - Vector Arithmetic Div

Input = double, double

Output = double

*/
func Div(real0, real1 []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real0))
	C.TA_DIV(0, C.int(len(real0)-1), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Dx - Directional Movement Index

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Dx(high, low, close []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_DX(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Ema - Exponential Moving Average

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Ema(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_EMA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Exp - Vector Arithmetic Exp

Input = double

Output = double

*/
func Exp(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_EXP(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Floor - Vector Floor

Input = double

Output = double

*/
func Floor(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_FLOOR(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*HtDcperiod - Hilbert Transform - Dominant Cycle Period

Input = double

Output = double

*/
func HtDcperiod(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_HT_DCPERIOD(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*HtDcphase - Hilbert Transform - Dominant Cycle Phase

Input = double

Output = double

*/
func HtDcphase(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_HT_DCPHASE(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*HtPhasor - Hilbert Transform - Phasor Components

Input = double

Output = double, double

*/
func HtPhasor(real []float64) ([]float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	outInPhase := make([]float64, len(real))
	outQuadrature := make([]float64, len(real))
	C.TA_HT_PHASOR(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outInPhase[0])), (*C.double)(unsafe.Pointer(&outQuadrature[0])))
	return outInPhase, outQuadrature
}

/*HtSine - Hilbert Transform - SineWave

Input = double

Output = double, double

*/
func HtSine(real []float64) ([]float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	outSine := make([]float64, len(real))
	outLeadSine := make([]float64, len(real))
	C.TA_HT_SINE(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outSine[0])), (*C.double)(unsafe.Pointer(&outLeadSine[0])))
	return outSine, outLeadSine
}

/*HtTrendline - Hilbert Transform - Instantaneous Trendline

Input = double

Output = double

*/
func HtTrendline(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_HT_TRENDLINE(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*HtTrendmode - Hilbert Transform - Trend vs Cycle Mode

Input = double

Output = int

*/
func HtTrendmode(real []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(real))
	C.TA_HT_TRENDMODE(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Kama - Kaufman Adaptive Moving Average

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Kama(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_KAMA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Linearreg - Linear Regression

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Linearreg(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_LINEARREG(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*LinearregAngle - Linear Regression Angle

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func LinearregAngle(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_LINEARREG_ANGLE(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*LinearregIntercept - Linear Regression Intercept

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func LinearregIntercept(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_LINEARREG_INTERCEPT(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*LinearregSlope - Linear Regression Slope

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func LinearregSlope(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_LINEARREG_SLOPE(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Ln - Vector Log Natural

Input = double

Output = double

*/
func Ln(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_LN(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Log10 - Vector Log10

Input = double

Output = double

*/
func Log10(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_LOG10(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Ma - Moving average

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

optInMaType:

Type of Moving Average

*/
func Ma(real []float64, timePeriod, mAType int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_MA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), C.TA_MAType(mAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Macd - Moving Average Convergence/Divergence

Input = double

Output = double, double, double

Optional Parameters

-------------------

optInFastPeriod:(From 2 to 100000)

Number of period for the fast MA

optInSlowPeriod:(From 2 to 100000)

Number of period for the slow MA

optInSignalPeriod:(From 1 to 100000)

Smoothing for the signal line (nb of period)

*/
func Macd(real []float64, fastPeriod, slowPeriod, signalPeriod int) ([]float64, []float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	outMACD := make([]float64, len(real))
	outMACDSignal := make([]float64, len(real))
	outMACDHist := make([]float64, len(real))
	C.TA_MACD(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(fastPeriod), C.int(slowPeriod), C.int(signalPeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outMACD[0])), (*C.double)(unsafe.Pointer(&outMACDSignal[0])), (*C.double)(unsafe.Pointer(&outMACDHist[0])))
	return outMACD, outMACDSignal, outMACDHist
}

/*Macdext - MACD with controllable MA type

Input = double

Output = double, double, double

Optional Parameters

-------------------

optInFastPeriod:(From 2 to 100000)

Number of period for the fast MA

optInFastMAType:

Type of Moving Average for fast MA

optInSlowPeriod:(From 2 to 100000)

Number of period for the slow MA

optInSlowMAType:

Type of Moving Average for slow MA

optInSignalPeriod:(From 1 to 100000)

Smoothing for the signal line (nb of period)

optInSignalMAType:

Type of Moving Average for signal line

*/
func Macdext(real []float64, fastPeriod, fastMAType, slowPeriod, slowMAType, signalPeriod, signalMAType int) ([]float64, []float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	outMACD := make([]float64, len(real))
	outMACDSignal := make([]float64, len(real))
	outMACDHist := make([]float64, len(real))
	C.TA_MACDEXT(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(fastPeriod), C.TA_MAType(fastMAType), C.int(slowPeriod), C.TA_MAType(slowMAType), C.int(signalPeriod), C.TA_MAType(signalMAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outMACD[0])), (*C.double)(unsafe.Pointer(&outMACDSignal[0])), (*C.double)(unsafe.Pointer(&outMACDHist[0])))
	return outMACD, outMACDSignal, outMACDHist
}

/*Macdfix - Moving Average Convergence/Divergence Fix 12/26

Input = double

Output = double, double, double

Optional Parameters

-------------------

optInSignalPeriod:(From 1 to 100000)

Smoothing for the signal line (nb of period)

*/
func Macdfix(real []float64, signalPeriod int) ([]float64, []float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	outMACD := make([]float64, len(real))
	outMACDSignal := make([]float64, len(real))
	outMACDHist := make([]float64, len(real))
	C.TA_MACDFIX(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(signalPeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outMACD[0])), (*C.double)(unsafe.Pointer(&outMACDSignal[0])), (*C.double)(unsafe.Pointer(&outMACDHist[0])))
	return outMACD, outMACDSignal, outMACDHist
}

/*Mama - MESA Adaptive Moving Average

Input = double

Output = double, double

Optional Parameters

-------------------

optInFastLimit:(From 0.01 to 0.99)

Upper limit use in the adaptive algorithm

optInSlowLimit:(From 0.01 to 0.99)

Lower limit use in the adaptive algorithm

*/
func Mama(real []float64, fastLimit, slowLimit float64) ([]float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	outMAMA := make([]float64, len(real))
	outFAMA := make([]float64, len(real))
	C.TA_MAMA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.double(fastLimit), C.double(slowLimit), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outMAMA[0])), (*C.double)(unsafe.Pointer(&outFAMA[0])))
	return outMAMA, outFAMA
}

/*Mavp - Moving average with variable period

Input = double, double

Output = double

Optional Parameters

-------------------

optInMinPeriod:(From 2 to 100000)

Value less than minimum will be changed to Minimum period

optInMaxPeriod:(From 2 to 100000)

Value higher than maximum will be changed to Maximum period

optInMAType:

Type of Moving Average

*/
func Mavp(real, periods []float64, minPeriod, maxPeriod, mAType int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_MAVP(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), (*C.double)(unsafe.Pointer(&periods[0])), C.int(minPeriod), C.int(maxPeriod), C.TA_MAType(mAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Max - Highest value over a specified period

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Max(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_MAX(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Maxindex - Index of highest value over a specified period

Input = double

Output = int

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Maxindex(real []float64, timePeriod int) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(real))
	C.TA_MAXINDEX(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Medprice - Median Price

Input = High, Low

Output = double

*/
func Medprice(high, low []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_MEDPRICE(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Mfi - Money Flow Index

Input = High, Low, Close, Volume

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Mfi(high, low, close, volume []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_MFI(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), (*C.double)(unsafe.Pointer(&volume[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Midpoint - MidPoint over period

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Midpoint(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_MIDPOINT(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Midprice - Midpoint Price over period

Input = High, Low

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Midprice(high, low []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_MIDPRICE(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Min - Lowest value over a specified period

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Min(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_MIN(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Minindex - Index of lowest value over a specified period

Input = double

Output = int

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Minindex(real []float64, timePeriod int) []int {
	var outBegIdx C.int
	var outNBElement C.int
	outInteger := make([]int, len(real))
	C.TA_MININDEX(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outInteger[0])))
	return outInteger
}

/*Minmax - Lowest and highest values over a specified period

Input = double

Output = double, double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Minmax(real []float64, timePeriod int) ([]float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	outMin := make([]float64, len(real))
	outMax := make([]float64, len(real))
	C.TA_MINMAX(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outMin[0])), (*C.double)(unsafe.Pointer(&outMax[0])))
	return outMin, outMax
}

/*Minmaxindex - Indexes of lowest and highest values over a specified period

Input = double

Output = int, int

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Minmaxindex(real []float64, timePeriod int) ([]int, []int) {
	var outBegIdx C.int
	var outNBElement C.int
	outMinIdx := make([]int, len(real))
	outMaxIdx := make([]int, len(real))
	C.TA_MINMAXINDEX(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&outMinIdx[0])), (*C.int)(unsafe.Pointer(&outMaxIdx[0])))
	return outMinIdx, outMaxIdx
}

/*MinusDi - Minus Directional Indicator

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func MinusDi(high, low, close []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_MINUS_DI(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*MinusDm - Minus Directional Movement

Input = High, Low

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func MinusDm(high, low []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_MINUS_DM(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Mom - Momentum

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Mom(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_MOM(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Mult - Vector Arithmetic Mult

Input = double, double

Output = double

*/
func Mult(real0, real1 []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real0))
	C.TA_MULT(0, C.int(len(real0)-1), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Natr - Normalized Average True Range

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Natr(high, low, close []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_NATR(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Obv - On Balance Volume

Input = double, Volume

Output = double

*/
func Obv(real, volume []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_OBV(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), (*C.double)(unsafe.Pointer(&volume[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*PlusDi - Plus Directional Indicator

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func PlusDi(high, low, close []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_PLUS_DI(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*PlusDm - Plus Directional Movement

Input = High, Low

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func PlusDm(high, low []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_PLUS_DM(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Ppo - Percentage Price Oscillator

Input = double

Output = double

Optional Parameters

-------------------

optInFastPeriod:(From 2 to 100000)

Number of period for the fast MA

optInSlowPeriod:(From 2 to 100000)

Number of period for the slow MA

optInMAType:

Type of Moving Average

*/
func Ppo(real []float64, fastPeriod, slowPeriod, mAType int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_PPO(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(fastPeriod), C.int(slowPeriod), C.TA_MAType(mAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Roc - Rate of change : ((price/prevPrice)-1)*100

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Roc(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_ROC(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Rocp - Rate of change Percentage: (price-prevPrice)/prevPrice

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Rocp(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_ROCP(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Rocr - Rate of change ratio: (price/prevPrice)

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Rocr(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_ROCR(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Rocr100 - Rate of change ratio 100 scale: (price/prevPrice)*100

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Rocr100(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_ROCR100(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Rsi - Relative Strength Index

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Rsi(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_RSI(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Sar - Parabolic Sar

Input = High, Low

Output = double

Optional Parameters

-------------------

optInAcceleration:(From 0 to TA_REAL_MAX)

Acceleration Factor used up to the Maximum value

optInMaximum:(From 0 to TA_REAL_MAX)

Acceleration Factor Maximum value

*/
func Sar(high, low []float64, acceleration, maximum float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_SAR(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), C.double(acceleration), C.double(maximum), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Sarext - Parabolic SAR - Extended

Input = High, Low

Output = double

Optional Parameters

-------------------

optInStartValue:(From TA_REAL_MIN to TA_REAL_MAX)

Start value and direction. 0 for Auto, >0 for Long, <0 for Short

optInOffsetOnReverse:(From 0 to TA_REAL_MAX)

Percent offset added/removed to initial stop on short/long reversal

optInAccelerationInitLong:(From 0 to TA_REAL_MAX)

Acceleration Factor initial value for the Long direction

optInAccelerationLong:(From 0 to TA_REAL_MAX)

Acceleration Factor for the Long direction

optInAccelerationMaxLong:(From 0 to TA_REAL_MAX)

Acceleration Factor maximum value for the Long direction

optInAccelerationInitShort:(From 0 to TA_REAL_MAX)

Acceleration Factor initial value for the Short direction

optInAccelerationShort:(From 0 to TA_REAL_MAX)

Acceleration Factor for the Short direction

optInAccelerationMaxShort:(From 0 to TA_REAL_MAX)

Acceleration Factor maximum value for the Short direction

*/
func Sarext(high, low []float64, startValue, offsetOnReverse, accelerationInitLong, accelerationLong, accelerationMaxLong, accelerationInitShort, accelerationShort, accelerationMaxShort float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_SAREXT(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), C.double(startValue), C.double(offsetOnReverse), C.double(accelerationInitLong), C.double(accelerationLong), C.double(accelerationMaxLong), C.double(accelerationInitShort), C.double(accelerationShort), C.double(accelerationMaxShort), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Sin - Vector Trigonometric Sin

Input = double

Output = double

*/
func Sin(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_SIN(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Sinh - Vector Trigonometric Sinh

Input = double

Output = double

*/
func Sinh(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_SINH(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Sma - Simple Moving Average

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Sma(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_SMA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Sqrt - Vector Square Root

Input = double

Output = double

*/
func Sqrt(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_SQRT(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Stddev - Standard Deviation

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

optInNbDev:(From TA_REAL_MIN to TA_REAL_MAX)

Nb of deviations

*/
func Stddev(real []float64, timePeriod int, nbDev float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_STDDEV(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), C.double(nbDev), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Stoch - Stochastic

Input = High, Low, Close

Output = double, double

Optional Parameters

-------------------

optInFastK_Period:(From 1 to 100000)

Time period for building the Fast-K line

optInSlowK_Period:(From 1 to 100000)

Smoothing for making the Slow-K line. Usually set to 3

optInSlowK_MAType:

Type of Moving Average for Slow-K

optInSlowD_Period:(From 1 to 100000)

Smoothing for making the Slow-D line

optInSlowD_MAType:

Type of Moving Average for Slow-D

*/
func Stoch(high, low, close []float64, fastKPeriod, slowKPeriod, slowKMAType, slowDPeriod, slowDMAType int) ([]float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	outSlowK := make([]float64, len(high))
	outSlowD := make([]float64, len(high))
	C.TA_STOCH(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(fastKPeriod), C.int(slowKPeriod), C.TA_MAType(slowKMAType), C.int(slowDPeriod), C.TA_MAType(slowDMAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outSlowK[0])), (*C.double)(unsafe.Pointer(&outSlowD[0])))
	return outSlowK, outSlowD
}

/*Stochf - Stochastic Fast

Input = High, Low, Close

Output = double, double

Optional Parameters

-------------------

optInFastK_Period:(From 1 to 100000)

Time period for building the Fast-K line

optInFastD_Period:(From 1 to 100000)

Smoothing for making the Fast-D line. Usually set to 3

optInFastD_MAType:

Type of Moving Average for Fast-D

*/
func Stochf(high, low, close []float64, fastKPeriod, fastDPeriod, fastDMAType int) ([]float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	outFastK := make([]float64, len(high))
	outFastD := make([]float64, len(high))
	C.TA_STOCHF(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(fastKPeriod), C.int(fastDPeriod), C.TA_MAType(fastDMAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outFastK[0])), (*C.double)(unsafe.Pointer(&outFastD[0])))
	return outFastK, outFastD
}

/*Stochrsi - Stochastic Relative Strength Index

Input = double

Output = double, double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

optInFastK_Period:(From 1 to 100000)

Time period for building the Fast-K line

optInFastD_Period:(From 1 to 100000)

Smoothing for making the Fast-D line. Usually set to 3

optInFastD_MAType:

Type of Moving Average for Fast-D

*/
func Stochrsi(real []float64, timePeriod, fastKPeriod, fastDPeriod, fastDMAType int) ([]float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	outFastK := make([]float64, len(real))
	outFastD := make([]float64, len(real))
	C.TA_STOCHRSI(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), C.int(fastKPeriod), C.int(fastDPeriod), C.TA_MAType(fastDMAType), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outFastK[0])), (*C.double)(unsafe.Pointer(&outFastD[0])))
	return outFastK, outFastD
}

/*Sub - Vector Arithmetic Substraction

Input = double, double

Output = double

*/
func Sub(real0, real1 []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real0))
	C.TA_SUB(0, C.int(len(real0)-1), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Sum - Summation

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Sum(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_SUM(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*T3 - Triple Exponential Moving Average (T3)

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

optInVFactor:(From 0 to 1)

Volume Factor

*/
func T3(real []float64, timePeriod int, vFactor float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_T3(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), C.double(vFactor), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Tan - Vector Trigonometric Tan

Input = double

Output = double

*/
func Tan(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_TAN(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Tanh - Vector Trigonometric Tanh

Input = double

Output = double

*/
func Tanh(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_TANH(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Tema - Triple Exponential Moving Average

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Tema(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_TEMA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Trange - True Range

Input = High, Low, Close

Output = double

*/
func Trange(high, low, close []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_TRANGE(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Trima - Triangular Moving Average

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Trima(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_TRIMA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Trix - 1-day Rate-Of-Change (ROC) of a Triple Smooth EMA

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

*/
func Trix(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_TRIX(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Tsf - Time Series Forecast

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Tsf(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_TSF(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Typprice - Typical Price

Input = High, Low, Close

Output = double

*/
func Typprice(high, low, close []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_TYPPRICE(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Ultosc - Ultimate Oscillator

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod1:(From 1 to 100000)

Number of bars for 1st period.

optInTimePeriod2:(From 1 to 100000)

Number of bars fro 2nd period

optInTimePeriod3:(From 1 to 100000)

Number of bars for 3rd period

*/
func Ultosc(high, low, close []float64, timePeriod1, timePeriod2, timePeriod3 int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_ULTOSC(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod1), C.int(timePeriod2), C.int(timePeriod3), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Var - Variance

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 1 to 100000)

Number of period

optInNbDev:(From TA_REAL_MIN to TA_REAL_MAX)

Nb of deviations

*/
func Var(real []float64, timePeriod int, nbDev float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_VAR(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), C.double(nbDev), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Wclprice - Weighted Close Price

Input = High, Low, Close

Output = double

*/
func Wclprice(high, low, close []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_WCLPRICE(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Willr - Williams' %R

Input = High, Low, Close

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Willr(high, low, close []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(high))
	C.TA_WILLR(0, C.int(len(high)-1), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}

/*Wma - Weighted Moving Average

Input = double

Output = double

Optional Parameters

-------------------

optInTimePeriod:(From 2 to 100000)

Number of period

*/
func Wma(real []float64, timePeriod int) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	outReal := make([]float64, len(real))
	C.TA_WMA(0, C.int(len(real)-1), (*C.double)(unsafe.Pointer(&real[0])), C.int(timePeriod), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&outReal[0])))
	return outReal
}
