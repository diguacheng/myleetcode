package main

import (
	"fmt"
)

func trap1(height []int) int {
	// 暴力法 76ms 15%时间 100%内存
	ans := 0
	l := len(height)
	for i := 1; i < l-1; i++ {
		maxL, maxR := 0, 0
		for j := i; j >= 0; j-- {
			maxL = max(maxL, height[j])
		}
		for j := i; j < l; j++ {
			maxR = max(maxR, height[j])
		}
		if maxR > maxL {
			ans += maxL - height[i]
		} else {
			ans += maxR - height[i]
		}

	}
	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y

}

func trap2(height []int) int {
	// 动态编程  4ms 86.97% 3.1mb 18.52%  以空间换时间 ，将左右的最大值保存起来
	l := len(height)
	if l == 0 {
		return 0
	}
	ans := 0
	lmax := make([]int, l)
	rmax := make([]int, l)
	lmax[0] = height[0]
	for i := 1; i < l; i++ {
		lmax[i] = max(height[i], lmax[i-1]) // 保存第i个点至最左范围内的最大的值
	}
	rmax[l-1] = height[l-1]
	for i := l - 2; i >= 0; i-- {
		rmax[i] = max(height[i], rmax[i+1]) // 保存第i个点至最右边范围内最大的值
	}
	for i := 1; i < l-1; i++ {
		if lmax[i] > rmax[i] {
			ans += rmax[i] - height[i]
		} else {
			ans += lmax[i] - height[i]
		}
	}
	return ans
}

func trap3(height []int) int {
	// 使用栈 4ms 86.97% 2.8mb 25.93%  一次遍历 相当于层次填充，只要 当前高度大于栈顶（前面最近的）就放一次水块，填平低洼
	ans := 0
	curr := 0
	stack := make([]int, 0)
	for curr < len(height) {
		for len(stack) != 0 && height[curr] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if len(stack) == 0 {
				break
			}
			dis := curr - stack[len(stack)-1] - 1
			h := height[stack[len(stack)-1]]
			if height[curr] > h {
				bHeight := h - height[top]
				ans += dis * bHeight
			} else {
				bHeight := height[curr] - height[top]
				ans += dis * bHeight
			}

		}
		stack = append(stack, curr)
		curr++
	}
	return ans

}

func trap(height []int) int {
	// 双指针
	left := 0
	right := len(height) - 1
	ans := 0
	lmax, rmax := 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= lmax {
				lmax = height[left]
			} else {
				// height[left]<height[right] 且 hheight[left]<lmax 说明 left 所处的点位于一个低洼 可以防水
				ans += lmax - height[left]
			}
			left++
		} else {
			if height[right] >= rmax {
				rmax = height[right]
			} else {
				// 同理 right的点位于一个低洼处 可以防水
				ans += rmax - height[right]
			}
			right--

		}
	}
	return ans

}

func main() {
	height := []int{4, 9, 4, 5, 3, 2}
	fmt.Println(trap3(height))
	fmt.Println(trap2(height))
	fmt.Println(trap(height))

}
