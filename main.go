package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"net/http"
	"path/filepath"
)

type KubeContext struct {
	clientset *kubernetes.Clientset
}

func NewHandlerContext(clientset *kubernetes.Clientset) *KubeContext {
	if clientset == nil {
		panic("nil Kubernetes client!")
	}
	return &KubeContext{clientset}
}

func (ctx *KubeContext) healthzHandler(w http.ResponseWriter, r *http.Request) {
	pods, err := ctx.clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "healthz: OK There are %d pods in the cluster\n", len(pods.Items))
}

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	kubectx := NewHandlerContext(clientset)

	mux := http.NewServeMux()

	// XXX: Not there yet
	mux.Handle("/metrics", http.NotFoundHandler())

	mux.HandleFunc("/healthz", kubectx.healthzHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
