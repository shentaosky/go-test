package main

import "fmt"

const diableSystemdMount = "testmount"

func main() {
    args := []string{
        "1", "2", "3", "testmount", "4",
    }
    isMount, args := isSystemdMount(args)
    fmt.Println(isMount, args)
}

func isSystemdMount(args []string) (bool, []string) {
    if len(args) <= 0 {
        return true, args
    }
    disableIndex := -1
    for i, argument := range args {
        if argument == diableSystemdMount {
            disableIndex = i
            break
        }
    }
    if disableIndex < 0 {
        return true, args
    }
    return false, append(args[:disableIndex], args[disableIndex+1:]...)[:len(args)-1]
}
