package main

import (
    "fmt"
    "os/exec"
    "strconv"
    "strings"
)

type azStruct struct {
    az string
    tmaz string
}
func main() {
    var err error
    var out []byte
    if err = exec.Command("sh", "-c", "tess kubectl config use-context tm-global").Run(); err != nil {
        fmt.Println(err)
        return
    }

    //if out, err = exec.Command("sh", "-c", "tess kubectl get availabilityZone |grep -v NAME  |awk {'print $1'}").Output(); err != nil {
    //    fmt.Println(string(out)+ ":", err.Error())
    //    return
    //}
    //azArray := strings.Split(string(out), "\n")
    //fmt.Println(azArray)
    //azMap := make(map[string]string, len(azArray))
    //for _, az := range azArray {
    //    cmd := fmt.Sprintf("tess kubectl get availabilityZone %s -o go-template --template='{{range $element := .metadata.annotations}}{{$element}}{{end}}'", az)
    //    if out, err = exec.Command("sh", "-c", cmd).Output(); err != nil {
    //        fmt.Println(string(out)+ ":", err.Error())
    //        return
    //    }
    //    if string(out) != "" &&  string(out) != " " {
    //        azMap[az] = string(out)
    //    }
    //}
    if out, err = exec.Command("sh", "-c", "tess kubectl get k8scluster |grep -v NAME  |awk {'print $1'}").Output(); err != nil {
        fmt.Println(string(out)+ ":", err.Error())
        return
    }
    clusterArray := strings.Split(string(out), "\n")
    //clusterMap := make(map[string]azStruct)
    //for _, cluster := range clusterArray {
    //    if cluster != "" && cluster != " " {
    //        cmd := fmt.Sprintf("tess kubectl get k8scluster %s -o go-template --template='{{.spec.availabilityZone.name}}'", cluster)
    //        if out, err = exec.Command("sh", "-c", cmd).Output(); err != nil {
    //            fmt.Println(string(out)+ ":", err.Error())
    //            return
    //        }
    //        az := string(out)
    //        cmd = fmt.Sprintf("tess kubectl get k8scluster %s -o go-template --template='{{range $index, $value := .metadata.labels}}{{if eq $index \"tm-az\"}}{{$value}}{{end}}{{end}}'", cluster)
    //        if out, err = exec.Command("sh", "-c", cmd).Output(); err != nil {
    //            fmt.Println(string(out)+ ":", err.Error())
    //        }
    //        tmaz := string(out)
    //        if len(tmaz) <=1 {
    //            tmaz = az
    //        }
    //        clusterMap[cluster] = azStruct{
    //            az: az,
    //            tmaz: tmaz,
    //        }
    //    }
    //}
    //fmt.Println(clusterMap)
    clusterArrayInt := []int{
        //41, 50, 120, 132, 56, 122, 17, 80, 95, 130, 30, 48, 84, 89, 93, 24, 28, 85, 29, 43, 44, 60, 92, 20, 35, 46, 47, 69, 97, 16, 42, 49, 58, 109, 113, 81, 14, 78, 99,
        70, 71, 72, 73,  74, 75, 76, 77, 78 ,79,
    }
    clusterArray = []string{}
    for _, cluster := range clusterArrayInt {
        clusterArray = append(clusterArray, strconv.Itoa(cluster))
    }
    for _, cluster := range clusterArray {
        if len(cluster) <= 1 {
            continue
        }
        cmd := fmt.Sprintf("tess kubectl config use-context %s ", cluster)
        if err = exec.Command("sh", "-c", cmd).Run(); err != nil {
            fmt.Println("use-context", cluster, err)
            continue
        }
        //cmd = fmt.Sprintf("tess kubectl get ds -n kube-system --cluster=%s node-problem-detector", cluster)
        //if out, err = exec.Command("sh", "-c", cmd).Output(); err == nil {
        //    fmt.Println("get addon npd", cluster)
        //    continue
        //}
        //cmd = fmt.Sprintf("tess kubectl get node --cluster=%s | wc -l ", cluster)
        //if out, err = exec.Command("sh", "-c", cmd).Output(); err != nil {
        //    fmt.Println("get tessmon", cluster, err, cmd)
        //    continue
        //}
        cmd = fmt.Sprintf("tess kubectl get pv --all-namespaces --cluster=%s |grep cinder | wc -l " , cluster)
        if out, err = exec.Command("sh", "-c", cmd).Output(); err != nil {
            continue
        }
        fmt.Printf("cluster %s, out: %s \n", cluster, string(out))
    }

    //for az, ep := range azMap {
    //    cmd := fmt.Sprintf("tess set cluster tm-%s --server=%s --realm=production --insecure-skip-tls-verify=true", az, ep)
    //    if out, err = exec.Command("sh", "-c", cmd).Output(); err != nil {
    //        fmt.Println(string(out)+ ":", err.Error())
    //        return
    //    }
    //    fmt.Printf("success set tm-%s \n", az)
    //}

    //exec.Command("tess kubectl get availabilityZone $(tess kubectl get k8scluster 16 -o yaml -o go-template --template='{{.spec.availabilityZone.name}}') -o yaml -o go-template --template='{{range $element := .metadata.annotations}}{{$element}} {{end}}'")
    //fmt.Println(res)

}
