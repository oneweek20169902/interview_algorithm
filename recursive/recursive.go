package recursive

import (
	"github.com/spf13/cast"
	"math"
)

func GetMax(arr []int) int {
	return process(arr, 0, len(arr))
}

func process(arr []int, l, r int) int {
	if l == r {
		return arr[l]
	}
	mid := l + ((r - l) >> 1) // 中点
	leftMax := process(arr, l, mid)
	rightMax := process(arr, mid+1, r)
	return cast.ToInt(math.Max(float64(leftMax), float64(rightMax)))
}

// DutchFlag 荷兰国旗
func DutchFlag(arr []int, num int) {
	/*
		1. i<num, i和<区域下一个交换，<区域右扩，i++
		2. i==num, i++
		3.i>num,i和>区域前一个交换，>区域右扩,i原地不动
	*/

}

// QuickSort 快排
func QuickSort(arr []int, l, r int) {
	if l < r {
		p := partition(arr, l, r)
		QuickSort(arr, l, p-1) //< 区
		QuickSort(arr, p+1, r) //> 区
	}
}

/*
这是一个处理arr[...]函数
默认以arr[r]做划分,arr[r] -> p  <p ==p   >p
返回等于区域(左边界，有边界) 所以返回一个长度为2的数组res,res[0] res[1]
*/
func partition(arr []int, l, r int) int {
	pivot := arr[r]
	i := l - 1 //<区右边界
	for j := l; j < r; j++ {
		if arr[j] >= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}
