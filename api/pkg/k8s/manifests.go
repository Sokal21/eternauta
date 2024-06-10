package k8s

import (
	"fmt"

	"eternauta/pkg"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateDeploymentManifest(name string, image string, port int32) *appsv1.Deployment {
	// Define the deployment
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s-deployment", name),
			Labels: map[string]string{
				"app": name,
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: pkg.Int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  name,
							Image: image,
							Ports: []apiv1.ContainerPort{
								{
									ContainerPort: port,
								},
							},
						},
					},
				},
			},
		},
	}
}

func CreateServiceManifest(name string, port int32) *apiv1.Service {
	return &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s-service", name),
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{"app": name},
			Ports:    []apiv1.ServicePort{{Protocol: "TCP", Port: port}},
		},
	}
}
