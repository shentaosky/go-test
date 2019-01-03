package main

// 可重复
func combinationSum(candidates []int, target int) [][]int {
    var res [][]int
    for i := 0; i < len(candidates); i++ {
        if target - candidates[i] > 0 {
            tmpRes := combinationSum(candidates[i:], target - candidates[i])
            for _, tmpRes_ := range tmpRes {
                tmp := make([]int, len(tmpRes_) + 1)
                copy(tmp, tmpRes_)
                tmp[len(tmp)-1] = candidates[i]
                res = append(res, tmp)
            }
        } else if target - candidates[i] == 0 {
            res = append(res, []int{target})
        }
    }
    return res
}

// 不重复
//Example 1:
//
//Input: candidates = [10,1,2,7,6,1,5], target = 8,
//A solution set is:
//[
//  [1, 7],
//  [1, 2, 5],
//  [2, 6],
//  [1, 1, 6]
//]
//Example 2:
//
//Input: candidates = [2,5,2,1,2], target = 5,
//A solution set is:
//[
//  [1,2,2],
//  [5]
//]
func combinationSum2(candidates []int, target int) [][]int {
    var res [][]int
    for i := 0; i < len(candidates); i++ {
        if target - candidates[i] > 0 {
            tmpRes := combinationSum2(candidates[i+1:], target - candidates[i])
            for _, tmpRes_ := range tmpRes {
                tmp := make([]int, len(tmpRes_) + 1)
                copy(tmp, tmpRes_)
                tmp[len(tmp)-1] = candidates[i]
                res = append(res, tmp)
            }
        } else if target - candidates[i] == 0 {
            res = append(res, []int{target})
        }
    }
    return res
}