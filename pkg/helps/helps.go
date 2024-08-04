// Package helps 公共辅助函数
package helps

import (
	"crypto/rand"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"slices"
	"strconv"
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

// RandomQuantity 随机数量生成
func RandomQuantity() int {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, 2)
	n, err := io.ReadAtLeast(rand.Reader, b, 2)
	if n != 2 {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	num, _ := strconv.Atoi(string(b))
	return num
}

// RandomPrice 随机价格生成
func RandomPrice() float64 {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, 4)
	n, err := io.ReadAtLeast(rand.Reader, b, 4)
	if n != 4 {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}

	priceStr := fmt.Sprintf("%s.%s", string(b[0:2]), string(b[2:]))
	return cast.ToFloat64(priceStr)
}

// RandomPublisher 随机出版社生成
func RandomPublisher() string {
	table := [...]string{"people出版社", "china书局", "五联书店",
		"前景印书馆", "铸刻文化", "海洋出版社", "童趣出版社", "乾南文化馆", "梧桐漫画书屋", "随笔毛斋"}
	b := make([]byte, 1)
	var word string
	n, err := io.ReadAtLeast(rand.Reader, b, 1)
	if n != 1 {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		word = table[int(b[i])%len(table)]
	}
	return word
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
