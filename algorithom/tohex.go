//https://leetcode.com/problems/convert-a-number-to-hexadecimal/description/
//Given an integer, write an algorithm to convert it to hexadecimal. For negative integer, two’s complement method is used.
//
//Note:
//
//All letters in hexadecimal (a-f) must be in lowercase.
//The hexadecimal string must not contain extra leading 0s. If the number is zero, it is represented by a single zero character '0'; otherwise, the first character in the hexadecimal string will not be the zero character.
//The given number is guaranteed to fit within the range of a 32-bit signed integer.
//You must not use any method provided by the library which converts/formats the number to hex directly.
//Example 1:
//
//Input:
//26
//
//Output:
//"1a"
//Example 2:
//
//Input:
//-1
//
//Output:
//Given an integer, write an algorithm to convert it to hexadecimal. For negative integer, two’s complement method is used.
//
//Note:
//
//All letters in hexadecimal (a-f) must be in lowercase.
//The hexadecimal string must not contain extra leading 0s. If the number is zero, it is represented by a single zero character '0'; otherwise, the first character in the hexadecimal string will not be the zero character.
//The given number is guaranteed to fit within the range of a 32-bit signed integer.
//You must not use any method provided by the library which converts/formats the number to hex directly.
//Example 1:
//
//Input:
//26
//
//Output:
//"1a"
//Example 2:
//
//Input:
//-1
//
//Output:
//"ffffffff"

package main

import (
    "fmt"
)

const hex string = "0123456789abcdef"

func main() {
    //fmt.Printf("%o", -26)
    fmt.Println(toHex(-26))
    //moveOpsTest(-10)
    fmt.Println(toHexSimple(-32))
}

// 负数右移补1, 正数右移补0, java有无符号右移'>>>' 操作, c和go中没有.
func moveOpsTest(num int) {
    var count int
    for num != 0 && count < 30{
        fmt.Println(num)
        fmt.Printf("%b \n", num)
        num >>= 1
        count++
    }
}

func toHexSimple(num int) string {
    return fmt.Sprintf("%b", uint32(num))
}

func toHex(num int) string {
    var result string
    var count int
    if num == 0 {
        return "0"
    }
    for num != 0 && count < 8 {
        result = string(hex[num&0xf]) + result
        num = num >> 4
        count++
    }
    return result
}

// string不能存某个索引, 只能取
func reverseString(s string) string {
    // s[1] = '1' error
    // a := s[1] right
    runes := []rune(s)
    for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
        runes[from], runes[to] = runes[to], runes[from]
    }
    return string(runes)
}