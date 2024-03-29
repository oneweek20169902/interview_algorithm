package recursive

import (
	"fmt"
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
func DutchFlag(arr []int, num int) []int {
	/*
		1. i<num, i和<区域下一个交换，<区域右扩，i++
		2. i==num, i++
		3.i>num,i和>区域前一个交换，>区域右扩,i原地不动
	*/
	i := 0
	less := -1
	more := len(arr) - 1
	for i <= more {
		if arr[i] < num {
			less += 1
			arr[i], arr[less] = arr[less], arr[i]
			i++
		} else if arr[i] == num {
			i++
		} else {
			arr[i], arr[more] = arr[more], arr[i]
			more--
		}
	}
	return []int{arr[i] - 1, arr[more]}
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

/*
给定一个包含红色、白色和蓝色、共n 个元素的数组nums，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

我们使用整数 0、1 和 2 分别表示红色、白色和蓝色。

必须在不使用库内置的 sort 函数的情况下解决这个问题

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/sort-colors
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func sortColors(nums []int) {
	p0, p1 := 0, 0
	for i, v := range nums {
		if v == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			if p0 < p1 {
				nums[p1], nums[i] = nums[i], nums[p1]
			}
			p0++
			p1++
		}
		if v == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		}
	}
}

/*
左 2*i+1
右 2*i+2
父 (i-1)/2
大根堆 每一个子树的头结点都是最大值
小根堆 每一个子树的头结点都是最小值
*/

func heapify(nums []int, parent, len int) {
	max := parent
	l := parent*2 + 1
	r := parent*2 + 2
	if l < len && nums[l] > nums[max] {
		max = l
	}
	if r < len && nums[r] > nums[max] {
		max = r
	}
	if parent != max {
		nums[max], nums[parent] = nums[parent], nums[max]
		heapify(nums, max, len)
	}
}

func buildHeep(nums []int, len int) {
	parent := len/2 - 1
	for parent >= 0 {
		heapify(nums, parent, len)
		parent--
	}
	fmt.Println(nums[0:len])
}

func HeapSort(nums []int) []int {
	i := len(nums)
	for i > 1 {
		buildHeep(nums, i)
		nums[0], nums[i-1] = nums[i-1], nums[0]
		i--
	}
	return nums
}

/*
·如果找到了1，那么将其与nums【p1】进行交换，并将p1向后移动一个位置，这与方法一是相同的；
·如果找到了0，那么将其与nums【po】进行交换，并将po向后移动一个位置。这样做是正确的吗？
我们可以注意到，因为连续的0之后是连续的1，因此如果我们将0与nums【po】进行交换，那么我们可能会把一个1交换出去。当p0<p1时，我们已经将一些1连续地放在头部，此时一定会把一个1交换出去，导致答案错误。
因此，如果p0<p1，那么我们需要再将nums【i与numsp1进行交换，其中i是当前遍历到的位置，在进行了第一次交换后，nums[i]的值为1，我们需要将这个1放到「头部】的末端。在最后，无论是否有p0<p1，
我们需要将p0和p1均向后移动一个位置，而不是仅将p0向后移动一个位置
*/
func sortColors1(nums []int) {
	p1, p2 := 0, 0
	for i, v := range nums {
		if v == 0 {
			nums[p1], nums[i] = nums[i], nums[p1]
			//if p1 < p2 { //当p0<p1时，我们已经将一些1连续地放在头部，此时一定会把一个1交换出去，导致答案错误。
			//	nums[p2], nums[i] = nums[i], nums[p2]
			//}
			p2++ //如果P2指针不走，在index对应的value满足p2索引条件时会替换掉p1指针已经处理过的值
			p1++
		}

		if v == 1 {
			nums[p2], nums[i] = nums[i], nums[p2]
			p2++
		}
		fmt.Println(nums)
	}
}
