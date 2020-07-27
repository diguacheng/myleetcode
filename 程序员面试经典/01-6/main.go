package main

import (
	"fmt"
	"strconv"
)

func compressString(S string) string {
	// 慢
	if len(S) <= 1 {
		return S
	}

	start := 0
	res := ""
	for i := 1; i < len(S); i++ {
		if S[i] != S[start] {
			res = res + string(S[start]) + fmt.Sprintf("%d", i-start)
			start = i
		}

	}
	res = res + string(S[start]) + fmt.Sprintf("%d", len(S)-start)
	if len(S) <= len(res) {
		return S
	}
	return res
}

func compressString1(S string) string {
	// 稍快
	if len(S) <= 1 {
		return S
	}

	start := 0
	res := make([]byte, 0)
	for i := 1; i < len(S); i++ {
		if S[i] != S[start] {
			res = append(res, S[start])
			res = append(res, []byte(fmt.Sprintf("%d", i-start))...)
			start = i
		}

	}
	res = append(res, S[start])
	res = append(res, []byte(fmt.Sprintf("%d", len(S)-start))...)
	if len(S) <= len(res) {
		return S
	}
	return string(res)
}

func compressString2(S string) string {
	// 最快
	if len(S) <= 1 {
		return S
	}

	start := 0
	res := make([]byte, 0)
	for i := 1; i < len(S); i++ {
		if S[i] != S[start] {
			res = append(res, S[start])
			res = append(res, []byte(strconv.Itoa(i-start))...)
			start = i
		}

	}
	res = append(res, S[start])
	res = append(res, []byte(strconv.Itoa(len(S)-start))...)
	if len(S) <= len(res) {
		return S
	}
	return string(res)
}

func main() {
	S := "aabcccccaaa"
	fmt.Println(compressString2(S))

}
