package main

func search(nums []int, target int) bool {
	for lo, hi := 0, len(nums)-1; lo <= hi; {
		mi := (lo + hi) / 2
		if nums[mi] == target {
			return true
		} else if nums[lo] < nums[mi] || nums[hi] < nums[mi] {
			if target < nums[lo] || target > nums[mi] {
				lo = mi + 1
			} else {
				hi = mi - 1
			}
		} else if nums[mi] < nums[hi] || nums[mi] < nums[lo] {
			if target < nums[mi] || target > nums[hi] {
				hi = mi - 1
			} else {
				lo = mi + 1
			}
		} else {
			lo++
		}
	}
	return false
}
