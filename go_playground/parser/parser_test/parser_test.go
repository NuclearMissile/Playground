package parser_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/NuclearMissile/Playground/go_playground/parser"
)

func BenchmarkParser(b *testing.B) {
	var tests = []struct {
		input    string
		env      parser.Env
		expected string
	}{
		{"x%2", nil, "unexpected '%'"},
		{"ln(e)", parser.Env{"e": math.E}, "1"},
		{"sqrt(A / pi)", parser.Env{"A": 87616, "pi": math.Pi}, "167"},
	}

	for _, test := range tests {
		expr, err := parser.Parse(test.input)
		if err == nil {
			expr.Check(map[parser.Var]bool{})
		}
		if err != nil {
			if err.Error() != test.expected {
				b.Errorf("%s: got %q, expected: %q", test.input, err, test.expected)
			}
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		if got != test.expected {
			b.Errorf("%s: %v => %s, expected: %s", test.input, test.env, got, test.expected)
		}
	}
}
func TestParser(t *testing.T) {
	var tests = []struct {
		input    string
		env      parser.Env
		expected string
	}{
		{"x%2", nil, "unexpected '%'"},
		{"ln(e)", parser.Env{"e": math.E}, "1"},
		{"sqrt(A / pi)", parser.Env{"A": 87616, "pi": math.Pi}, "167"},
	}

	for _, test := range tests {
		expr, err := parser.Parse(test.input)
		if err == nil {
			expr.Check(map[parser.Var]bool{})
		}
		if err != nil {
			if err.Error() != test.expected {
				t.Errorf("%s: got %q, expected: %q", test.input, err, test.expected)
			}
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		if got != test.expected {
			t.Errorf("%s: %v => %s, expected: %s", test.input, test.env, got, test.expected)
		}
	}
}
