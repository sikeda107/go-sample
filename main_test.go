package main

import (
	"testing"
)

func Test_hello(t *testing.T) {
	actual := hello()
	expected := "Hello, world."
	if actual != expected {
		t.Errorf("got: %v\nwant: %v", actual, expected)
	}
}

func Test_greet(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"テスト0", args{0}, "Hello, world."},
		{"テスト1", args{1}, "Good Morning"},
		{"テスト2", args{2}, "Hello"},
		{"テスト3", args{3}, "Good Evening"},
		{"テスト4", args{4}, "Hello, world!!!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greet(tt.args.n); got != tt.want {
				t.Errorf("greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
