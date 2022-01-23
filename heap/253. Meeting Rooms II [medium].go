package main

import (
	"github.com/emirpasic/gods/trees/binaryheap"
	"sort"

	//"github.com/emirpasic/gods/utils"
)

func minMeetingRooms(intervals [][]int) int {
	return minMeetingRoomsByTimeOrderingTwoPointer(intervals)
}

func minMeetingRoomsByTimeOrderingTwoPointer(intervals [][]int) int {
	beginTimes:=make([]int,len(intervals))
	endTimes:=make([]int,len(intervals))
	//init
	for i,interval:=range intervals{
		beginTimes[i]=interval[0]
		endTimes[i]=interval[1]
	}
	sort.Slice(beginTimes, func(i, j int) bool {
		return beginTimes[i]-beginTimes[j]<0
	})
	sort.Slice(endTimes, func(i, j int) bool {
		return endTimes[i]-endTimes[j]<0
	})
	rooms:=0
	//iterate begintimes
	endtimePtr:=0
	//關注會議起始時間是否在最近的結束時間之後，如果是代表可延續該會議室使用(不須assign room)，
	//然後endtimePtr+1 往後考慮最近的會議結束時間;
	//若不是就需要assign room
	//這邊做法跟leetcode解不同，有需要才assign room, 而不是扣掉又再次assign
	for _,begintime:=range beginTimes{
		recentEndtime:=endTimes[endtimePtr]
		if begintime>=recentEndtime{
			endtimePtr++
		}else{
			rooms++
		}
	}
	return rooms
}

func minMeetingRoomsByHeap(intervals [][]int) int {
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

		//若當前會議start time在heap 頂部的時間之後(代表最近的會議結束時間)，代表可以在該會議室接續開會，
		//所以pop掉最近的會議結束時間，
		//然後push新的會議的結束時間
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