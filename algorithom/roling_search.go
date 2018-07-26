package main

//Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
//(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
//
//You are given a target value to search. If found in the array return its index, otherwise return -1.
//
//You may assume no duplicate exists in the array.
//
//Your algorithm's runtime complexity must be in the order of O(log n).

func search(nums []int, target int) int {
    left := 0
    right := len(nums) - 1
    for left <= right {
        mid := left + (right - left) >> 1
        if nums[mid] == target {
            return mid
        }
        if nums[mid] >= nums[left] { // left is ordered
            if target < nums[mid] && target >= nums[left] {
                if target == nums[left] {
                    return left
                }
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else { // right is ordered
            if target > nums[mid] && target <= nums[right]{
                if target == nums[right] {
                    return right
                }
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }
    return -1
}
