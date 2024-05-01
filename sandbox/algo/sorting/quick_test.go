package sorting

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{10, 7, 8, 9, 1, 5}
	QuickSort(arr)
	fmt.Println("Sorted array: ", arr)
}
