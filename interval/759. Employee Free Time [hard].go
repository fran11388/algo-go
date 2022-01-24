package main

import (
	"github.com/emirpasic/gods/trees/binaryheap"
)

type Interval struct {
	Start int
	End   int
}

func employeeFreeTimeByHeap(schedule [][]*Interval) []*Interval {
	heap:=binaryheap.NewWith(func(a, b interface{}) int{
		ia:=a.(*Interval)
		ib:=b.(*Interval)
		return ia.Start-ib.Start
	})
	//push every interval to heap
	for _,intervals:=range schedule{
		for _,interval:=range intervals{
			heap.Push(interval)
		}
	}

	var result []*Interval
	//pop one
	tmp,_:=heap.Pop()
	tmpInterval:=tmp.(*Interval)
	for !heap.Empty(){
		ipeak,_:=heap.Peek()
		peak:=ipeak.(*Interval)
		if peak.Start>tmpInterval.End{ //找到gap(free time)
			result=append(result,&Interval{tmpInterval.End,peak.Start})
		}
		if peak.End>tmpInterval.End{ //要更新tmpInterval 作為下次比較的區間
			tmpInterval=peak
		}
		heap.Pop() //pop這次比較的interval
	}
	return result
}
