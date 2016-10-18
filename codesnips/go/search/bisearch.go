package search

// BiSearch return the any position of nums whilch the value is equal to target
// return -1 when there is no elenemt equal to target
func BiSearch(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// UpperBound return the number of element is smaller or equal than target in list
func UpperBound(nums []int, target int) int {
	s, e := 0, len(nums)-1
	for s <= e {
		mid := (s + e) / 2
		if nums[mid] > target {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	return s
}

// LowerBound return the number of element is smaller than target in list
func LowerBound(nums []int, target int) int {
	s, e := 0, len(nums)-1
	for s <= e {
		mid := (s + e) / 2
		if nums[mid] >= target {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	return s
}
