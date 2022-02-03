package main
func mySqrt(x int) int {
	// range = 2<= a <x/2
	if x<2{
		return x
	}

	left:=2
	right :=x/2
	for left<=right{

		pivot:=(left+right)/2
		n:=pivot*pivot

		if n==x{
			return pivot
		}else if n>x{
			right=pivot-1
		}else{
			left=pivot+1
		}
	}
	return right
}