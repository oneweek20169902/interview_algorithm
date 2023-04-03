package base_sort

// BubbleSort 冒泡排序
func BubbleSort(nums []int) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// QuickSort 快速排序
func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := arr[left]
	i, j := left, right
	for i < j {
		// 从右往左找到第一个小于等于pivot的元素
		for i < j && arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		// 从左往右找到第一个大于pivot的元素
		for i < j && arr[j] <= pivot {
			i++
		}

		if i < j {
			arr[i] = arr[j]
			i++
		}

	}
	arr[i] = pivot
	QuickSort(arr, left, i-1)  // 递归排序左边的子数组
	QuickSort(arr, i+1, right) // 递归排序右边的子数组

}
