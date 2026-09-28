package uniq

import (
	"reflect"
	"testing"
	"uniq/types"
)

func TestUniq(t *testing.T) {
	tests := []struct {
		name string
		lines []string
		options types.Options
		want []string
		wantErr bool
	}{
		{
			name: "empty input",
			lines: []string{},
			options: types.Options{},
			want: []string{},
		},
		{
			name: "single line",
			lines: []string{"a"},
			options: types.Options{},
			want: []string{"a"},
		},
		{
			name: "all unique",
			lines: []string{"a", "b", "c"},
			options: types.Options{},
			want: []string{"a", "b", "c"},
		},
		{
			name: "all identical",
			lines: []string{"x", "x", "x"},
			options: types.Options{},
			want: []string{"x"},
		},
		{
			name: "count",
			lines: []string{"a", "a", "b", "b", "b", "c", "a"},
			options: types.Options{Count: true},
			want: []string{"2 a", "3 b", "1 c", "1 a"},
		},
		{
			name: "repeated",
			lines: []string{"a", "a", "b", "b", "b", "c", "a"},
			options: types.Options{Repeated: true},
			want: []string{"a", "b"},
		},
		{
			name: "unique",
			lines: []string{"a", "a", "b", "b", "b", "c", "a"},
			options: types.Options{Unique: true},
			want: []string{"c", "a"},
		},
		{
			name: "ignore case",
			lines: []string{"A", "a", "b", "B", "c"},
			options: types.Options{IgnoreCase: true},
			want: []string{"A", "b", "c"},
		},
		{
			name: "ignore case with count",
			lines: []string{"A", "a", "b"},
			options: types.Options{Count: true, IgnoreCase: true},
			want: []string{"2 A", "1 b"},
		},
		{
			name: "skip fields",
			lines: []string{"1 a", "2 a", "1 b", "2 b"},
			options: types.Options{SkipFields: 1},
			want: []string{"1 a", "1 b"},
		},
		{
			name: "skip chars",
			lines: []string{"x a", "y a", "x b"},
			options: types.Options{SkipChars: 2},
			want: []string{"x a", "x b"},
		},
		{
			name: "lines with spaces",
			lines: []string{"   ", "   ", "x"},
			options: types.Options{},
			want: []string{"   ", "x"},
		},
		{
			name: "mutually exclusive c and d",
			lines: []string{"a"},
			options: types.Options{Count: true, Repeated: true},
			wantErr: true,
		},
		{
			name: "mutually exclusive c and u",
			lines: []string{"a"},
			options: types.Options{Count: true, Unique: true},
			wantErr: true,
		},
		{
			name: "mutually exclusive d and u",
			lines: []string{"a"},
			options: types.Options{Repeated: true, Unique: true},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uniq(tt.lines, tt.options)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Uniq() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
