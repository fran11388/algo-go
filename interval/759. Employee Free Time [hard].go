package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/binaryheap"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func main() {
	arr := [][]*Interval{{&Interval{1, 3}, &Interval{5, 7}}}
	fmt.Println(employeeFreeTime(arr))
}
func employeeFreeTime(schedule [][]*Interval) []*Interval {
	var result []*Interval
	allintervals := []*Interval{}
	for _, intervals := range schedule {
		for _, interval := range intervals {
			allintervals = append(allintervals, interval)
		}
	}
	if len(allintervals) == 0 {
		return result
	}

	sort.Slice(allintervals, func(i, j int) bool {
		return allintervals[i].Start-allintervals[j].Start < 0
	})

	lastEndTime := allintervals[0].End
	for i := 1; i < len(allintervals); i++ {
		currIntervalStart := allintervals[i].Start
		currIntervalEnd := allintervals[i].End
		if currIntervalStart > lastEndTime { //判斷每個interval起始有沒有在最近的結束時間之後，有的話形成gap(free time)
			result = append(result, &Interval{Start: lastEndTime, End: currIntervalStart})
		}
		if currIntervalEnd > lastEndTime {
			lastEndTime = currIntervalEnd
		}
	}
	return result
}

func employeeFreeTimeByHeap(schedule [][]*Interval) []*Interval {
	heap := binaryheap.NewWith(func(a, b interface{}) int {
		ia := a.(*Interval)
		ib := b.(*Interval)
		return ia.Start - ib.Start
	})
	//push every interval to heap
	for _, intervals := range schedule {
		for _, interval := range intervals {
			heap.Push(interval)
		}
	}

	var result []*Interval
	//pop one
	tmp, _ := heap.Pop()
	tmpInterval := tmp.(*Interval)
	for !heap.Empty() {
		ipeak, _ := heap.Peek()
		peak := ipeak.(*Interval)
		if peak.Start > tmpInterval.End { //找到gap(free time)
			result = append(result, &Interval{tmpInterval.End, peak.Start})
		}
		if peak.End > tmpInterval.End { //要更新tmpInterval 作為下次比較的區間
			tmpInterval = peak
		}
		heap.Pop() //pop這次比較的interval
	}
	return result
}
