package main

import (
	"github.com/emirpasic/gods/trees/binaryheap"
	"sort"

	//"github.com/emirpasic/gods/utils"
)
func minMeetingRooms(intervals [][]int) int {
	//sorting
	sort.Slice(intervals,func(i, j int) bool{
		 if intervals[i][0]-intervals[j][0]<0{
			 return true
		 }
		 return false
	})

	heap:=binaryheap.NewWith(func(a, b interface{}) int{
		ia:=a.(int)
		ib:=b.(int)
		return ia-ib
	})

	if len(intervals)==0{
		return 0
	}

	endtime:=intervals[0][1]
	heap.Push(endtime)
	for i:=1;i<len(intervals);i++{
		start:=intervals[i][0]
		endtime=intervals[i][1]

		//若當前會議start time在heap 頂部的時間之後，代表可以在該會議室接續開會，所以pop掉
		ipeak,ok:=heap.Peek()
		if ok{
			peakEndTime:=ipeak.(int)
			if start>=peakEndTime{
				heap.Pop()
			}
		}
		//push endtime
		heap.Push(endtime)

	}
	return heap.Size()
}