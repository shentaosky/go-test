package main

import (
	"encoding/json"
	"fmt"
    "io"
    "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "time"

    apps "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "os/exec"
)

func main() {

	cfg, err := clientcmd.BuildConfigFromFlags("", "/var/lib/kube-controller/kubeconfig")
	if err != nil {
		fmt.Println(err)
		return
	}
	client := kubernetes.NewForConfigOrDie(cfg)
	deployment1 := "systemdhangtest1"
	deployment2 := "systemdhangtest2"
	d1, err := client.AppsV1().Deployments("default").Create(generateDeployment(deployment1, "tess-node-klccb.stratus.lvs.ebay.com", "hub.tess.io/tashen/busybox:test"))
	//for i:= 100; i < 200; i++  {
	data, _ := json.MarshalIndent(d1, "", "    ")
	fmt.Println(string(data))
	if err != nil && !apierrors.IsNotFound(err){
		fmt.Println(err)
		return
	}
	d2, err := client.AppsV1().Deployments("default").Create(generateDeployment(deployment2, "tess-node-lr88h.stratus.lvs.ebay.com", "hub.tess.io/tashen/busybox:test"))
	if err != nil && !apierrors.IsNotFound(err) {
		fmt.Println(err)
		return
	}
	//    podName := name + strconv.Itoa(i)
	//    podYaml := generateTestPod(podName)
	//    if i == 100 {
	//        fmt.Println(podYaml)
	//    }
	//
	//    if err := kubeCreateYaml(podYaml); err != nil {
	//        fmt.Println(err)
	//        continue
	//    }
	//    fmt.Printf("%s done \n", podName)
	//}
	for true {
		scale1, err := client.AppsV1().Deployments(d1.Namespace).GetScale(d1.Name, metav1.GetOptions{})
		if err != nil {
			fmt.Printf("%v: %v", scale1, err)
			time.Sleep(time.Second)
			continue
		}
		scale1.Spec.Replicas = 90
		_, err = client.AppsV1().Deployments(d1.Namespace).UpdateScale(d1.Name, scale1)
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)
			continue
		}
		scale2, err := client.AppsV1().Deployments(d1.Namespace).GetScale(d2.Name, metav1.GetOptions{})
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)
			continue
		}

		scale2.Spec.Replicas = 90
		_, err = client.AppsV1().Deployments(d2.Namespace).UpdateScale(d2.Name, scale2)
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)
			continue
		}
		if !waitDeploymentReady(client, d1, d2) {
			time.Sleep(time.Hour * 2)
		}
		scale1, err = client.AppsV1().Deployments(d1.Namespace).GetScale(d1.Name, metav1.GetOptions{})
		if err != nil {
			fmt.Printf("%v: %v", scale1, err)
			time.Sleep(time.Second)
			continue
		}
		scale1.Spec.Replicas = 0
		scale2, err = client.AppsV1().Deployments(d1.Namespace).GetScale(d2.Name, metav1.GetOptions{})
		if err != nil {
			fmt.Println(err)
			break
		}
		scale2.Spec.Replicas = 0
		_, err = client.AppsV1().Deployments(d1.Namespace).UpdateScale(d1.Name, scale1)
		if err != nil {
			fmt.Println(err)
			break
		}
		_, err = client.AppsV1().Deployments(d2.Namespace).UpdateScale(d2.Name, scale2)
		if err != nil {
			fmt.Println(err)
			break
		}

        waitDeploymentDown(client, d1, d2)

	}
}

func waitDeploymentReady(client *kubernetes.Clientset, deployment1, deployment2 *apps.Deployment) bool {
	t2 := time.NewTicker(time.Minute * 3)
	t3 := time.NewTicker(time.Minute * 6)
	t4 := time.NewTicker(time.Minute * 10)
	for {
		select {

		case <-t2.C:
			fmt.Printf("exceed 3 min to ready for deployment: %v \n", time.Now())
		case <-t3.C:
			fmt.Printf("exceed 6 min to ready for deployment: %v \n", time.Now())
		case <-t4.C:
            d1, _ := client.AppsV1().Deployments(deployment1.Namespace).Get(deployment1.Name, metav1.GetOptions{})
            d2, _ := client.AppsV1().Deployments(deployment2.Namespace).Get(deployment2.Name, metav1.GetOptions{})
			fmt.Printf("exceed 10 min to ready for deployment: %s . Exist loop \n", time.Now())
            fmt.Println(d1.Status.AvailableReplicas, d2.Status.AvailableReplicas)
			return false
		default:
			time.Sleep(time.Second* 20)
			d1, err := client.AppsV1().Deployments(deployment1.Namespace).Get(deployment1.Name, metav1.GetOptions{})
			if err != nil {
				fmt.Println(err)
				continue
			}
			d2, err := client.AppsV1().Deployments(deployment2.Namespace).Get(deployment2.Name, metav1.GetOptions{})
			if err != nil {
				fmt.Println(err)
				continue
			}
			if d1.Status.AvailableReplicas == d1.Status.Replicas && d2.Status.AvailableReplicas == d2.Status.Replicas {
				return true
			}
		}
	}
}

func waitDeploymentDown(client *kubernetes.Clientset, deployment1, deployment2 *apps.Deployment) bool {
	for {
		time.Sleep(time.Second * 20)
		d1, err := client.AppsV1().Deployments(deployment1.Namespace).Get(deployment1.Name, metav1.GetOptions{})
		if err != nil {
			fmt.Println(err)
			continue
		}
		d2, err := client.AppsV1().Deployments(deployment2.Namespace).Get(deployment2.Name, metav1.GetOptions{})
		if err != nil {
			fmt.Println(err)
			continue
		}
		if d1.Status.Replicas == 0 && d2.Status.Replicas == 0 {
			return true
		}
	}
}

// generateDeployment creates a deployment, with the input image as its template
func generateDeployment(name, hostname, image string) *apps.Deployment {
	podLabels := map[string]string{"systemd_test": name}
	return &apps.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Annotations: make(map[string]string),
			Labels:      podLabels,
		},
		Spec: apps.DeploymentSpec{
			Replicas: func(i int32) *int32 { return &i }(0),
			Selector: &metav1.LabelSelector{MatchLabels: podLabels},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: podLabels,
				},
				Spec: v1.PodSpec{
					Tolerations: []v1.Toleration{
						{
							Key:      "test",
							Operator: v1.TolerationOpEqual,
							Value:    "systemdhang",
							Effect:   v1.TaintEffectNoSchedule,
						},
					},
					Containers: []v1.Container{
						{
							Name:                   "test",
							Resources: v1.ResourceRequirements{
									Limits: v1.ResourceList{
										v1.ResourceName(v1.ResourceCPU): resource.MustParse("5"),
										v1.ResourceName(v1.ResourceMemory): resource.MustParse("5G"),
									},
									Requests: v1.ResourceList{
										v1.ResourceName(v1.ResourceCPU): resource.MustParse("50m"),
										v1.ResourceName(v1.ResourceMemory): resource.MustParse("50m"),
									},
							},
							Args: []string{"/bin/sh", "-c", "sleep 3600"},
							Image:                  image,
							ImagePullPolicy:        v1.PullAlways,
							TerminationMessagePath: v1.TerminationMessagePathDefault,
							VolumeMounts: []v1.VolumeMount{
								{
									MountPath: "/var/nfs/test1",
									Name:      "nfs1",
								},
								{
									MountPath: "/var/nfs/test2",
									Name:      "nfs2",
								},
								{
									MountPath: "/var/nfs/test3",
									Name:      "nfs3",
								},
								{
									MountPath: "/var/nfs/test4",
									Name:      "nfs4",
								},
								{
									MountPath: "/var/nfs/test5",
									Name:      "nfs5",
								},
								{
									MountPath: "/var/nfs/test6",
									Name:      "nfs6",
								},
								{
									MountPath: "/var/nfs/test7",
									Name:      "nfs7",
								},
								{
									MountPath: "/var/nfs/test8",
									Name:      "nfs8",
								},
								{
									MountPath: "/var/nfs/test9",
									Name:      "nfs9",
								},
								{
									MountPath: "/var/nfs/test10",
									Name:      "nfs10",
								},
								{
									MountPath: "/var/nfs/test11",
									Name:      "nfs11",
								},
								{
									MountPath: "/var/nfs/test12",
									Name:      "nfs12",
								},
								{
									MountPath: "/var/nfs/test13",
									Name:      "nfs13",
								},
								{
									MountPath: "/var/nfs/test14",
									Name:      "nfs14",
								},
								{
									MountPath: "/var/nfs/test15",
									Name:      "nfs15",
								},
							},
						},
					},
					HostNetwork:                   true,
					NodeName:                      hostname,
					RestartPolicy:                 v1.RestartPolicyAlways,
					Volumes: []v1.Volume{
						{
							Name: "nfs1",
							VolumeSource: v1.VolumeSource{
								NFS: &v1.NFSVolumeSource{
									Server: "10.196.147.93",
									Path:   "/var/nfsshare",
								},
							},
						},
						{
							Name: "nfs2",
							VolumeSource: v1.VolumeSource{
								NFS: &v1.NFSVolumeSource{
									Server: "10.196.147.93",
									Path:   "/var/nfsshare",
								},
							},
						},
						{
							Name: "nfs3",
							VolumeSource: v1.VolumeSource{
								NFS: &v1.NFSVolumeSource{
									Server: "10.196.147.93",
									Path:   "/var/nfsshare",
								},
							},
						},
						{
							Name: "nfs4",
							VolumeSource: v1.VolumeSource{
								NFS: &v1.NFSVolumeSource{
									Server: "10.196.147.93",
									Path:   "/var/nfsshare",
								},
							},
						},
						{
							Name: "nfs5",
							VolumeSource: v1.VolumeSource{
								NFS: &v1.NFSVolumeSource{
									Server: "10.196.147.93",
									Path:   "/var/nfsshare",
								},
							},
						},
						{
							Name: "nfs6",
							VolumeSource: v1.VolumeSource{
								NFS: &v1.NFSVolumeSource{
									Server: "10.196.147.93",
									Path:   "/var/nfsshare",
								},
							},
						}, {
							Name: "nfs7",
							VolumeSource: v1.VolumeSource{
								NFS: &v1.NFSVolumeSource{
									Server: "10.196.147.93",
									Path:   "/var/nfsshare",
								},
							},
						},
						{
							Name: "nfs8",
							VolumeSource: v1.VolumeSource{
								NFS: &v1.NFSVolumeSource{
									Server: "10.196.147.93",
									Path:   "/var/nfsshare",
								},
							},
						},
						{
							Name: "nfs9",
							VolumeSource: v1.VolumeSource{
								NFS: &v1.NFSVolumeSource{
									Server: "10.196.147.93",
									Path:   "/var/nfsshare",
								},
							},
						},
						{
							Name: "nfs10",
							VolumeSource: v1.VolumeSource{
								NFS: &v1.NFSVolumeSource{
									Server: "10.196.147.93",
									Path:   "/var/nfsshare",
								},
							},
						},
						{
							Name: "nfs11",
							VolumeSource: v1.VolumeSource{
								NFS: &v1.NFSVolumeSource{
									Server: "10.196.147.93",
									Path:   "/var/nfsshare",
								},
							},
						},
						{
							Name: "nfs12",
							VolumeSource: v1.VolumeSource{
								NFS: &v1.NFSVolumeSource{
									Server: "10.196.147.93",
									Path:   "/var/nfsshare",
								},
							},
						},
						{
							Name: "nfs13",
							VolumeSource: v1.VolumeSource{
								NFS: &v1.NFSVolumeSource{
									Server: "10.196.147.93",
									Path:   "/var/nfsshare",
								},
							},
						},
						{
							Name: "nfs14",
							VolumeSource: v1.VolumeSource{
								NFS: &v1.NFSVolumeSource{
									Server: "10.196.147.93",
									Path:   "/var/nfsshare",
								},
							},
						},
						{
							Name: "nfs15",
							VolumeSource: v1.VolumeSource{
								NFS: &v1.NFSVolumeSource{
									Server: "10.196.147.93",
									Path:   "/var/nfsshare",
								},
							},
						},
					},
				},
			},
		},
	}
}

func generateTestPod(name string) string {
	podYaml := fmt.Sprintf("apiVersion: v1 \n"+
		"kind: Pod \n"+
		"metadata: \n"+
		"  name: %s \n"+
		"  labels: \n"+
		"    systemd: test \n"+
		"    systemd2: test \n"+
		"spec: \n"+
		"  nodeName: tess-node-klccb.stratus.lvs.ebay.com \n"+
		"  hostNetwork: true \n"+
		"  containers: \n"+
		"    - name: explorer \n"+
		"      image: hub.tess.io/tashen/busybox:test \n"+
		"      resources: \n"+
		"        requests: \n"+
		"          cpu: 50m \n"+
		"          memory: 50m \n"+
		"        limits: \n"+
		"          cpu: 1 \n"+
		"          memory: 1G \n"+
		"      args: \n"+
		"        - \"/bin/sh\" \n"+
		"        - \"-c\" \n"+
		"        - \"sleep 3600\" \n"+
		"      volumeMounts: \n"+
		"        - mountPath: \"/mount/test-volume\" \n"+
		"          name: test-volume \n"+
		"        - name: nfs-volume1 \n"+
		"          mountPath: /var/nfs/test1 \n"+
		"        - name: nfs-volume2 \n"+
		"          mountPath: /var/nfs/test2 \n"+
		"        - name: nfs-volume3 \n"+
		"          mountPath: /var/nfs/test3 \n"+
		"        - name: nfs-volume4 \n"+
		"          mountPath: /var/nfs/test4 \n"+
		"        - name: nfs-volume5 \n"+
		"          mountPath: /var/nfs/test5 \n"+
		"        - name: nfs-volume6 \n"+
		"          mountPath: /var/nfs/test6 \n"+
		"        - name: nfs-volume7 \n"+
		"          mountPath: /var/nfs/test7 \n"+
		"        - name: nfs-volume8 \n"+
		"          mountPath: /var/nfs/test8 \n"+
		"        - name: nfs-volume9 \n"+
		"          mountPath: /var/nfs/test9 \n"+
		"        - name: nfs-volume10 \n"+
		"          mountPath: /var/nfs/test10 \n"+
		"  volumes: \n"+
		"    - name: test-volume \n"+
		"      emptyDir: {} \n"+
		"    - name: nfs-volume1 \n"+
		"      nfs:\n"+
		"        server: 10.196.147.93\n"+
		"        path: /var/nfsshare \n"+
		"    - name: nfs-volume2 \n"+
		"      nfs:\n"+
		"        server: 10.196.147.93\n"+
		"        path: /var/nfsshare \n"+
		"    - name: nfs-volume3 \n"+
		"      nfs:\n"+
		"        server: 10.196.147.93\n"+
		"        path: /var/nfsshare \n"+
		"    - name: nfs-volume4 \n"+
		"      nfs:\n"+
		"        server: 10.196.147.93\n"+
		"        path: /var/nfsshare \n"+
		"    - name: nfs-volume5 \n"+
		"      nfs:\n"+
		"        server: 10.196.147.93\n"+
		"        path: /var/nfsshare \n"+
		"    - name: nfs-volume6 \n"+
		"      nfs:\n"+
		"        server: 10.196.147.93\n"+
		"        path: /var/nfsshare \n"+
		"    - name: nfs-volume7 \n"+
		"      nfs:\n"+
		"        server: 10.196.147.93\n"+
		"        path: /var/nfsshare \n"+
		"    - name: nfs-volume8 \n"+
		"      nfs:\n"+
		"        server: 10.196.147.93\n"+
		"        path: /var/nfsshare \n"+
		"    - name: nfs-volume9 \n"+
		"      nfs:\n"+
		"        server: 10.196.147.93\n"+
		"        path: /var/nfsshare \n"+
		"    - name: nfs-volume10 \n"+
		"      nfs:\n"+
		"        server: 10.196.147.93\n"+
		"        path: /var/nfsshare \n",
		name)
	return podYaml
}

func kubeCreateYaml(yamlContent string) error {
	cmd := exec.Command("kubectl", "--kubeconfig=/var/lib/kube-controller/kubeconfig", "create", "-f", "-")
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
