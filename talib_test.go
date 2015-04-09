package talib_test

import (
	"github.com/d4l3k/talib"

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
