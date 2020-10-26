package main

import "fmt"

type testStruct struct {
    az string
    tmaz string
}

func main() {

    testArray := []*string{}
    a := "s"
    b := "ss"
    testArray = append(testArray, &a, &b, nil)
    for i, v := range testArray {
        fmt.Printf("%d, %v", i, v)
    }
    testMap := map[string]testStruct{map[102:{ams ams} 11:{slc07 preprod} 111:{sjclab sjclab} 112:{sjclab sjclab} 12:{lvsaz01 preprod} 132:{lvsaz01 lvsaz01} 16:{slc07 slc07} 17:{lvsaz02 lvsaz02} 19:{lvsaz01 soak} 22:{lvs02 lvs02} 223:{syd syd} 23:{slc07 slc07} 231:{mia mia} 24:{rnoaz01 rnoaz01} 241:{ewr ewr} 25:{rnoaz01 rnoaz01} 258:{fra fra} 26:{rnoaz02 rnoaz02} 261:{dus dus} 27:{lvsaz01 lvsaz01} 271:{lhr lhr} 28:{rnoaz01 rnoaz01} 281:{sin sin} 29:{rnoaz02 rnoaz02} 32:{lvs02 lvs02} 33:{slc07 slc07} 34:{slc07 slc07} 35:{rnoaz03 rnoaz03} 37:{rnoaz03 rnoaz03} 38:{lvsaz02 lvsaz02} 40:{slcaz01 slcaz01} 41:{lvs02 lvs02} 42:{slc07 slc07} 43:{rnoaz02 rnoaz02} 44:{rnoaz02 rnoaz02} 45:{slcaz02 slcaz02} 46:{rnoaz03 rnoaz03} 47:{rnoaz03 rnoaz03} 48:{lvsaz03 lvsaz03} 49:{slc07 slc07} 51:{lvs02 lvs02} 52:{slc07 slc07} 53:{lvs02 lvs02} 54:{slc07 slc07} 56:{lvsaz01 lvsaz01} 58:{slc07 slc07} 60:{rnoaz02 rnoaz02} 61:{ams ams} 62:{syd syd} 64:{mdw mdw} 65:{dfw dfw} 66:{lax lax} 67:{mrs mrs} 68:{sjc sjc} 70:{slcaz01 slcaz01} 71:{rnoaz03 rnoaz03} 73:{hkg edge} 75:{lvs02 lvs02} 77:{slc07 slc07} 85:{rnoaz01 rnoaz01} 86:{rnoaz03 rnoaz03} 91:{slc07 slc07} 99:{lvsaz01 soak}]
    }
}
