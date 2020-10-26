package main

import (
    "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
    "os/exec"
)

func main() {
    out, err := exec.Command("kubectl get pods -o wide -n kube-system |grep dynamic | awk {'print $1'}").CombinedOutput()
    if err != nil {
        fmt.Println(err)
        return
    }

    exec.Command("for i in $(kubectl get pods -o wide -n kube-system |grep dynamic | awk {'print $1'} ); do echo $i; kubectl exec -it $i -n kube-system -c localcsi -- lvs")
}
