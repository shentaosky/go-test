package main

import (
	"fmt"
	"k8s.io/api/apps/v1"

	//"io"
	//"log"
	"io"
	"math/rand"
	"os/exec"
	"strings"
	"time"
)

const (
	BINARY_DIR = "/root/convoy_test/kubernetes/bin/"
	APISERVER  = BINARY_DIR + "kube-apiserver"
	MANAGER    = BINARY_DIR + "kube-controller-manager"
	PROXY      = BINARY_DIR + "kube-proxy"
	SCHEDULER  = BINARY_DIR + "kube-scheduler"
	KUBELETE   = BINARY_DIR + "kubelet"
	ETCD       = BINARY_DIR + "etcd"
	KUBECTL    = "/home/shentao/gopath/src/k8s.io/kubernetes/_output/bin/kubectl"
	poolname   = "bronze"
	pod        = "podtest"
	petset     = "petsettest"
	RUNTIME    = 90
	NOWAIT     = 0
	WAIT       = 1
	RANDOMWAIT = 2
	RETRY      = 500
)

type k8sAcceptanceSuit struct {
	root          string
	dataDir       string
	logDir        string
	containerId   string
	dataDev       string
	metaDev       string
	etcdPid       int
	kubeletPid    int
	proxyPid      int
	managerPid    int
	schedulerPid  int
	apiserverPid  int
	etcdArgs      []string
	kubeletArgs   []string
	proxyArgs     []string
	managerArgs   []string
	schedulerArgs []string
	apiserverArgs []string
}

func main() {
	//tosdiskYaml := fmt.Sprintf("kind: StorageClass \n" +
	//        "apiVersion: storage.k8s.io/v1beta1 \n" +
	//        "metadata: \n" +
	//        "  name: %s \n" +
	//        "provisioner: transwarp.io/tosdisk", POOLNAME)
	//err := kubeCreate(tosdiskYaml)
	//
	//petsetYaml := fmt.Sprintf("apiVersion: apps/v1beta1 \n" +
	//        "kind: StatefulSet \n" +
	//        "metadata: \n" +
	//        "  name: %s \n" +
	//        "spec: \n" +
	//        "  serviceName: \"nginx\" \n" +
	//        "  replicas: 1 \n" +
	//        "  template: \n" +
	//        "    metadata: \n" +
	//        "      labels: \n" +
	//        "        app: nginx \n" +
	//        "      annotations: \n" +
	//        "        pod.alpha.kubernetes.io/initialized: \"true\" \n" +
	//        "    spec: \n" +
	//        "      terminationGracePeriodSeconds: 0 \n" +
	//        "      containers: \n" +
	//        "      - name: nginx \n" +
	//        "        image: 172.16.1.41:5000/transwarp-centos6:20150311-01 \n" +
	//        "        args: \n" +
	//        "          - \"/bin/bash\" \n" +
	//        "          - \"-c\" \n" +
	//        "          - \"echo REPLICAID: $REPLICAS; while true;do sleep 1;done\" \n" +
	//        "        env: \n" +
	//        "          - name: REPLICAS \n" +
	//        "            valueFrom: \n" +
	//        "              fieldRef: \n" +
	//        "                fieldPath: metadata.annotations.transwarp.replicaid \n" +
	//        "        volumeMounts: \n" +
	//        "        - name: www \n" +
	//        "          mountPath: /usr/share/nginx/html \n" +
	//        "  volumeClaimTemplates: \n" +
	//        "  - metadata: \n" +
	//        "      name: www \n" +
	//        "      annotations: \n" +
	//        "        volume.beta.kubernetes.io/storage-class: %s \n" +
	//        "    spec: \n" +
	//        "      accessModes: [ \"ReadWriteOnce\" ] \n" +
	//        "      resources: \n" +
	//        "        requests: \n" +
	//        "          storage: 1Gi \n" +
	//        "        limits: \n" +
	//        "          blkio.throttle.read_iops_device: 100 \n" +
	//        "          blkio.throttle.write_iops_device: 200 \n" +
	//        "          blkio.throttle.read_bps_device: 100M \n" +
	//        "          blkio.throttle.write_bps_device: 200M \n", "web", POOLNAME)
	//err = kubeCreate(petsetYaml)
	//if err != nil {
	//    fmt.Println(err)
	//}
	//waitStartContainer("pod", "web-0")

	//out, err := exec.Command("sh", "-c", "/home/shentao/gopath/src/k8s.io/kubernetes/_output/bin/kubectl describe pods web-0 | grep ClaimName").Output()
	//if err != nil {
	//    fmt.Println(err)
	//} else {
	//    res := strings.Split(string(out), ":")[1]
	//    res = strings.Replace(res, "\t", "", -1)
	//    res = strings.Replace(res, "\n", "", -1)
	//    if res == "www-web-0" {
	//        fmt.Println(res)
	//    }
	//}
	//str := "    str   str str"
	//str = strings.Replace(str, " ", "", -1)
	//fmt.Println(str)
	////waitStartContainer()
	//waitPVCDeleted()
	//err := exec.Command(KUBECTL, "get", "pods").Run()
	//if err == nil {
	//    fmt.Println("done")
	//} else {
	//    fmt.Printf("%v", err)
	//}
	//stop := make(chan struct{}, 2)
	//go stop1(stop)
	//go stop2(stop)
	//time.Sleep(time.Second * 2)
	//stop <- struct {}{}
	//stop <- struct {}{}
	//time.Sleep(time.Second * 1)
	//label := "123"
	//podYaml := generatePetSetYaml("1", label)
	//err := kubeCreate(podYaml)
	//if err != nil {
	//    fmt.Printf("%v", err)
	//    fmt.Print(podYaml)
	//} else {
	//    println("123")
	//}
	//str := "1\n2\n3\n"
	//println(str)
	//list := strings.Split(str, "\n")
	//if list[3]=="" {
	//    println(list)
	//}
	//waitStartStatefulSet("4d6", 1

}

func (s *k8sAcceptanceSuit) tosdiskCreateAndDelete(petsetYaml, petsetname, label string, replica, delPetSetMode, delPVCMode, createMode int, stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			return
		default:
			count := 0

			for {
				if err := kubeCreate(petsetYaml); err == nil {
					break
				} else {
					count++
					if count > 10 {
						fmt.Println("err kubeCreate")
					}
				}
				time.Sleep(time.Millisecond * 200)
			}

			if delPetSetMode == NOWAIT {
			} else if delPetSetMode == RANDOMWAIT {
				time.Sleep(time.Second * time.Duration(rand.Intn(25)))
			} else if delPetSetMode == WAIT {
				waitStartStatefulSet(label, replica)
			}

			//c.Errorf("after waitStartStatefulSet")

			count = 0

			for {
				if err := exec.Command(KUBECTL, "delete", "statefulset", petsetname).Run(); err == nil {
					break
				} else {
					count++
					if count > 10 {
						fmt.Println("err kubeCreate")
					}
				}
				time.Sleep(time.Millisecond * 200)
			}

			if delPVCMode == NOWAIT {
			} else if delPVCMode == RANDOMWAIT {
				time.Sleep(time.Second * time.Duration(rand.Intn(5)))
			} else if delPVCMode == WAIT {
				waitUnmountVolume(s, c)
			}

			//c.Errorf("after waitUnmountVolume")

			count = 0

			for {
				if err := exec.Command("sh", "-c", fmt.Sprintf("%s delete pvc $(%s get pvc -l test=%s | awk {'print $1'} |grep -v NAME)", KUBECTL, KUBECTL, label)).Run(); err == nil {
					break
				} else {
					count++
					if count > 10 {
						fmt.Println("err kubeCreate")
					}
				}
				time.Sleep(time.Millisecond * 200)
			}

			if createMode == NOWAIT {
			} else if createMode == RANDOMWAIT {
				time.Sleep(time.Second * time.Duration(rand.Intn(15)))
			} else if createMode == WAIT {
				waitConvoyCleanUp(s, c)
			} else {
				c.Errorf("error create mode")
			}
		}

		//c.Errorf("after waitConvoyCleanUp")
	}
}

func (s *k8sAcceptanceSuit) restartConvoy() {
	t := time.NewTicker(time.Second * 1000)
	for {
		select {
		case <-t.C:
			return
		default:
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(25000)))
			err := exec.Command("systemctl", "restart", "convoy").Run()
			fmt.Println(err)
		}
	}
}

func waitUnmountVolume(s *k8sAcceptanceSuit) {
	t1 := time.NewTicker(time.Second * TIMESTEP)
	t2 := time.NewTicker(time.Second * TIMEOUT)

	for {
		select {
		case <-t1.C:
			volumeResps, err := s.convoyClient.ListDevices()
			c.Assert(err, IsNil)
			i := 0
			for _, volumeInfo := range volumeResps.Items {
				if volumeInfo.ThinPoolVolume.MountPoint == "" {
					i++
				}
			}
			if i == len(volumeResps.Items) {
				return
			}
		case <-t2.C:
			return
		}
	}
}

func waitStartStatefulSet(label string, replica int) {
	t1 := time.NewTicker(time.Second * 1)
	t2 := time.NewTicker(time.Second * 100)

	for {
		select {
		case <-t1.C:
			output, err := exec.Command("sh", "-c", fmt.Sprintf("%s get pod -l test=%s | grep -v NAME | awk {'print $1'}", KUBECTL, label)).Output()
			if err != nil {
				println("err:", err)
				println("output:", string(output))
				continue
			}
			podList := strings.Split(string(output), "\n")
			if len(podList)-1 < replica {
				println("string(output))", string(output))
				println("b")
				continue
			}
			expectNum := 0
			for _, pod := range podList {
				if pod != "" {
					output, err := exec.Command(KUBECTL, "get", "pod", pod, "--template={{.status.phase}}").Output()
					if err == nil && string(output) == "Running" {
						println("expectNum: ", expectNum)
						expectNum++
					}
				}
			}
			if expectNum == replica {
				println("1")
			}
		case <-t2.C:
			println("2")
		}
	}
}

func generatePod(label string) string {
	podYaml := fmt.Sprintf("apiVersion: v1 \n"+
		"kind: Pod \n"+
		"metadata: \n"+
		"  name: %s \n"+
		"spec: \n"+
		"  containers: \n"+
		"    - name: explorer \n"+
		"      image: 172.16.1.41:5000/transwarp-centos6:20150311-01 \n"+
		"      args: \n"+
		"        - \"/bin/bash\" \n"+
		"        - \"-c\" \n"+
		"        - \"while true;do sleep 1;done\" \n"+
		"      volumeMounts: \n"+
		"        - mountPath: \"/mount/test-volume\" \n"+
		"          name: test-volume \n"+
		"  volumes: \n"+
		"    - name: test-volume \n"+
		"      tosDisk: \n"+
		"        name: %s \n"+
		"        storageType: %s \n"+
		"        capability: 1Gi \n"+
		"        limits: \n"+
		"          blkio.throttle.read_iops_device: 100 \n"+
		"          blkio.throttle.write_iops_device: 200 \n"+
		"          blkio.throttle.read_bps_device: 100M \n"+
		"          blkio.throttle.write_bps_device: 200M \n",
		pod+label, pod+label, poolname)
	return podYaml
}

func generatePetSetYaml(replica, label string) string {
	petsetYaml := fmt.Sprintf("apiVersion: apps/v1beta1 \n"+
		"kind: StatefulSet \n"+
		"metadata: \n"+
		"  name: %s \n"+
		"spec: \n"+
		"  serviceName: \"nginx\" \n"+
		"  replicas: %s \n"+
		"  template: \n"+
		"    metadata: \n"+
		//"      labels: \n"+
		//"        test: %s \n"+
		"      annotations: \n"+
		"        pod.alpha.kubernetes.io/initialized: \"true\" \n"+
		"    spec: \n"+
		"      terminationGracePeriodSeconds: 0 \n"+
		"      containers: \n"+
		"      - name: nginx \n"+
		"        image: 172.16.1.41:5000/transwarp-centos6:20150311-01 \n"+
		"        args: \n"+
		"          - \"/bin/bash\" \n"+
		"          - \"-c\" \n"+
		"          - \"echo REPLICAID: $REPLICAS; while true;do sleep 1;done\" \n"+
		"        env: \n"+
		"          - name: REPLICAS \n"+
		"            valueFrom: \n"+
		"              fieldRef: \n"+
		"                fieldPath: metadata.annotations.transwarp.replicaid \n"+
		"        volumeMounts: \n"+
		"        - name: www1 \n"+
		"          mountPath: /usr/share/nginx/html1 \n"+
		//"        - name: www2 \n"+
		//"          mountPath: /usr/share/nginx/html2 \n"+
		"  volumeClaimTemplates: \n"+
		"  - metadata: \n"+
		"      name: www1 \n"+
		"      annotations: \n"+
		"        volume.beta.kubernetes.io/storage-class: %s \n"+
		"    spec: \n"+
		"      accessModes: [ \"ReadWriteOnce\" ] \n"+
		"      resources: \n"+
		"        requests: \n"+
		"          storage: 1Gi \n"+
		"        limits: \n"+
		"          blkio.throttle.read_iops_device: 100 \n"+
		"          blkio.throttle.write_iops_device: 200 \n"+
		"          blkio.throttle.read_bps_device: 100M \n"+
		"          blkio.throttle.write_bps_device: 200M \n",
		//"  - metadata: \n"+
		//"      name: www2 \n"+
		//"      annotations: \n"+
		//"        volume.beta.kubernetes.io/storage-class: %s \n"+
		//"    spec: \n"+
		//"      accessModes: [ \"ReadWriteOnce\" ] \n"+
		//"      resources: \n"+
		//"        requests: \n"+
		//"          storage: 1Gi \n"+
		//"        limits: \n"+
		//"          blkio.throttle.read_iops_device: 100 \n"+
		//"          blkio.throttle.write_iops_device: 200 \n"+
		//"          blkio.throttle.read_bps_device: 100M \n"+
		//"          blkio.throttle.write_bps_device: 200M \n",
		petset+label, replica, poolname)
	return petsetYaml
}

func kubeCreate(yamlContent string) error {
	cmd := exec.Command(KUBECTL, "create", "-f", "-")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, yamlContent)
	}()
	err = cmd.Run()
	return err
}

func stop1(stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			fmt.Println("stop1")
			return
		}
	}
}

func stop2(stopCh <-chan struct{}) {
	for {
		select {
		case <-stopCh:
			fmt.Println("stop2")
			return
		}
	}
}

func waitPVCDeleted() {
	t1 := time.NewTicker(time.Second * 1)
	t2 := time.NewTicker(time.Second * 60)

	for {
		select {
		case <-t1.C:
			countSum, err := exec.Command("sh", "-c", "/home/shentao/gopath/src/k8s.io/kubernetes/_output/bin/kubectl"+" get node | grep -v NAME| wc -l").Output()
			if err != nil {
				fmt.Println(err)
			}
			countReady, err := exec.Command("sh", "-c", "/home/shentao/gopath/src/k8s.io/kubernetes/_output/bin/kubectl"+" get node | grep Ready | wc -l").Output()
			if err != nil {
				fmt.Println(err)
			}
			if string(countReady) == "1\n" && string(countSum) == string(countReady) {
				fmt.Println(string(countReady) + "123")
				return
			}
		case <-t2.C:
			fmt.Println("time out")
			return
		}
	}
}


func execCmd(commmand string, args []string) (int, error) {
	cmd := exec.Command(commmand, args...)
	if err := cmd.Start(); err != nil {
		return -1, err
	}
	return cmd.Process.Pid, nil
}
