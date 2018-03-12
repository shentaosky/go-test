package util

type Node struct {
    Left  *Node
    Right *Node
    Value int
}

type TestStruct struct {
    MCase map[string]string
    IntCase int64
    ObjectCase TestObject
    PointerCase *TestObject
    InterfaceCase interface{}
}

type TestObject struct {
    Name string
}

//type TestStruct struct {
//    testStruct
//}
//
//type testStruct struct{
//    Name string
//}
type ListNode struct {
    Val  int
    Next *ListNode
}
