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
