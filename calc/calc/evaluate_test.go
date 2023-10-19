package calc

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	// arrange
	tbl := []struct {
		input  string
		result string
	}{
		{"1 + 2 * 3", "7"},
		{"( 1 + 2 ) * 3", "9"},
		{"1 - -7", "8"},
		{"( 1 + 2 * ( 1 + 1 )) * 9", "45"},
		{"( -1 + 2 * ( 1 + 1 )) * 9", "27"},
		{"1+2*3", "7"},
		{"1+3/2", "2.5"},
		{"1+ (2+1)/3", "2"},
	}

	for i, tt := range tbl {
		t.Run(fmt.Sprintf("%v %v", i, tt.input), func(t *testing.T) {
			// act
			resFloat, err := Eval(tt.input)

			// assert
			if err != nil {
				t.Fatal(err)
			}

			res := fmt.Sprintf("%v", resFloat)
			if res != tt.result {
				t.Errorf("invalid outcome '%v', expected '%v'", res, tt.result)
			}
		})
	}
}
