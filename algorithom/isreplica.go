package main

import "fmt"

func main() {
    test := []int{
        3, 1, 2, 5, 4, 0,
    }
    fmt.Println(isReplica(test))
}
// [3   1   2   5   3   0]
func isReplica(num []int) bool {
    if len(num) < 1 {
        return false
    }
    j := 0
    for i := 0; i < len(num); i++ {
        if i == num[i] {
            j++
            continue
        }
        for num[j] != j {
            if num[num[j]] == num[j] {
                return false
            }
            swap1(num, j, num[j])
        }
        j++
    }
    return true
}

func swap1(num []int, i, j int) {
    num[i] = num[i] ^ num[j]
    num[j] = num[i] ^ num[j]
    num[i] = num[i] ^ num[j]
}

