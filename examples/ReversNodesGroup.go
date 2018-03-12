package main

import "fmt"
import . "../util"

func main() {
    head := &ListNode{Val: 1}
    head.Next = &ListNode{Val: 2}
    //head.Next.Next = &ListNode{Val: 3}
    //head.Next.Next.Next = &ListNode{Val: 4}
    //head.Next.Next.Next.Next = &ListNode{Val: 5}
    result := reverseKGroup(head, 2)
    for result != nil {
        fmt.Printf("%4d", result)
        result = result.Next
    }
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil || k <= 1 {
        return head
    }
    length := 0
    stack := []*ListNode{}
    for i := head; i != nil ; i = i.Next {
        stack = append(stack, i)
        length++
    }
    if length < k {
        return head
    }
    gourpNumber := length / k
    // leftout := length % k
    for n := 0; n < gourpNumber; n++ {
        for i := 1; i < k; i++ {
            stack[n*k+i].Next = stack[n*k+i-1]
        }
    }
    for n := 0; n < gourpNumber-1; n++ {
        stack[n*k].Next = stack[(n+2)*k-1]
    }
    if length % k == 0 {
        stack[(gourpNumber-1)*k].Next = nil
    } else {
        stack[(gourpNumber-1)*k].Next = stack[gourpNumber*k]
    }
    return stack[k-1]
}
