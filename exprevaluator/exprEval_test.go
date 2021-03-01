package exprevaluator

import (
	"fmt"
	"math"
	"testing"
)

// TestEval test the concrete types to satisify Expr interface
func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / Pi)", Env{"A": 87616, "Pi": math.Pi}, "167"},
		{"pow(x, 3)+pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"5*a+6*b", Env{"a": 1, "b": 2}, "17"},
		{"5+a*b", Env{"a": 1, "b": 2}, "7"},
		{"5 / 9 * (f - 32)", Env{"f": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32.0}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212.0}, "100"},
		{"6 / 9 * (x-1)", Env{"x": 4}, "2"},
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
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n", test.expr, test.env, got, test.want)
		}
	}
}
