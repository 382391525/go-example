package ltest

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "2 + 3",
			args: args{
				a: 2,
				b: 3,
			},
			want: 5,
		},
		{
			name: "2 + 4",
			args: args{
				a: 2,
				b: 4,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "2 * 3",
			args: args{
				a: 2,
				b: 3,
			},
			want: 6,
		},
		{
			name: "2 * 4",
			args: args{
				a: 2,
				b: 4,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mul(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}
