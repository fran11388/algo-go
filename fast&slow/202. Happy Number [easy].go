package main

func isHappy(n int) bool {
	slow:=getNext(n)
	fast:=getNext(slow)

	//持續往下
	// 最終結果要馬快的先出現1  要馬兩者相等

	for fast!=1 &&slow!=fast{

		slow=getNext(slow)
		fast=getNext(getNext(fast))


	}
	return fast==1
}

func getNext(n int)int{
	sum:=0
	// loop get digit
	for n>0{
		d:=n%10
		sum+=(d*d)
		n=n/10
	}
	return sum
}

