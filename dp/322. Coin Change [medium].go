package main

import "fmt"

func main(){
	coins:=[]int{1,2,5}
	amount:=11
	fmt.Println(coinChange(coins,amount))
}

func coinChange(coins []int, amount int) int {
	dpTable:=make([]int,amount+1) //
	//init dptable

	for i,_:=range dpTable{
		dpTable[i]=amount+1
	}
	dpTable[0]=0

	for i:=0;i<amount;i++{
		for _,coin:=range coins{
			currAmountMinCoin:=dpTable[i]
			if currAmountMinCoin==amount+1{ //代表他還是初始直=換不到
				continue
			}

			if (i+coin)<=amount{
				newMinCoin:=currAmountMinCoin+1 //加1代表當前硬幣
				if newMinCoin<dpTable[i+coin]{
					dpTable[i+coin]=newMinCoin
				}
			}

		}
	}


	if dpTable[amount]==(amount+1){
		return -1
	}
	return dpTable[amount]
}