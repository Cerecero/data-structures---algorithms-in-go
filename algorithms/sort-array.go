package main
import "fmt"

func insertInOrder(slice []int, val int) []int{
	i := 0
	for i < len(slice) && slice[i] < val {
		i++
	}

	slice = append(slice,0)

	copy(slice[i+1:], slice[i:])

	slice[i] = val
	return slice
}

func main(){
	slice := []int{10,32,54, 68}

	fmt.Println("Original slice: ", slice)

	slice = insertInOrder(slice, 25)
	fmt.Println("After inserting 25: ", slice)
	slice = insertInOrder(slice, -23)
	fmt.Println("After inserting -23: ", slice)
}
