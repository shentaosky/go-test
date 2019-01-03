package main

import "fmt"

var phone = map[int32][]string{
    '2': {"a", "b", "c"},
    '3': {"d", "e", "f"},
    '4': {"g", "h", "i"},
    '5': {"j", "k", "l"},
    '6': {"m", "n", "o"},
    '7': {"p", "q", "s"},
    '8': {"t", "u", "v"},
    '9': {"w", "x", "y", "z"},
}

func main() {
    res := letterCombinations("23")
    fmt.Println(res)
}

func letterCombinations(digits string) []string {
    length := len(digits)
    var res = []string{}
    re := []string{}
    var search = make([][]string, length)

    for i, c := range digits {
        if chars, exist := phone[c]; exist {
            search[i] = chars
            re = append(re, chars[0])
        } else {
            return nil
        }
    }
    res = append(res, combanate(re))
    m := len(search)
    for i := 0; i < m; i++ {
        for j := 1; j < len(search[i]); j++ {
            re[i] = search[i][j]
            res = append(res, combanate(re))
        }
    }
    return res
}

func combanate(re []string) string {
    res := ""
    for _, r := range re {
        res += r
    }
    return res
}