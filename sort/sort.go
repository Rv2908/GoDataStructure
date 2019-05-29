package sort

// BubbleSort Given an unordered list, we compare adjacent elements in the list,
// each time, putting in the right order of magnitude, only two elements.
// The algorithm hinges on a swap procedure.
// Knowing how many times to swap is important when implementing a bubble sort algorithm.
// To sort a list of numbers such as [3, 2, 1],
// we need to swap the elements a maximum of twice. This is equal to the length of the list minus 1
func BubbleSort(list []int) []int {

	var (
		n      = len(list)
		sorted = false
	)

	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if list[i] > list[i+1] {
				list[i+1], list[i] = list[i], list[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}

		n = n - 1
	}
	return list
}

// QuickSort The quick sort algorithm falls under the divide and conquer class of algorithms, 
// where we break (divide) a problem into smaller chunks that are much simpler to solve (conquer). 
// In this case, an unsorted array is broken into sub-arrays that are partially sorted, 
// until all elements in the list are in the right position, 
// by which time our unsorted list will have become sorted.
func QuickSort(list []int) []int{

	return list
}
