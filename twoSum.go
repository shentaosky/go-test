package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"syscall"
)

func main() {
	res, err := replaceDeviceNameWithId("/dev/sda")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func replaceDeviceNameWithId(deviceName string) (string, error) {
	strs := strings.Split(deviceName, "/")
	files, err := ioutil.ReadDir("/dev/disk/by-id/")
	if err != nil {
		return "", err
	}
	deviceName = "../../" + strs[len(strs)-1]
	for _, file := range files {
		out := make([]byte, 4096)
		filePath := "/dev/disk/by-id/" + file.Name()
		lenth, err := syscall.Readlink(filePath, out)
		if err != nil {
			return "", err
		}
		if string(out)[:lenth] == deviceName {
			return filePath, nil
		}
	}
	return "", fmt.Errorf("not found")
	//args := fmt.Sprintf("ls -g /dev/disk/by-id/ | grep %s | awk {'print $8'}", deviceName)
	//out, err := exec.Command("sh", "-c", args).CombinedOutput()
	//if err != nil {
	//    return "", err
	//}
	//if string(out) == "" {
	//    return "", fmt.Errorf("err")
	//}
	//return filepath.Join("/dev/disk/by-id/", strings.Split(string(out), "\n")[0]), err
}

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}
		if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := ListNode{}
	p := &res
	q := p
	more := 0
	for l1 != nil || l2 != nil || more != 0 {
		newVal := 0
		if l1 != nil {
			newVal = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			newVal = newVal + l2.Val
			l2 = l2.Next
		}
		newVal += more
		p.Val = newVal % 10
		more = newVal / 10
		p.Next = new(ListNode)
		q = p
		p = p.Next
	}
	q.Next = nil
	return &res
}
