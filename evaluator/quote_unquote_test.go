package evaluator

import (
	"testing"

	"github.com/nownabe/monkey-language-go/object"
)

func TestQuote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`quote(5)`, "5"},
		{`quote(5 + 8)`, `(5 + 8)`},
		{`quote(foobar)`, `foobar`},
		{`quote(foobar + barfoo)`, `(foobar + barfoo)`},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			evaluated := testEval(tt.input)
			quote, ok := evaluated.(*object.Quote)
			if !ok {
				t.Fatalf("expected *object.Quote. got=%T (%+v)", evaluated, evaluated)
			}

			if quote.Node == nil {
				t.Fatalf("quote.Node is nil")
			}

			if quote.Node.String() != tt.expected {
				t.Errorf("not equal. got=%q, wnat=%q", quote.Node.String(), tt.expected)
			}
		})
	}
}

func TestQuoteUnquote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`quote(unquote(4))`, "4"},
		{`quote(unquote(4 + 4))`, "8"},
		{`quote(8 + unquote(4 + 4))`, "(8 + 8)"},
		{`quote(unquote(4 + 4) + 8)`, "(8 + 8)"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			evaluated := testEval(tt.input)
			quote, ok := evaluated.(*object.Quote)
			if !ok {
				t.Fatalf("expected *object.Quote. got=%T (%+v)", evaluated, evaluated)
			}

			if quote.Node == nil {
				t.Fatalf("quote.Node is nil")
			}

			if quote.Node.String() != tt.expected {
				t.Errorf("not equal. got=%q, wnat=%q", quote.Node.String(), tt.expected)
			}
		})
	}
}
