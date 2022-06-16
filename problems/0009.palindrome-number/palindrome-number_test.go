package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "provided_1",
			args: args{121},
			want: true,
		},
		{
			name: "provided_2",
			args: args{-121},
			want: false,
		},
		{
			name: "provided_3",
			args: args{10},
			want: false,
		},
		{
			name: "zero",
			args: args{0},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
