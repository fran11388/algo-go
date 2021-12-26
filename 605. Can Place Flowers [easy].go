package main

import "fmt"

func main(){
	t:=[]int{0,0,1,0,0}
	n:=1
	fmt.Println(canPlaceFlowers(t,n))

}
func canPlaceFlowers(flowerbed []int, n int) bool {
	//loop arr
	//	if can place
	for i, _ := range flowerbed {
		if canPlace(flowerbed, i) {
			place(flowerbed, i)
			n--
		}
	}
	return n <= 0
}

func canPlace(flowerbed []int, idx int)bool{
	leftNoFlower:=true
	if idx-1>=0&&flowerbed[idx-1]==1{
		leftNoFlower=false
	}

	rightNoFlower:=true
	if idx+1<len(flowerbed)&&flowerbed[idx+1]==1{
		rightNoFlower=false
	}

	if flowerbed[idx]==0&&leftNoFlower&&rightNoFlower{
		return true
	}
	return false
}

func place(flowerbed []int, idx int){
	flowerbed[idx]=1
}