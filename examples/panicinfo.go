package main

import (
    "errors"
)

type X struct {
    i int
}

type Y struct {
}

func (y *Y) foo(x *X) {
    panic("panic in foo")
}

func (y *Y) bar(x *X) (*Y) {
    panic("panic in bar")
    return y
}

func (y *Y) baz(x *X) (error) {
    panic("panic in baz")
    return errors.New("error in baz")
}

func (y *Y) bam() {
    panic("panic in bam")
}

func main() {
    //defer func() {
    //    out := recover()
    //    fmt.Println(out)
    //}()
    y := new(Y)
    x := new(X)
    // comment out the ones you don't want to check
    y.foo(x)
    y.bar(x)
    y.baz(x)
    y.bam()
}