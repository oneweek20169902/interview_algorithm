package recursive

import (
	"github.com/spf13/cast"
	"math"
	"math/rand"
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
		swap(arr, rand.Intn(l+(r-l+1)), r) //有随机行为所以复杂度是O(n*logn)
		p := partition(arr, l, r)
		QuickSort(arr, l, p[0]-1) //< 区
		QuickSort(arr, p[0]+1, r) //> 区
	}
}

/*
这是一个处理arr[...]函数
默认以arr[r]做划分,arr[r] -> p  <p ==p   >p
返回等于区域(左边界，有边界) 所以返回一个长度为2的数组res,res[0] res[1]
*/
func partition(arr []int, l, r int) []int {
	less := l - 1  //<区右边界
	more := r      // >区左边界
	for l < more { //L表示当前熟的位置 arr[R]    -> 划分值
		if arr[l] < arr[r] { //当前书 <划分值
			lless := less + 1
			swap(arr, lless, l)
		} else if arr[l] > arr[r] { //当前数>划分值
			mmore := more - 1
			swap(arr, mmore, l)
		} else {
			l++
		}
	}
	swap(arr, more, r)
	return []int{less + 1, more}
}

func swap(arr []int, l, r int) {

}
