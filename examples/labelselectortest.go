package main

import (
    "fmt"
    "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/labels"
    "tess.io/ebay/vm-volume/pkg/pdcontroller"
)

func main() {
    label := labels.SelectorFromSet(labels.Set(map[string]string{"k8s-app": "kube-dns"}))
    labelSelector := v1.LabelSelector{MatchLabels: map[string]string{pdcontroller.ClusterSelector: "12"}}
    fmt.Println(labelSelector.String(), label)
}
