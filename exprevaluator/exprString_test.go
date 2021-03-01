package exprevaluator

import (
	"fmt"
	"testing"
)

// TestString test the concrete type for implementing
// String() method
func TestString(t *testing.T) {
	tests := []struct {
		expr    string
		wantout string
	}{
		{"2*b-c", "2 * b - c"},
		{"sqrt(A/Pi)", "sqrt(A / Pi)"},
		{"pow(3,4)+6*7", "pow(3, 4) + 6 * 7"},
	}

	var preExpr string
	for _, test := range tests {
		if test.expr != preExpr {
			fmt.Printf("\n%s\n", test.expr)
			preExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%s", expr)
		fmt.Printf("\t%s pretty print is %s", test.expr, got)
		if got != test.wantout {
			t.Errorf("%s pretty print want is %s, want %s\n", test.expr, got, test.wantout)
		}
	}
}
