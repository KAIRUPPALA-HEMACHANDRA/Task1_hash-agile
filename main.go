package main
import "fmt"

func SecLargest(arr []int) int {
	Largest := arr[0]
	secondLargest := -1
	for _, num := range arr {
		if num > Largest {
			secondLargest = Largest
			Largest = num
		} else if num > secondLargest && num != Largest {
			secondLargest = num
		}
	}
	if secondLargest == -1 {
		fmt.Println("No second largest element")
		return -1
	}
	return secondLargest
}

func main() {
	var size int
	fmt.Print("Enter the size of the array: ")
	fmt.Scan(&size)
	arr := make([]int, size)
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}
	secondLargest := SecLargest(arr)
	fmt.Println("The second largest element is:", secondLargest)
}
