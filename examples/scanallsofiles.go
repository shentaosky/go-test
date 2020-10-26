package main

import (
    "fmt"
    "os/exec"
    "strings"
)
var result = map[string]struct{}{}

func main() {
    result = make(map[string]struct{}, 2000)
    scan("/usr/bin/rbd")
    for res := range result {
        if res == "" {
            continue
        }
        out, err := exec.Command("realpath", res).CombinedOutput()
        if err != nil {
           fmt.Printf("out: %s, err: %v", string(out), err)
        }
        start := strings.LastIndex(string(out), "/")
        end := strings.LastIndex(string(out), "\n")
        fmt.Printf(" %s", string(out)[start+1:end])
        //fmt.Printf(" %s", res[start+1:])
    }
}

func scan(path string) {
    out, err := exec.Command("ldd", path).CombinedOutput()
    if err != nil {
        fmt.Printf("out: %s, err: %v", string(out), err)
    }
    res := string(out)
    resSplits := strings.Split(res, "\n")
    resultNew := make([]string, 200)
    for i, resS := range resSplits {
        if i == 0 || i == len(resSplits) - 1 {
            continue
        }
        start := strings.LastIndex(resS, ">") + 2
        end := strings.LastIndex(resS, "(") - 1
        if _, ok := result[resS[start: end]]; !ok {
            resultNew = append(resultNew, resS[start: end])
        }
        result[resS[start: end]] = struct{}{}
    }
    for _, newStr := range resultNew {
        if newStr != "" {
            scan(newStr)
        }
    }
}