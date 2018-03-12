package main

import (
    "fmt"

    "../util"
)

func main() {
    testCase := &util.TestStruct{
        IntCase: 1,
        MCase: map[string]string{
            "1": "1",
        },
        PointerCase: &util.TestObject{},
    }
    fmt.Printf("%v", testCase.PointerCase.Name) // not panic

    testCase.PointerCase = nil   // pointer could be nil, then

    fmt.Printf("%v", testCase.PointerCase.Name) // will panic

    // testCase.ObjectCase = nil // struct can't be nil

    testCase.ObjectCase = util.TestObject{}  // empty struct

    fmt.Printf("%v", testCase.ObjectCase.Name) // not panic
}


