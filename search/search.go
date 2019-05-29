package search

// LinearSearch This technique pass over the list of elements,
// by using the index to move from the beginning of the list to the end.
// Each element is examined and if it does not match the search item,
// the next item is examined.
// By hopping from one item to its next, the list is passed over sequentially.
func LinearSearch(key int, list []int) int {
	for index, item := range list {
		if item == key {
			return index
		}
	}
	return -1
}

// BinarySearch  A binary search is a search strategy used to find elements within a list
// by consistently reducing the amount of data to be searched
// and thereby increasing the rate at which the search term is found.
// To use a binary search algorithm, the list to be operated on must have already been sorted.
func BinarySearch(key int, list []int) int {
	low := 0
	high := len(list) - 1

	for i := low; i <= high; i++ {
		median := (low + high) / 2
		if list[median] == key {
			return i
		} else if list[median] < key {
			low = median + 1
			continue
		} else if list[median] > key {
			high = median - 1
			continue
		}
	}

	return -1
}

