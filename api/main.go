package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"eternauta/pkg"
	"eternauta/pkg/k8s"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func createDeployment(name string) {

}

func main() {

	// generateUID generates a random UID string of 16 bytes in hexadecimal format.

	uid := pkg.GenerateUID()
	fmt.Printf("Generated UID: %s\n", uid)

	name := fmt.Sprintf("test-%s", uid)
	port := int32(8080)
	image := "kicbase/echo-server:1.0"

	// Create a Kubernetes client
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
	if err != nil {
		log.Fatalf("Failed to create k8s config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create k8s clientset: %v", err)
	}

	// Define the deployment
	deployment := k8s.CreateDeploymentManifest(name, image, port)
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	resultDeployment, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Failed to create deployment: %v", err)
	}
	log.Printf("Created deployment %q.\n", resultDeployment.GetObjectMeta().GetName())

	// Define the service
	service := k8s.CreateServiceManifest(name, port)
	serviceClient := clientset.CoreV1().Services(apiv1.NamespaceDefault)
	resultService, err := serviceClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Failed to create Service: %v", err)
	}
	log.Printf("Created Service %q.\n", resultService.GetObjectMeta().GetName())

}
