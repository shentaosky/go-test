package main

import "fmt"

func main() {
    nums := []int{
        0, 1, 2, 0, 1,
    }
    sortColors(nums)
    fmt.Println(nums)
}

func sortColors(nums []int)  {
    red := -1
    white := -1

    if len(nums) <= 1 {
        return
    }
    for i, num := range nums {
        if num == 0 {
            exchange(nums, white + 1, i)
            exchange(nums, red + 1, white + 1)
            white ++
            red ++
        } else if num == 1 {
            exchange(nums, white + 1, i)
            white ++
        }
    }
}

func exchange(nums []int, i, j int)  {
    tmp := nums[j]
    nums[j] = nums[i]
    nums[i] = tmp
}


