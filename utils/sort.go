package utils

func BubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func MergeSort(nums []int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}
	middle := length / 2
	return merge(MergeSort(nums[:middle]), MergeSort(nums[middle:]))
}
func merge(left, right []int) []int {
	res := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			res[k] = left[i]
			i++
		} else {
			res[k] = right[j]
			j++
		}
		k++
	}
	if i < len(left) {
		for ; i < len(left); i++ {
			res[k] = left[i]
			k++
		}
	}
	if j < len(right) {
		for ; j < len(right); j++ {
			res[k] = right[j]
			k++
		}
	}
	return res
}

func QuickSort(nums []int) {
}
