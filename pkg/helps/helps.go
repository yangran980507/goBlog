// Package helps 公共辅助函数
package helps

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
	case length - 1:
		newNums = nums[:length-1]
	default:
		newNums = append(nums[:index], nums[index+1:]...)
	}
	return
}
