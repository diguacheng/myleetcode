package main

import "fmt"

func findTheLongestSubstring(s string) int {
    ans, status := 0, 0
    pos := make([]int, 1 << 5)// 32个  初始化为-1
    for i := 0; i < len(pos); i++ {
		// pos的记录状态为i时，它的前缀元素的个数为多少个
        pos[i] = -1
    }
    pos[0] = 0 // 第一个元素前的元素个数为0
    for i := 0; i < len(s); i++ {
		// status用五位2进制数记录五个元素出现的奇偶
        switch s[i] {
        case 'a':
            status ^= 1 << 0
        case 'e':
            status ^= 1 << 1
        case 'i':
            status ^= 1 << 2
        case 'o':
            status ^= 1 << 3
        case 'u':
            status ^= 1 << 4
		}
		// 第一次将status对应得前缀元素个数保存保存到pos数组
		// 经过几个元素后，status变成与原来一样得，说明经过了偶数个元音字符。用当前元素个数减去前缀即得到
        if pos[status] >= 0 {
            ans = max(ans, i + 1 - pos[status])
        } else {
            pos[status] = i + 1
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

func main() {
	for i := 0; i < 5;i++{
		fmt.Printf("%d %b %b",1<<i,1<<i,(1<<i)^1)
	}

}