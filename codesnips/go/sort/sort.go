package sort

func IsSorted(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			return false
		}
	}
	return true
}

func QuickSort(nums []int) {
	doPivot(nums, 0, len(nums)-1)
}

func doPivot(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivot := lo
	i := lo + 1
	j := hi
	for {
		for ; i <= hi; i++ {
			if nums[pivot] < nums[i] {
				break
			}
		}
		for ; j > lo; j-- {
			if nums[j] < nums[pivot] {
				break
			}
		}
		if i > j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[j], nums[pivot] = nums[pivot], nums[j]
	doPivot(nums, lo, j-1)
	doPivot(nums, j+1, hi)
}
