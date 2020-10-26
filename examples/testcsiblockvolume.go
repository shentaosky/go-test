package main

import (
    "fmt"
    "io"
    "math/rand"
    "os/exec"
    "time"
)

func main() {
    testcliams1 := make([]string, 20)
    testcliams2 := make([]string, 20)
    testpods := make([]string, 20)
    for i := 0; i < 20 ; i++ {
        testpods[i] = newPod(i)
        testcliams1[i] = newCliam1(i)
        testcliams2[i] = newCliam2(i)
    }
    t := time.NewTimer(10*time.Minute)
    for {
        select{
            case <- t.C:
                fmt.Println("done")
                return
        default:
            for i := 0; i < 10; i++ {
                str1, err1 := kubeCreate1(testcliams1[i])
                fmt.Println(str1, err1)
                str1, err1 = kubeCreate1(testcliams2[i])
                fmt.Println(str1, err1)
                str2, err2 := kubeCreate1(testpods[i])
                fmt.Println(str2, err2)
            }

            time.Sleep(time.Duration(rand.Intn(100))*time.Second)
            str, err := kubeDeletePo()
            fmt.Println(str, err)
            time.Sleep(time.Duration(rand.Intn(20))*time.Second)
            str, err = kubeDeletePvc()
            fmt.Println(str, err)
            time.Sleep(time.Duration(rand.Intn(100))*time.Second)
        }

    }
}

func kubeCreate1(yamlContent string) (string, error) {
    cmd := exec.Command("kubectl", "create", "-f", "-")
    stdin, err := cmd.StdinPipe()
    if err != nil {
        return "", err
    }
    go func() {
        defer stdin.Close()
        io.WriteString(stdin, yamlContent)
    }()
    out, err := cmd.CombinedOutput()
    return string(out), err
}

func kubeDeletePo() (string, error) {
    cmd := exec.Command("kubectl", "delete", "pod", "--all", "--namespace", "shentao")
    out, err := cmd.CombinedOutput()
    return string(out), err
}

func kubeDeletePvc() (string, error) {
    cmd := exec.Command("kubectl", "delete", "pvc", "--all", "--namespace", "shentao")
    out, err := cmd.CombinedOutput()
    return string(out), err
}

func newPod(num int) string {
    return fmt.Sprintf("apiVersion: v1 \n" +
        "kind: Pod \n" +
        "metadata: \n" +
        "  name: testpod%d \n" +
        "  namespace: shentao \n" +
        "spec: \n" +
        "  containers: \n" +
        "  - image: busybox \n" +
        "    command: \n" +
        "      - sleep \n" +
        "      - \"3600\" \n" +
        "    imagePullPolicy: IfNotPresent \n" +
        "    name: test \n" +
        "    volumeDevices: \n" +
        "    - devicePath: /mnt/disk1 \n" +
        "      name: test-vol1 \n" +
        "    volumeMounts: \n" +
        "    - mountPath: /mnt/volume1 \n" +
        "      name: test-vol2 \n" +
        "  volumes: \n" +
        "    - persistentVolumeClaim: \n" +
        "          claimName: testpvc%d \n" +
        "      name: test-vol1 \n" +
        "    - persistentVolumeClaim: \n" +
        "          claimName: testcinder%d \n" +
        "      name: test-vol2 \n" +
        "  restartPolicy: Always \n", num, num, num)
}

func newCliam1(num int) string {
    return fmt.Sprintf("kind: PersistentVolumeClaim \n" +
                        "apiVersion: v1 \n" +
                        "metadata: \n" +
                        "  name: testpvc%d \n" +
                        "  namespace: shentao\n" +
                        "spec: \n" +
                        "  storageClassName: local-dynamic \n" +
                        "  accessModes: \n" +
                        "    - ReadWriteOnce \n" +
                        "  resources: \n" +
                        "    requests: \n" +
                        "      storage: 1Gi \n" +
                        "  volumeMode: Block \n", num)
}

func newCliam2(num int) string {
    return fmt.Sprintf("kind: PersistentVolumeClaim \n" +
        "apiVersion: v1 \n" +
        "metadata: \n" +
        "  name: testcinder%d \n" +
        "  namespace: shentao\n" +
        "spec: \n" +
        "  storageClassName: cinder-standard \n" +
        "  accessModes: \n" +
        "    - ReadWriteOnce \n" +
        "  resources: \n" +
        "    requests: \n" +
        "      storage: 1Gi \n", num)
}
