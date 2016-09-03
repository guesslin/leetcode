package search

// BiSearch return the any position of nums whilch the value is equal to target
// return -1 when there is no elenemt equal to target
func BiSearch(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
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
