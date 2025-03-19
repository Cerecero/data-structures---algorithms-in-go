package main

import "fmt"

func insertInOrder(slice []int, val int) []int {
	i := 0
	//Find the right index to where to insert the value
	// If i is less than the length of the slice
	// And the value in slice[i] is less than the given value
	// Then increment i to find the right index
	for i < len(slice) && slice[i] < val {
		i++
	}

	// Grow the slice, (append inserts at the end of the slice)
	slice = append(slice, 0)

	// copy(dst, src)
	// This will copy from a src (slice[i:]) into a dst (slice[i+1:])
	// dst slice starts from the index + 1 to the end of the slice or len(slice)
	// src starts from index to the end of the slice
	copy(slice[i+1:], slice[i:])

	// insert into i the value
	slice[i] = val
	return slice
}

func linearSearch(slice []int, val int) (int, bool) {

	//iterate throught the slice
	for i, v := range slice {
		// if the value is found return the index and a boolean value
		if val == v {
			return i, true
			// If v is greater than the value that we are searching for exit the loop
		} else if v > val {
			break
		}
	}
	// return 0 and false if the value is not whiting the array
	return 0, false
}

func binarySearch(slice []int, searchVal int) (int, bool) {
	lower_bound := 0
	upper_bound := len(slice) - 1 //last index of the slice

	for lower_bound <= upper_bound { // While loop
		// midpoint := lower_bound + (upper_bound-lower_bound)/2
		midpoint := (lower_bound + upper_bound)/2 //index

		valueMid := slice[midpoint]

		if searchVal == valueMid {
			return midpoint, true
		} else if searchVal < valueMid {

			// since we already checked that the midpoint is not the value we want
			// we move the upper bound to the midpoint minus one to not include the index we already checked
			// and search the left had of the array
			upper_bound = midpoint - 1
		} else if searchVal > valueMid {
			//search the right half
			lower_bound = midpoint + 1
		}
	}

	return -1, false
}

func main() {
	slice := []int{10, 32, 44, 54, 68}

	fmt.Println("Original slice: ", slice)

	slice = insertInOrder(slice, 25)
	fmt.Println("After inserting 25: ", slice)
	slice = insertInOrder(slice, -24)
	fmt.Println("After inserting -24: ", slice)

	search, found := linearSearch(slice, 32)
	fmt.Printf("Searching for 32: Found at index %d, %t\n", search, found)

	search, found = linearSearch(slice, 19)
	fmt.Printf("Searching for 19: Not found %t\n", found)

	search, found = binarySearch(slice, 54)
	fmt.Printf("Binary search for 54: Found: %t at index: %d\n", found, search)
	search, found = binarySearch(slice, 8)
	fmt.Printf("Binary search for 8: Found: %t\n", found)
	search, found = binarySearch(slice, 10)
	fmt.Printf("Binary search for 10: Found: %t at index: %d\n", found, search)
}
