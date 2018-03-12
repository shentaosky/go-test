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
    DepthSearchAndPreorder(root)
    reverseTree(root)
    DepthSearchAndPreorder(root)
    reverseTreeNoRecrut(root)
    DepthSearchAndPreorder(root)
    //  origin tree:
    //           0
    //       1        2
    //         3    4   6
    //             5   7
}

func DepthSearchAndPreorder(root *util.Node) {
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

func reverseTree(node *util.Node) {
    if node == nil {
        return
    }
    tmpNodeLink := node.Left
    node.Left = node.Right
    node.Right = tmpNodeLink
    reverseTree(node.Left)
    reverseTree(node.Right)
}

func reverseTreeNoRecrut(root *util.Node) {
    if root == nil {
        return
    }
    l := []*util.Node{}
    l = append(l, root)
    for len(l) > 0 {
        node := l[len(l)-1]
        l = l[:len(l)-1]
        tmpNodeLink := node.Left
        node.Left = node.Right
        node.Right = tmpNodeLink
        if node.Right != nil {
            l = append(l, node.Right)
        }
        if node.Left != nil {
            l = append(l, node.Left)
        }
    }
}
