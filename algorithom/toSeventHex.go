//给定一个整数，将其转化为7进制，并以字符串形式输出。
//
//示例 1:
//
//输入: 100
//输出: "202"
//示例 2:
//
//输入: -7
//输出: "-10"
package main

import (
    "strconv"
    "fmt"
)

func main() {
    fmt.Println(convertToBase7(-100))
}

func convertToBase7(num int) string {
    var result string
    var symbol string
    if num == 0 {
        return "0"
    }
    if num < 0 {
        symbol = "-"
        num *= -1
    }
    for num != 0 {
        result = strconv.Itoa(num % 7) + result
        num = num / 7
    }
    return symbol + result
}
