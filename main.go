package main

import (
	"fmt"
	"github.com/GoDataStructure/list"
	"github.com/GoDataStructure/search"
	"github.com/GoDataStructure/sort"
)

func main() {

	L := list.List{}
	L.Push(1)
	L.Push(2)
	L.Push(3)

	L.IterateForward()
	L.IterateBackward()

	fmt.Println("DataStructure And Algorithm Examples")
	
	nums := getNewArray()
	//Linear Search
	printResponse("Linear Search", search.LinearSearch(5, nums))
	//Binary Search
	printResponse("Binary Search", search.BinarySearch(3, nums))
	//Before Sort
	fmt.Println("Before Bubble Sort", nums)
	printResponse("After Bubble Sort", sort.BubbleSort(nums))
	nums = getNewArray()
	fmt.Println("Before Quick Sort", nums)
	printResponse("After Quick Sort", sort.QuickSort(nums))
}

func getNewArray() []int {
	return []int{1, 25, 14, 3, 5, -1}
}

func printResponse(message string, index interface{}) {
	// fmt.Println(message, index)
}
