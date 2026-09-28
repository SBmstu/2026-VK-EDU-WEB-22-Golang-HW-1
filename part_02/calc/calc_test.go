package calc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEval(t *testing.T) {
	tests := []struct {
		name    string
		expr    string
		want    float64
		wantErr bool
	}{
		{name: "addition", expr: "1+2", want: 3},
		{name: "subtraction", expr: "5-3", want: 2},
		{name: "multiplication", expr: "4*5", want: 20},
		{name: "division", expr: "10/4", want: 2.5},
		{name: "parentheses", expr: "(1+2)*3", want: 9},
		{name: "unary minus", expr: "-5+2", want: -3},
		{name: "precedence", expr: "2+3*4", want: 14},
		{name: "multiline", expr: "(1\n+ 2) *\n3\n/ 4", want: 2.25},
		{name: "error", expr: "1/0", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Eval(tt.expr)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.InDelta(t, tt.want, got, 1e-9)
		})
	}
}
