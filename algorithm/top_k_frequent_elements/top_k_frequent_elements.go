package main

import (
	"fmt"
	"reflect"
)

type HeapNode struct {
	Val, Count int
}

func topKFrequent(nums []int, k int) []int {
	record := make(map[int]int)
	heap := make([]HeapNode, 0)
	for _, c := range nums {
		if i, ok := record[c]; ok {
			heap[i].Count++
			adjustRecordHeap(record, heap)
		} else {
			heap = append(heap, HeapNode{Val: c, Count: 1})
			record[c] = len(heap) - 1
			adjustRecordHeap(record, heap)
		}
	}
	fmt.Println("Heap", heap)
	fmt.Println("Record", record)
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, heap[0].Val)
		delete(record, heap[0].Val)
		heap = heap[1:]
		adjustRecordHeap(record, heap)
	}
	return res
}

func adjustRecordHeap(record map[int]int, heap []HeapNode) {
	root := 0
	for {
		child := root*2 + 1
		if child >= len(heap) {
			return
		}
		if child+1 < len(heap) && heap[child].Count < heap[child+1].Count {
			child++
		}
		if heap[root].Count >= heap[child].Count {
			root++
			continue
		}
		record[heap[root].Val], record[heap[child].Val] = record[heap[child].Val], record[heap[root].Val]
		heap[root].Val, heap[root].Count, heap[child].Val, heap[child].Count = heap[child].Val, heap[child].Count, heap[root].Val, heap[root].Count
		root = child
	}
}

func main() {
	testCases := []struct {
		nums   []int
		k      int
		expect []int
	}{
		{nums: []int{1, 1, 1, 2, 2, 3}, k: 2, expect: []int{1, 2}},
		{nums: []int{2, 3, 4, 1, 4, 0, 4, -1, -2, -1}, k: 2, expect: []int{4, -1}},
	}
	for i, c := range testCases {
		res := topKFrequent(c.nums, c.k)
		eq := reflect.DeepEqual(res, c.expect)
		if !eq {
			fmt.Println(i, res, c.expect, c.k, c.nums)
		}
	}

}
