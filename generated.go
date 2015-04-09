package talib

// #cgo LDFLAGS: -lta_lib
// #include "ta-lib/ta_libc.h"
import "C"

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
	real := make([]float64, len(real))
	C.ACOS(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Ad - Chaikin A/D Line
Input = High, Low, Close, Volume
Output = double
*/
func Ad(high, low, close, volume []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(high))
	C.AD(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), (*C.double)(unsafe.Pointer(&volume[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Add - Vector Arithmetic Add
Input = double, double
Output = double
*/
func Add(real0, real1 []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real0))
	C.ADD(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(high))
	C.ADOSC(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), (*C.double)(unsafe.Pointer(&volume[0])), fastPeriod, slowPeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(high))
	C.ADX(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(high))
	C.ADXR(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
func Apo(real []float64, fastPeriod, slowPeriod int, mAType TA_MAType) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.APO(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), fastPeriod, slowPeriod, mAType, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	aroonDown := make([]float64, len(high))
	aroonUp := make([]float64, len(high))
	C.AROON(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&aroonDown[0])), (*C.double)(unsafe.Pointer(&aroonUp[0])))
	return aroonDown, aroonUp
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
	real := make([]float64, len(high))
	C.AROONOSC(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Asin - Vector Trigonometric ASin
Input = double
Output = double
*/
func Asin(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.ASIN(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Atan - Vector Trigonometric ATan
Input = double
Output = double
*/
func Atan(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.ATAN(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(high))
	C.ATR(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Avgprice - Average Price
Input = Open, High, Low, Close
Output = double
*/
func Avgprice(open, high, low, close []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(open))
	C.AVGPRICE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
func Bbands(real []float64, timePeriod int, nbDevUp, nbDevDn float64, mAType TA_MAType) ([]float64, []float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	realUpperBand := make([]float64, len(real))
	realMiddleBand := make([]float64, len(real))
	realLowerBand := make([]float64, len(real))
	C.BBANDS(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, nbDevUp, nbDevDn, mAType, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&realUpperBand[0])), (*C.double)(unsafe.Pointer(&realMiddleBand[0])), (*C.double)(unsafe.Pointer(&realLowerBand[0])))
	return realUpperBand, realMiddleBand, realLowerBand
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
	real := make([]float64, len(real0))
	C.BETA(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Bop - Balance Of Power
Input = Open, High, Low, Close
Output = double
*/
func Bop(open, high, low, close []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(open))
	C.BOP(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(high))
	C.CCI(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Cdl2crows - Two Crows
Input = Open, High, Low, Close
Output = int
*/
func Cdl2crows(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDL2CROWS(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdl3blackcrows - Three Black Crows
Input = Open, High, Low, Close
Output = int
*/
func Cdl3blackcrows(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDL3BLACKCROWS(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdl3inside - Three Inside Up/Down
Input = Open, High, Low, Close
Output = int
*/
func Cdl3inside(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDL3INSIDE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdl3linestrike - Three-Line Strike
Input = Open, High, Low, Close
Output = int
*/
func Cdl3linestrike(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDL3LINESTRIKE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdl3outside - Three Outside Up/Down
Input = Open, High, Low, Close
Output = int
*/
func Cdl3outside(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDL3OUTSIDE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdl3starsinsouth - Three Stars In The South
Input = Open, High, Low, Close
Output = int
*/
func Cdl3starsinsouth(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDL3STARSINSOUTH(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdl3whitesoldiers - Three Advancing White Soldiers
Input = Open, High, Low, Close
Output = int
*/
func Cdl3whitesoldiers(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDL3WHITESOLDIERS(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
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
	integer := make([]int, len(open))
	C.CDLABANDONEDBABY(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), penetration, &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdladvanceblock - Advance Block
Input = Open, High, Low, Close
Output = int
*/
func Cdladvanceblock(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLADVANCEBLOCK(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlbelthold - Belt-hold
Input = Open, High, Low, Close
Output = int
*/
func Cdlbelthold(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLBELTHOLD(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlbreakaway - Breakaway
Input = Open, High, Low, Close
Output = int
*/
func Cdlbreakaway(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLBREAKAWAY(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlclosingmarubozu - Closing Marubozu
Input = Open, High, Low, Close
Output = int
*/
func Cdlclosingmarubozu(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLCLOSINGMARUBOZU(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlconcealbabyswall - Concealing Baby Swallow
Input = Open, High, Low, Close
Output = int
*/
func Cdlconcealbabyswall(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLCONCEALBABYSWALL(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlcounterattack - Counterattack
Input = Open, High, Low, Close
Output = int
*/
func Cdlcounterattack(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLCOUNTERATTACK(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
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
	integer := make([]int, len(open))
	C.CDLDARKCLOUDCOVER(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), penetration, &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdldoji - Doji
Input = Open, High, Low, Close
Output = int
*/
func Cdldoji(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLDOJI(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdldojistar - Doji Star
Input = Open, High, Low, Close
Output = int
*/
func Cdldojistar(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLDOJISTAR(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdldragonflydoji - Dragonfly Doji
Input = Open, High, Low, Close
Output = int
*/
func Cdldragonflydoji(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLDRAGONFLYDOJI(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlengulfing - Engulfing Pattern
Input = Open, High, Low, Close
Output = int
*/
func Cdlengulfing(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLENGULFING(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
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
	integer := make([]int, len(open))
	C.CDLEVENINGDOJISTAR(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), penetration, &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
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
	integer := make([]int, len(open))
	C.CDLEVENINGSTAR(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), penetration, &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlgapsidesidewhite - Up/Down-gap side-by-side white lines
Input = Open, High, Low, Close
Output = int
*/
func Cdlgapsidesidewhite(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLGAPSIDESIDEWHITE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlgravestonedoji - Gravestone Doji
Input = Open, High, Low, Close
Output = int
*/
func Cdlgravestonedoji(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLGRAVESTONEDOJI(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlhammer - Hammer
Input = Open, High, Low, Close
Output = int
*/
func Cdlhammer(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLHAMMER(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlhangingman - Hanging Man
Input = Open, High, Low, Close
Output = int
*/
func Cdlhangingman(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLHANGINGMAN(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlharami - Harami Pattern
Input = Open, High, Low, Close
Output = int
*/
func Cdlharami(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLHARAMI(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlharamicross - Harami Cross Pattern
Input = Open, High, Low, Close
Output = int
*/
func Cdlharamicross(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLHARAMICROSS(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlhighwave - High-Wave Candle
Input = Open, High, Low, Close
Output = int
*/
func Cdlhighwave(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLHIGHWAVE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlhikkake - Hikkake Pattern
Input = Open, High, Low, Close
Output = int
*/
func Cdlhikkake(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLHIKKAKE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlhikkakemod - Modified Hikkake Pattern
Input = Open, High, Low, Close
Output = int
*/
func Cdlhikkakemod(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLHIKKAKEMOD(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlhomingpigeon - Homing Pigeon
Input = Open, High, Low, Close
Output = int
*/
func Cdlhomingpigeon(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLHOMINGPIGEON(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlidentical3crows - Identical Three Crows
Input = Open, High, Low, Close
Output = int
*/
func Cdlidentical3crows(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLIDENTICAL3CROWS(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlinneck - In-Neck Pattern
Input = Open, High, Low, Close
Output = int
*/
func Cdlinneck(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLINNECK(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlinvertedhammer - Inverted Hammer
Input = Open, High, Low, Close
Output = int
*/
func Cdlinvertedhammer(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLINVERTEDHAMMER(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlkicking - Kicking
Input = Open, High, Low, Close
Output = int
*/
func Cdlkicking(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLKICKING(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlkickingbylength - Kicking - bull/bear determined by the longer marubozu
Input = Open, High, Low, Close
Output = int
*/
func Cdlkickingbylength(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLKICKINGBYLENGTH(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlladderbottom - Ladder Bottom
Input = Open, High, Low, Close
Output = int
*/
func Cdlladderbottom(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLLADDERBOTTOM(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdllongleggeddoji - Long Legged Doji
Input = Open, High, Low, Close
Output = int
*/
func Cdllongleggeddoji(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLLONGLEGGEDDOJI(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdllongline - Long Line Candle
Input = Open, High, Low, Close
Output = int
*/
func Cdllongline(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLLONGLINE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlmarubozu - Marubozu
Input = Open, High, Low, Close
Output = int
*/
func Cdlmarubozu(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLMARUBOZU(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlmatchinglow - Matching Low
Input = Open, High, Low, Close
Output = int
*/
func Cdlmatchinglow(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLMATCHINGLOW(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
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
	integer := make([]int, len(open))
	C.CDLMATHOLD(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), penetration, &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
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
	integer := make([]int, len(open))
	C.CDLMORNINGDOJISTAR(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), penetration, &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
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
	integer := make([]int, len(open))
	C.CDLMORNINGSTAR(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), penetration, &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlonneck - On-Neck Pattern
Input = Open, High, Low, Close
Output = int
*/
func Cdlonneck(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLONNECK(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlpiercing - Piercing Pattern
Input = Open, High, Low, Close
Output = int
*/
func Cdlpiercing(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLPIERCING(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlrickshawman - Rickshaw Man
Input = Open, High, Low, Close
Output = int
*/
func Cdlrickshawman(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLRICKSHAWMAN(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlrisefall3methods - Rising/Falling Three Methods
Input = Open, High, Low, Close
Output = int
*/
func Cdlrisefall3methods(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLRISEFALL3METHODS(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlseparatinglines - Separating Lines
Input = Open, High, Low, Close
Output = int
*/
func Cdlseparatinglines(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLSEPARATINGLINES(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlshootingstar - Shooting Star
Input = Open, High, Low, Close
Output = int
*/
func Cdlshootingstar(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLSHOOTINGSTAR(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlshortline - Short Line Candle
Input = Open, High, Low, Close
Output = int
*/
func Cdlshortline(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLSHORTLINE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlspinningtop - Spinning Top
Input = Open, High, Low, Close
Output = int
*/
func Cdlspinningtop(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLSPINNINGTOP(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlstalledpattern - Stalled Pattern
Input = Open, High, Low, Close
Output = int
*/
func Cdlstalledpattern(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLSTALLEDPATTERN(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlsticksandwich - Stick Sandwich
Input = Open, High, Low, Close
Output = int
*/
func Cdlsticksandwich(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLSTICKSANDWICH(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdltakuri - Takuri (Dragonfly Doji with very long lower shadow)
Input = Open, High, Low, Close
Output = int
*/
func Cdltakuri(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLTAKURI(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdltasukigap - Tasuki Gap
Input = Open, High, Low, Close
Output = int
*/
func Cdltasukigap(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLTASUKIGAP(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlthrusting - Thrusting Pattern
Input = Open, High, Low, Close
Output = int
*/
func Cdlthrusting(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLTHRUSTING(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdltristar - Tristar Pattern
Input = Open, High, Low, Close
Output = int
*/
func Cdltristar(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLTRISTAR(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlunique3river - Unique 3 River
Input = Open, High, Low, Close
Output = int
*/
func Cdlunique3river(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLUNIQUE3RIVER(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlupsidegap2crows - Upside Gap Two Crows
Input = Open, High, Low, Close
Output = int
*/
func Cdlupsidegap2crows(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLUPSIDEGAP2CROWS(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Cdlxsidegap3methods - Upside/Downside Gap Three Methods
Input = Open, High, Low, Close
Output = int
*/
func Cdlxsidegap3methods(open, high, low, close []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(open))
	C.CDLXSIDEGAP3METHODS(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&open[0])), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Ceil - Vector Ceil
Input = double
Output = double
*/
func Ceil(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.CEIL(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.CMO(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real0))
	C.CORREL(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Cos - Vector Trigonometric Cos
Input = double
Output = double
*/
func Cos(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.COS(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Cosh - Vector Trigonometric Cosh
Input = double
Output = double
*/
func Cosh(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.COSH(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.DEMA(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Div - Vector Arithmetic Div
Input = double, double
Output = double
*/
func Div(real0, real1 []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real0))
	C.DIV(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(high))
	C.DX(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.EMA(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Exp - Vector Arithmetic Exp
Input = double
Output = double
*/
func Exp(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.EXP(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Floor - Vector Floor
Input = double
Output = double
*/
func Floor(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.FLOOR(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*HtDcperiod - Hilbert Transform - Dominant Cycle Period
Input = double
Output = double
*/
func HtDcperiod(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.HT_DCPERIOD(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*HtDcphase - Hilbert Transform - Dominant Cycle Phase
Input = double
Output = double
*/
func HtDcphase(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.HT_DCPHASE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*HtPhasor - Hilbert Transform - Phasor Components
Input = double
Output = double, double
*/
func HtPhasor(real []float64) ([]float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	inPhase := make([]float64, len(real))
	quadrature := make([]float64, len(real))
	C.HT_PHASOR(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&inPhase[0])), (*C.double)(unsafe.Pointer(&quadrature[0])))
	return inPhase, quadrature
}

/*HtSine - Hilbert Transform - SineWave
Input = double
Output = double, double
*/
func HtSine(real []float64) ([]float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	sine := make([]float64, len(real))
	leadSine := make([]float64, len(real))
	C.HT_SINE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&sine[0])), (*C.double)(unsafe.Pointer(&leadSine[0])))
	return sine, leadSine
}

/*HtTrendline - Hilbert Transform - Instantaneous Trendline
Input = double
Output = double
*/
func HtTrendline(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.HT_TRENDLINE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*HtTrendmode - Hilbert Transform - Trend vs Cycle Mode
Input = double
Output = int
*/
func HtTrendmode(real []float64) []int {
	var outBegIdx C.int
	var outNBElement C.int
	integer := make([]int, len(real))
	C.HT_TRENDMODE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
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
	real := make([]float64, len(real))
	C.KAMA(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.LINEARREG(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.LINEARREG_ANGLE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.LINEARREG_INTERCEPT(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.LINEARREG_SLOPE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Ln - Vector Log Natural
Input = double
Output = double
*/
func Ln(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.LN(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Log10 - Vector Log10
Input = double
Output = double
*/
func Log10(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.LOG10(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
func Ma(real []float64, timePeriod int, mAType TA_MAType) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.MA(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, mAType, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	mACD := make([]float64, len(real))
	mACDSignal := make([]float64, len(real))
	mACDHist := make([]float64, len(real))
	C.MACD(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), fastPeriod, slowPeriod, signalPeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&mACD[0])), (*C.double)(unsafe.Pointer(&mACDSignal[0])), (*C.double)(unsafe.Pointer(&mACDHist[0])))
	return mACD, mACDSignal, mACDHist
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
func Macdext(real []float64, fastPeriod int, fastMAType TA_MAType, slowPeriod int, slowMAType TA_MAType, signalPeriod int, signalMAType TA_MAType) ([]float64, []float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	mACD := make([]float64, len(real))
	mACDSignal := make([]float64, len(real))
	mACDHist := make([]float64, len(real))
	C.MACDEXT(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), fastPeriod, fastMAType, slowPeriod, slowMAType, signalPeriod, signalMAType, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&mACD[0])), (*C.double)(unsafe.Pointer(&mACDSignal[0])), (*C.double)(unsafe.Pointer(&mACDHist[0])))
	return mACD, mACDSignal, mACDHist
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
	mACD := make([]float64, len(real))
	mACDSignal := make([]float64, len(real))
	mACDHist := make([]float64, len(real))
	C.MACDFIX(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), signalPeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&mACD[0])), (*C.double)(unsafe.Pointer(&mACDSignal[0])), (*C.double)(unsafe.Pointer(&mACDHist[0])))
	return mACD, mACDSignal, mACDHist
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
	mAMA := make([]float64, len(real))
	fAMA := make([]float64, len(real))
	C.MAMA(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), fastLimit, slowLimit, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&mAMA[0])), (*C.double)(unsafe.Pointer(&fAMA[0])))
	return mAMA, fAMA
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
func Mavp(real, periods []float64, minPeriod, maxPeriod int, mAType TA_MAType) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.MAVP(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), (*C.double)(unsafe.Pointer(&periods[0])), minPeriod, maxPeriod, mAType, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.MAX(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	integer := make([]int, len(real))
	C.MAXINDEX(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
}

/*Medprice - Median Price
Input = High, Low
Output = double
*/
func Medprice(high, low []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(high))
	C.MEDPRICE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(high))
	C.MFI(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), (*C.double)(unsafe.Pointer(&volume[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.MIDPOINT(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(high))
	C.MIDPRICE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.MIN(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	integer := make([]int, len(real))
	C.MININDEX(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&integer[0])))
	return integer
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
	min := make([]float64, len(real))
	max := make([]float64, len(real))
	C.MINMAX(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&min[0])), (*C.double)(unsafe.Pointer(&max[0])))
	return min, max
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
	minIdx := make([]int, len(real))
	maxIdx := make([]int, len(real))
	C.MINMAXINDEX(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.int)(unsafe.Pointer(&minIdx[0])), (*C.int)(unsafe.Pointer(&maxIdx[0])))
	return minIdx, maxIdx
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
	real := make([]float64, len(high))
	C.MINUS_DI(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(high))
	C.MINUS_DM(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.MOM(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Mult - Vector Arithmetic Mult
Input = double, double
Output = double
*/
func Mult(real0, real1 []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real0))
	C.MULT(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(high))
	C.NATR(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Obv - On Balance Volume
Input = double, Volume
Output = double
*/
func Obv(real, volume []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.OBV(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), (*C.double)(unsafe.Pointer(&volume[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(high))
	C.PLUS_DI(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(high))
	C.PLUS_DM(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
func Ppo(real []float64, fastPeriod, slowPeriod int, mAType TA_MAType) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.PPO(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), fastPeriod, slowPeriod, mAType, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.ROC(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.ROCP(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.ROCR(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.ROCR100(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.RSI(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(high))
	C.SAR(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), acceleration, maximum, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(high))
	C.SAREXT(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), startValue, offsetOnReverse, accelerationInitLong, accelerationLong, accelerationMaxLong, accelerationInitShort, accelerationShort, accelerationMaxShort, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Sin - Vector Trigonometric Sin
Input = double
Output = double
*/
func Sin(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.SIN(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Sinh - Vector Trigonometric Sinh
Input = double
Output = double
*/
func Sinh(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.SINH(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.SMA(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Sqrt - Vector Square Root
Input = double
Output = double
*/
func Sqrt(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.SQRT(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.STDDEV(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, nbDev, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
func Stoch(high, low, close []float64, fastKPeriod, slowKPeriod int, slowKMAType TA_MAType, slowDPeriod int, slowDMAType TA_MAType) ([]float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	slowK := make([]float64, len(high))
	slowD := make([]float64, len(high))
	C.STOCH(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), fastKPeriod, slowKPeriod, slowKMAType, slowDPeriod, slowDMAType, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&slowK[0])), (*C.double)(unsafe.Pointer(&slowD[0])))
	return slowK, slowD
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
func Stochf(high, low, close []float64, fastKPeriod, fastDPeriod int, fastDMAType TA_MAType) ([]float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	fastK := make([]float64, len(high))
	fastD := make([]float64, len(high))
	C.STOCHF(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), fastKPeriod, fastDPeriod, fastDMAType, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&fastK[0])), (*C.double)(unsafe.Pointer(&fastD[0])))
	return fastK, fastD
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
func Stochrsi(real []float64, timePeriod, fastKPeriod, fastDPeriod int, fastDMAType TA_MAType) ([]float64, []float64) {
	var outBegIdx C.int
	var outNBElement C.int
	fastK := make([]float64, len(real))
	fastD := make([]float64, len(real))
	C.STOCHRSI(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, fastKPeriod, fastDPeriod, fastDMAType, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&fastK[0])), (*C.double)(unsafe.Pointer(&fastD[0])))
	return fastK, fastD
}

/*Sub - Vector Arithmetic Substraction
Input = double, double
Output = double
*/
func Sub(real0, real1 []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real0))
	C.SUB(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real0[0])), (*C.double)(unsafe.Pointer(&real1[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.SUM(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.T3(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, vFactor, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Tan - Vector Trigonometric Tan
Input = double
Output = double
*/
func Tan(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.TAN(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Tanh - Vector Trigonometric Tanh
Input = double
Output = double
*/
func Tanh(real []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(real))
	C.TANH(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.TEMA(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Trange - True Range
Input = High, Low, Close
Output = double
*/
func Trange(high, low, close []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(high))
	C.TRANGE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.TRIMA(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.TRIX(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.TSF(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Typprice - Typical Price
Input = High, Low, Close
Output = double
*/
func Typprice(high, low, close []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(high))
	C.TYPPRICE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(high))
	C.ULTOSC(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), timePeriod1, timePeriod2, timePeriod3, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.VAR(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, nbDev, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Wclprice - Weighted Close Price
Input = High, Low, Close
Output = double
*/
func Wclprice(high, low, close []float64) []float64 {
	var outBegIdx C.int
	var outNBElement C.int
	real := make([]float64, len(high))
	C.WCLPRICE(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(high))
	C.WILLR(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&high[0])), (*C.double)(unsafe.Pointer(&low[0])), (*C.double)(unsafe.Pointer(&close[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
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
	real := make([]float64, len(real))
	C.WMA(0, C.int(len(double-1)), (*C.double)(unsafe.Pointer(&real[0])), timePeriod, &outBegIdx, &outNBElement, (*C.double)(unsafe.Pointer(&real[0])))
	return real
}

/*Setunstableperiod - Some TA functions takes a certain amount of input data
before stabilizing and outputing meaningful data. This is
a behavior pertaining to the algo of some TA functions and
is not particular to the TA-Lib implementation.
TA-Lib allows you to automatically strip off these unstabl
data from your output and from any internal processing.
(See documentation for more info)
Examples:
TA_Setunstableperiod( TA_FUNC_UNST_EMA, 30 );
Always strip off 30 price bar for the TA_EMA function.
TA_Setunstableperiod( TA_FUNC_UNST_ALL, 30 );
Always strip off 30 price bar from ALL functions
having an unstable period.
See ta_defs.h for the enumeration TA_FuncUnstId
*/
func Setunstableperiod() {

	C.SetUnstablePeriod()
	return
}

/*Setunstableperiod - Some TA functions takes a certain amount of input data
before stabilizing and outputing meaningful data. This is
a behavior pertaining to the algo of some TA functions and
is not particular to the TA-Lib implementation.
TA-Lib allows you to automatically strip off these unstabl
data from your output and from any internal processing.
(See documentation for more info)
Examples:
TA_Setunstableperiod( TA_FUNC_UNST_EMA, 30 );
Always strip off 30 price bar for the TA_EMA function.
TA_Setunstableperiod( TA_FUNC_UNST_ALL, 30 );
Always strip off 30 price bar from ALL functions
having an unstable period.
See ta_defs.h for the enumeration TA_FuncUnstId
*/
func Setunstableperiod() {

	C.SetUnstablePeriod()
	return
}

/*Setcompatibility - You can change slightly the behavior of the TA functions
by requesting compatibiliy with some existing software.
By default, the behavior is as close as the original
author of the TA functions intend it to be.
See ta_defs.h for the enumeration TA_Compatibility.
*/
func Setcompatibility() {

	C.SetCompatibility()
	return
}

/*Setcompatibility - You can change slightly the behavior of the TA functions
by requesting compatibiliy with some existing software.
By default, the behavior is as close as the original
author of the TA functions intend it to be.
See ta_defs.h for the enumeration TA_Compatibility.
*/
func Setcompatibility() {

	C.SetCompatibility()
	return
}

/*Setcandlesettings - Call TA_Setcandlesettings to set that when comparing a candle
basing on settingType it must be compared with the average
of the last avgPeriod candles' rangeType multiplied by factor.
This setting is valid until TA_RestoreCandleDefaultSettings is called
*/
func Setcandlesettings() {

	C.SetCandleSettings()
	return
}

/*Restorecandledefaultsettings - Call TA_Restorecandledefaultsettings after using custom settings
to restore the default settings for the specified settingType
*/
func Restorecandledefaultsettings() {

	C.RestoreCandleDefaultSettings()
	return
}
