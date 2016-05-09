package sgcore

import (
	"fmt"

	"github.com/supergiant/guber"
	"github.com/supergiant/supergiant-cli/spacetime"
)

func initSGAPI(c guber.Client, k *spacetime.Kube, version string) error {
	// The default core version.
	if version == "" {
		version = "latest"
	}

	fmt.Println("Installing Supergiant API version:", version)
	provider, err := spacetime.GetProvider(k.Provider)
	if err != nil {
		return err
	}

	service := &guber.Service{
		Metadata: &guber.Metadata{
			Name: "supergiant-api",
			Labels: map[string]string{
				"deployment": "Production",
				"instance":   "supergiant-api",
			},
		},
		Spec: &guber.ServiceSpec{
			Type: "ClusterIP",
			Selector: map[string]string{
				"deployment": "Production",
				"instance":   "supergiant-api",
			},
			Ports: []*guber.ServicePort{
				&guber.ServicePort{
					Name:       "frontend",
					Port:       80,
					TargetPort: 8080,
					Protocol:   "TCP",
				},
			},
		},
	}

	_, err = c.Services("supergiant").Create(service)
	if err != nil {
		return err
	}

	rc := &guber.ReplicationController{
		Metadata: &guber.Metadata{
			Name: "supergiant-api",
		},
		Spec: &guber.ReplicationControllerSpec{
			Selector: map[string]string{
				"instance": "supergiant-api",
			},
			Replicas: 1,
			Template: &guber.PodTemplate{
				Metadata: &guber.Metadata{
					Name: "etcd", // pod base name is same as RC
					Labels: map[string]string{
						"deployment": "Production",     // for Service
						"instance":   "supergiant-api", // for RC
					},
				},
				Spec: &guber.PodSpec{
					Containers: []*guber.Container{
						&guber.Container{
							Name:            "supergiant-api",
							ImagePullPolicy: "Always",
							Ports: []*guber.ContainerPort{
								&guber.ContainerPort{
									Name:          "frontend",
									ContainerPort: 8080,
									Protocol:      "TCP",
								},
							},
							Image: "supergiant/supergiant-api:" + version + "",
							Command: []string{
								"/supergiant-api",
								"--etcd-hosts",
								"http://etcd:2379",
								"--k8s-host",
								k.IP,
								"--k8s-user",
								k.User,
								"--k8s-pass",
								k.Pass,
								"--aws-sg-id",
								k.SGID,
								"--k8s-insecure-https",
								"--aws-access-key",
								provider.AccessKey,
								"--aws-secret-key",
								provider.SecretKey,
								"--enable-capacity-service",
							},
						},
					},
				},
			},
		},
	}

	_, err = c.ReplicationControllers("supergiant").Create(rc)
	if err != nil {
		return err
	}
	return nil

}
