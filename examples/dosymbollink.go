package main

import (
    "fmt"
    "os"
    "path/filepath"
    "regexp"
)

func main() {
    //if err :=os.MkdirAll("/mnt/test/dir1", 777); err != nil {
    //    fmt.Println(err)
    //    return
    //}
    //if err := doSymlink("/dev/mapper/vg10000-vg10000_pvc--7aac5eff3a3311e9notMnt", "/mnt/test/dir1"); err != nil {
    //    fmt.Println(err)
    //}
    //res, err := filepath.EvalSymlinks("/Users/tashen/work/devs/src/github.com/go-test/test1/test4")
    //if err != nil {
    //    fmt.Println(err)
    //    return
    //}
    //fmt.Println(res)
    disk := "/dev/nvme12n123"
    if loc := regexp.MustCompile("[0-9]+").FindStringIndex(disk); loc != nil {
        res := disk[:loc[1]]
        fmt.Println(res, loc)
    }
    fmt.Println(filepath.Join("/", "dev", "nvme"))
}

func doSymlink(sourcePath, linkFile string) error {
    mapPath := filepath.Dir(linkFile)
    if !filepath.IsAbs(mapPath) {
        return fmt.Errorf("The map path should be absolute: map path: %s", mapPath)
    }

    // Check and create mapPath
    _, err := os.Stat(mapPath)
    if err != nil && !os.IsNotExist(err) {
        return err
    }
    if err = os.MkdirAll(mapPath, 0750); err != nil {
        return fmt.Errorf("failed to mkdir %s: %v", mapPath, err)
    }
    // Remove old symbolic link(or file) then create new one.
    // This should be done because current symbolic link is
    // stale across node reboot.
    if err = os.Remove(linkFile); err != nil && !os.IsNotExist(err) {
        return err
    }
    err = os.Symlink(sourcePath, linkFile)
    return err
}
