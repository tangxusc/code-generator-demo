package main

import (
	"code-generator-test/generated/clientset/versioned"
	"code-generator-test/generated/clientset/versioned/typed/samplecontroller/v1alpha1"
	"code-generator-test/generated/informers/externalversions"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	config, e := clientcmd.BuildConfigFromFlags("10.30.21.238:6443", "/home/tangxu/.kube/config")
	if e != nil {
		panic(e.Error())
	}
	client, e := v1alpha1.NewForConfig(config)
	if e != nil {
		panic(e.Error())
	}
	fooList, e := client.Foos("test").List(metav1.ListOptions{})
	fmt.Println(fooList, e)

	clientset, e := versioned.NewForConfig(config)
	factory := externalversions.NewSharedInformerFactory(clientset, 30*time.Second)
	foo, e := factory.Samplecontroller().V1alpha1().Foos().Lister().Foos("test").Get("test")
	if e != nil {
		panic(e.Error())
	}
	fmt.Println(foo, e)
}
