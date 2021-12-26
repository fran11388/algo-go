package main

import "fmt"

//import "fmt"

func main(){
	e:=[]int{0,1}
	a:=e[:len(e)-1]
	b:=e[1:]
	fmt.Println(a,b)
}

func show9(i int){

}
func getS()[]string{
	return []string{}
}

func addsome(arr *[]int){
*arr=append(*arr,1)
}
//func swap(s []int, idx1,idx2 int){
//	t:=s[idx1]
//	s[idx1]=s[idx2]
//	s[idx2]=t
//}