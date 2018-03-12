package main

import (
    "github.com/docker/docker/pkg/devicemapper"
    "fmt"
)

func main() {
    startSector, totalSector, targetType, params, err := devicemapper.GetStatus("/dev/mapper/pvc-e3d486fc-d66c-11e7-bcec-b083feb83778")
    if err != nil {
        fmt.Printf("%v", err)
        return
    }

    var TransactionID, UsedMetaBlks uint64
    _, err = fmt.Sscanf(params, "%d %d", &TransactionID, &UsedMetaBlks)
    if err != nil {
        fmt.Printf("%v\n", err)
    }

    fmt.Printf("%d, %d", TransactionID, UsedMetaBlks)
}
