package main

import (
    "../util"
    "fmt"
)

func main() {
    root := &util.Node{
        Value: 0,
    }
    root.Left = &util.Node{
        Value: 1,
    }
    root.Right = &util.Node{
        Value: 2,
    }
    root.Left.Right = &util.Node{
        Value: 3,
    }
    root.Right.Left = &util.Node{
        Value: 4,
    }
    root.Right.Left.Left = &util.Node{
        Value: 5,
    }
    root.Right.Right = &util.Node{
        Value: 6,
    }
    root.Right.Right.Left = &util.Node{
        Value: 7,
    }
    //widthSearch(root)
    //inorder(root)
    //depthSearchAndPreorder(root)
    //preorder(root)
    postorder(root)
    postOrdertest(root)
}

func inordertest(root *util.Node){
    if root == nil {
        return
    }
    p := root
    stack := []*util.Node{}
    stack = append(stack, p)
    for len(stack) > 0 {
        for getTop(stack, &p) && p != nil {
            stack = append(stack, p.Left)
        }
        stack = stack[:len(stack)-1]
        if len(stack) > 0 {
            p = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            fmt.Printf("%4d", p.Value)
            stack = append(stack, p.Right)
        }
    }
}

func inordertest2(root *util.Node){
    if root == nil {
        return
    }
    p := root
    stack := []*util.Node{}
    for len(stack) > 0 || p != nil {
        for p != nil {
            stack = append(stack, p)
            p = p.Left
        }
        p = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        fmt.Printf("%4d", p.Value)
        p = p.Right
    }
}

func postOrdertest(root *util.Node) {
    if root == nil {
        return
    }
    p := root
    stack := []*util.Node{}
    stack = append(stack, p)
    for p.Left != nil {
        p = p.Left
        stack = append(stack, p)
    }
    var lastRight *util.Node
    for len(stack) > 0 {
        p = stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        // 两情情况都可以打印该节点, 右边为空, 或者右边已经遍历完成.
        if p.Right == nil || p.Right == lastRight {
            fmt.Printf("%4d", p.Value)
            lastRight = p
        } else {
            // 原来出栈的再次入栈
            stack = append(stack, p)
            p = p.Right
            for p != nil {
                stack = append(stack, p)
                p = p.Left
            }
        }
    }
}


func getTop(stack []*util.Node, p **util.Node) bool{
    *p = stack[len(stack) - 1]
    //fmt.Println(p.Value)
    return true
}

// 深度优先必须用栈, 不能用队列, 访问顺序与先序遍历相同, 但是思路不同
func depthSearchAndPreorder(root *util.Node) {
    if root == nil {
        return
    }
    l := []*util.Node{}
    l = append(l, root)
    for len(l) > 0 {
        node := l[len(l)-1]
        l = l[:len(l)-1]
        fmt.Printf("%4d", node.Value)
        if node.Right != nil {
            l = append(l, node.Right)
        }
        if node.Left != nil {
            l = append(l, node.Left)
        }
    }
    fmt.Println()
}

// 广度优先必须用队列
func widthSearch(root *util.Node) {
    if root == nil {
        return
    }
    l := []*util.Node{}
    l = append(l, root)
    for len(l) > 0 {
        node := l[0]
        l = l[1:]
        fmt.Printf("%4d", node.Value)
        if node.Left != nil {
            l = append(l, node.Left)
        }
        if node.Right != nil {
            l = append(l, node.Right)
        }
    }
    fmt.Println()
}

func preorder(root *util.Node) {
    if root == nil {
        return
    }
    l := []*util.Node{}
    node := root
    for len(l) > 0 || node != nil {
        for node != nil {
            l = append(l, node)
            fmt.Printf("%4d", node.Value)
            node = node.Left
        }
        node = l[len(l)-1]
        l = l[:len(l)-1]
        node = node.Right
    }
    fmt.Println()
}

// 中序遍历的难点在于要先找到最左边的节点
func inorder(root *util.Node) {
    if root == nil {
        return
    }
    l := []*util.Node{}
    node := root
    for len(l) > 0 || node != nil {
        for node != nil {
            l = append(l, node)
            node = node.Left
        }
        node = l[len(l)-1]
        l = l[:len(l)-1]
        fmt.Printf("%4d", node.Value)
        node = node.Right
    }
    fmt.Println()
}

// 后序是最难的, 思路是先找最左边的入栈, 然后右孩子是不是访问过或者为空, 如果不是就继续访问右孩子的左子树循环.
func postorder(root *util.Node) {
    if root == nil {
        return
    }
    l := []*util.Node{}
    node := root
    lastRight := &util.Node{}
    for node != nil {
        l = append(l, node)
        node = node.Left
    }
    for len(l) > 0 {
        node = l[len(l)-1]
        l = l[:len(l)-1]
        if node.Right == nil || node.Right == lastRight {
            fmt.Printf("%4d", node.Value)
            lastRight = node
        } else {
            l = append(l, node)
            node = node.Right
            for node != nil {
                l = append(l, node)
                node = node.Left
            }
        }
    }
    fmt.Println()
}



