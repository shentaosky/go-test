package main

import(
    "fmt"
    "github.com/sirupsen/logrus"
    "github.com/sirupsen/logrus/hooks/test"
)

var log = logrus.WithFields(logrus.Fields{"pkg": "devmapper"})

func main(){
    logrus.SetLevel(logrus.DebugLevel)
    hook := test.NewLocal(log.Logger)
    log.Debugln("123xxxxxxxxxxxxxxxxx")
    fmt.Println(hook.LastEntry().Message)
    //logger.Errorln("Helloerror")
    //fmt.Println(log.Logger.Level)
}
