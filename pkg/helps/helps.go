// Package helps 公共辅助函数
package helps

import (
	"crypto/rand"
	"io"
	"slices"
)

// JudgeElementInSliceExist 判断数组是否含指定有元素
func JudgeElementInSliceExist(num int64, nums []int64) bool {
	for _, v := range nums {
		if num == v {
			return true
		}
	}
	return false
}

// DeleteElementInSliceExist 判断数组是否含指定有元素
func DeleteElementInSliceExist(index int64, nums []int64) (newNums []int64) {
	length := int64(len(nums))

	switch index {
	case 0:
		newNums = nums[1:length]
	case length - 1:
		newNums = nums[:length-1]
	default:
		newNums = append(nums[:index], nums[index+1:]...)
	}
	return
}

// RandomNumber 随机数字生成
func RandomNumber(length int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(rand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// GenerateNewSliceByDeleteOldSlice 删除切片1中给出的切片2的值
func GenerateNewSliceByDeleteOldSlice(original, condition []int64) (result []int64) {
	result = slices.DeleteFunc(original, func(i int64) bool {
		var answer bool
		for _, v := range condition {
			if v == i {
				answer = true
				break
			}
		}
		return answer
	})
	return
}
