package talib_test

import (
	"github.com/d4l3k/talib"

	"math"
	"reflect"
	"testing"
)

func TestInit(t *testing.T) {
}
func TestAcos(t *testing.T) {
	expected := []float64{1.5707963267948966, 0}
	out := talib.Acos([]float64{0, 1})
	if !reflect.DeepEqual(expected, out) {
		t.Errorf("Expected %#v got %#v.", expected, out)
	}
}
func TestSin(t *testing.T) {
	out := talib.Sin([]float64{0, math.Pi / 2})
	expected := []float64{0, 1}
	if !reflect.DeepEqual(expected, out) {
		t.Errorf("Expected %#v got %#v.", expected, out)
	}
}
func TestMacd(t *testing.T) {
	data := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}
	fast, slow, signal := talib.Macd(data, 12, 26, 9)
	expectedFast := []float64{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	expectedSlow := []float64{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	expectedSignal := []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if !reflect.DeepEqual(expectedFast, fast) {
		t.Errorf("Expected %#v got %#v.", expectedFast, fast)
	}
	if !reflect.DeepEqual(expectedSlow, slow) {
		t.Errorf("Expected %#v got %#v.", expectedSlow, slow)
	}
	if !reflect.DeepEqual(expectedSignal, signal) {
		t.Errorf("Expected %#v got %#v.", expectedSignal, signal)
	}
}
