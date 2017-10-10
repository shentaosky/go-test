package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

//捕获因未知输入导致的程序异常
func catch(nums ...int) int {
	defer func() {
		if r := recover(); r != nil {
			log.Println("[E]", r)
			time.Sleep(time.Second)

		}
	}()
	fmt.Println("1")
	time.Sleep(time.Second)
	return nums[1] * nums[2] * nums[3] //index out of range
}

//主动抛出 panic，不推荐使用，可能会导致性能问题
func toFloat64(num string) (float64, error) {

	if num == "" {
		fmt.Println("2")
		time.Sleep(time.Second)

		panic("param is null") //主动抛出 panic,是在整个程序（不是当前函数）执行之后打印
		//并且panic只会执行一个离他最近的defer，如果该defer没有recover，整个程序将会退出。
	}

	return strconv.ParseFloat(num, 10)
}

func main() {
	defer func() {
		//即使是panic，也会执行defer,并且panic的message传递给了recover
		fmt.Println("123")
		time.Sleep(time.Second)

	}()
	catch(2, 8)
	toFloat64("")
	fmt.Println("End") //如果使用panic，就不会再往后执行，因此不会打印End
	//打印顺序不代表执行顺序
}
