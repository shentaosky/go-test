package main

import "fmt"

type ServerTest struct {
    name string
    Age  string
}

func (s *ServerTest) func1() {
    fmt.Println(s.name)
}

////right
//func main() {
//  stMap := make(map[string]*ServerTest)
//  st := &ServerTest{
//     name: "st",
//     Age: "18",
//  }
//  stMap["1"] = st
//  st1 := stMap["1"]
//  fmt.Printf("%#v\n", st)
//  fmt.Printf("%#v\n", &st1)
//}
//
//// right
//func main() {
//  stMap := make(map[string]ServerTest)
//  st := ServerTest{
//      name: "st",
//  }
//  stMap["1"] = st
//  fmt.Printf("%v\n", st)
//  fmt.Printf("%v\n", stMap["1"])
//  st1 := stMap["1"]
//  st1.func1()
//}

////error
//func main() {
//  stMap := make(map[string]ServerTest)
//  st := ServerTest{
//      name: "st",
//  }
//  stMap["1"] = st
//  stMap["1"].func1()
//}

////error
//func main() {
//    stMap := []ServerTest{}
//    st := ServerTest{
//        name: "st",
//    }
//    stMap = append(stMap, st)
//    stMap[0].func1()
//}

func main() {
    stMap := make(map[string]*ServerTest)
    st := &ServerTest{
        name: "st",
    }

    stMap["1"] = st

    fmt.Printf("%p\n", st)
    fmt.Printf("%p\n", &st)

    fmt.Printf("%p\n", stMap["1"])

    tmp := stMap["1"]
    fmt.Printf("%p\n", tmp)
    fmt.Printf("%p\n", &tmp)
}

