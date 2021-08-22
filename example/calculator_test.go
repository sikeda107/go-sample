package example

import (
	"testing"
)

func TestCalculator_Add1(t *testing.T) {
	type fields struct {
		A int
		B int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{"テスト0", fields{0, 0}, 0},
		{"テスト1", fields{0, 1}, 1},
		{"テスト2", fields{1, 0}, 1},
		{"テスト3", fields{1, 2}, 3},
		{"テスト4", fields{-1, -2}, -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				A: tt.fields.A,
				B: tt.fields.B,
			}
			if got := c.Add1(); got != tt.want {
				t.Errorf("Calculator.Add1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculator_Add2_1(t *testing.T) {
	type fields struct {
		A int
		B int
	}
	type args struct {
		option int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
		{"テスト0", fields{A: 0, B: 0}, args{option: 0}, 0},
		{"テスト1", fields{A: 0, B: 0}, args{option: 1}, 1},
		{"テスト2", fields{A: 0, B: 0}, args{option: -1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Calculator{
				A: tt.fields.A,
				B: tt.fields.B,
			}
			if got := c.Add2(tt.args.option); got != tt.want {
				t.Errorf("Calculator.Add2() = %v, want %v", got, tt.want)
			}
		})
	}
}
