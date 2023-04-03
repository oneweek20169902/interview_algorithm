package base_sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// 测试用例1
	arr1 := []int{3, 2, 1, 5, 4}
	BubbleSort(arr1)
	if !reflect.DeepEqual(arr1, []int{1, 2, 3, 4, 5}) {
		t.Errorf("BubbleSort(%v) = %v; want [1 2 3 4 5]", []int{3, 2, 1, 5, 4}, arr1)
	}

	// 测试用例2
	arr2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	BubbleSort(arr2)
	if !reflect.DeepEqual(arr2, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) {
		t.Errorf("BubbleSort(%v) = %v; want [1 2 3 4 5 6 7 8 9 10]", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, arr2)
	}

	// 测试用例3
	arr3 := []int{1, 2, 3, 4, 5}
	BubbleSort(arr3)
	if !reflect.DeepEqual(arr3, []int{1, 2, 3, 4, 5}) {
		t.Errorf("BubbleSort(%v) = %v; want [1 2 3 4 5]", []int{1, 2, 3, 4, 5}, arr3)
	}
}

func TestQuickSort(t *testing.T) {
	arr := []int{3, 2, 1, 5, 4}
	QuickSort(arr, 0, len(arr)-1)
	expected := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		if arr[i] != expected[i] {
			t.Errorf("Sorting failed, expected %v but got %v", expected, arr)
			return
		}
	}
	fmt.Println("Sorting succeeded, result:", arr)
}
