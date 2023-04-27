package recursive

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestGetMax(t *testing.T) {
	testCases := []struct {
		name        string
		input       []int
		expectedMax int
	}{
		{
			name:        "Valid positive integers",
			input:       []int{1, 3, 5, 2, 4},
			expectedMax: 5,
		},
		{
			name:        "Valid negative integers",
			input:       []int{-1, -3, -5, -2, -4},
			expectedMax: -1,
		},
		{
			name:        "Valid mix of positive and negative integers",
			input:       []int{-1, 3, 5, -2, 4},
			expectedMax: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GetMax(tc.input)
			if result != tc.expectedMax {
				t.Errorf("Expected maximum value to be %d, but got %d", tc.expectedMax, result)
			}
		})
	}

	// Additional test cases for edge cases:
	// Empty slice case
	result := GetMax([]int{})
	if result != math.MinInt64 {
		t.Errorf("Expected maximum value to be %d, but got %d", math.MinInt64, result)
	}
	// Single element slice case
	result = GetMax([]int{42})
	if result != 42 {
		t.Errorf("Expected maximum value to be %d, but got %d", 42, result)
	}
}

func TestQuickSort(t *testing.T) {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	QuickSort(arr, 0, len(arr)-1)
	if !sort.IntsAreSorted(arr) {
		t.Errorf("failed to sort arr: %v", arr)
	}
}

func TestDutchFlag(t *testing.T) {
	arr := []int{2, 3, 1, 9, 7, 6, 1, 4, 5}
	v := DutchFlag(arr, 4)
	t.Log(v)

}

func TestSortColors(t *testing.T) {
	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors1(arr)

}

func TestHeapSort(t *testing.T) {

	str := "【慧泽园·乐凯店】尊敬的贵宾，您好！感谢您莅临慧泽园！\\\\r\\\\n餐位(包厢):三楼雅间-321\\\\r\\\\n日期:2023年04月11日 周二 晚市 \\\\r\\\\n人数:1桌10人\\\\r\\\\n联系人:0312-3120333(詹红艳18617777513)\\\\r\\\\n导航:https://dwz.cn/o5tBITwv\\\\r\\\\n邀请函:https://dwz.cn/EAfghYJ0\\\\n地址：乐凯北大街980号，慧泽园·乐凯店全体员工恭候您的光临！\\\\r\\\\n\\"

	fmt.Println(len([]rune(str)))

	nums := []int{3, 5, 3, 0, 8, 6}
	newNums := HeapSort(nums)
	t.Log(newNums)

}
