package main

import (
    "context"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os/exec"
    "strconv"
    "strings"
    "time"
)

type MdsResponse struct {
    Request MdsRequestStruct `json:request,omitempty`
}

type MdsRequestStruct struct {
    Tid int64 `json:"tid,omitempty"`
    Age float64 `json:"age,omitempty"`
    Mds int    `json:"mds,omitempty"`
}

type ObjectRequest struct {
    Ops []ObjectRequestStuct
}

type ObjectRequestStuct struct {
    LastSend string `json:"last_sent,omitempty"`
    Osd      int `json:"osd,omitempty"`
}

type ObjectRes struct {
    Age float64
    Count int
}

func main() {
    //file, err := ioutil.ReadFile("/Users/tashen/mds_requests.log")
    file, err := ioutil.ReadFile("examples/cephfusemds.json")

    if err != nil {
        fmt.Println(err)
        return
    }
    data := string(file)

    abnormalCnt := 0
    maxAge := float64(0)
    maxAgeMds := -1

    lastMdsIndex := 0
    for true {
        mdsIndex := strings.Index(data[lastMdsIndex+7:], "\"mds\":")
        if mdsIndex == -1 {
            fmt.Println(lastMdsIndex)
            break
        }
        mdsIndex += lastMdsIndex + 7
        ageIndex := strings.Index(data[lastMdsIndex:mdsIndex], "\"age\":")
        if ageIndex == -1 {
            abnormalCnt += 1
            lastMdsIndex = mdsIndex
            continue
        }
        ageIndex += lastMdsIndex
        lastMdsIndex = mdsIndex
        //fmt.Println(ageIndex, strings.Index(data[ageIndex+7:], ","))
        age, err := strconv.ParseFloat(data[ageIndex+7:ageIndex+7+strings.Index(data[ageIndex+7:], ",")], 64)
        if err != nil {
            fmt.Println(data[ageIndex+7:ageIndex+7+strings.Index(data[ageIndex+7:], ",")], err)
        }
        if age > maxAge {
            mds, err := strconv.Atoi(data[mdsIndex+7:mdsIndex+7+strings.Index(data[mdsIndex+7:], ",")])
            if err != nil {
                fmt.Println(data[mdsIndex+7:mdsIndex+7+strings.Index(data[mdsIndex+7:], ",")], err)
                continue
            }
            maxAge = age
            maxAgeMds = mds
        }
    }

    fmt.Println(maxAgeMds, maxAge, abnormalCnt)

    out, err := ioutil.ReadFile("/Users/tashen/objecter_requests.log")
    if err != nil {
        fmt.Println(err)
        return
    }

    res := make(map[int]*ObjectRes)
    var v ObjectRequest
    err = json.Unmarshal(out, &v)
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, op := range v.Ops {
        age, err := strconv.ParseFloat(op.LastSend[:strings.Index(op.LastSend, "s")], 64)
        if err != nil {
            fmt.Println(err)
            return
        }
        r, ok := res[op.Osd]
        if !ok {
            res[op.Osd] = &ObjectRes{
                Age: age,
                Count: 1,
            }
        } else if r.Age < age {
            r.Count++
            r.Age = age
        } else {
            r.Count++
        }
    }

    for k, v := range res {
        fmt.Println(k, v.Age, v.Count)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
    defer cancel()

    if err := exec.CommandContext(ctx, "sleep", "2").Run(); err != nil {
        fmt.Println(err)
    }
    ctx, cancel = context.WithTimeout(context.Background(), 4*time.Second)

    if err := exec.CommandContext(ctx, "sleep", "2").Run(); err != nil {
        fmt.Println(err)
    }
    ctx, cancel = context.WithTimeout(context.Background(), 4*time.Second)

    if err := exec.CommandContext(ctx, "sleep", "2").Run(); err != nil {
        fmt.Println(err)
    }
    ctx, cancel = context.WithTimeout(context.Background(), 4*time.Second)

    if err := exec.CommandContext(ctx, "sleep", "2").Run(); err != nil {
        fmt.Println(err)
    }
    select {
    case <-ctx.Done():
        fmt.Println("done")
    }




    //
    //mdsResponse := &MdsResponse{
    //}
    //if err := json.Unmarshal(file, mdsResponse); err != nil {
    //    fmt.Println(err)
    //    return
    //}
    //fmt.Println(mdsResponse.Request)
    //fmt.Println("done")

}