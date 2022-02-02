package _3Sum_Closest

import (
	"testing"
)

func Test_threeSumClosestByTwoPointer(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "ex1",
			args: args{[]int{-1,2,1,-4},1},
			want: 2,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosestByTwoPointer(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosestByTwoPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}
