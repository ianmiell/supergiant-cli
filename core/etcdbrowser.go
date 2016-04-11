package sgcore

import "github.com/supergiant/guber"

//This may be stupid, but for now it will aid with dubugs.

func initETCDBrowser(c *guber.Client) error {
	service := &guber.Service{
		Metadata: &guber.Metadata{
			Name: "sg-etcd-browser",
			Labels: map[string]string{
				"deployment": "Production",
				"instance":   "sg-etcd-browser",
			},
		},
		Spec: &guber.ServiceSpec{
			Type: "ClusterIP",
			Selector: map[string]string{
				"deployment": "Production",
				"instance":   "sg-etcd-browser",
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

	_, err := c.Services("supergiant").Create(service)
	if err != nil {
		return err
	}

	rc := &guber.ReplicationController{
		Metadata: &guber.Metadata{
			Name: "sg-etcd-browser",
			Labels: map[string]string{
				"deployment": "Production",
				"instance":   "sg-etcd-browser",
			},
		},
		Spec: &guber.ReplicationControllerSpec{
			Selector: map[string]string{},
			Replicas: 1,
			Template: &guber.PodTemplate{
				Metadata: &guber.Metadata{
					Name: "etcd",
					Labels: map[string]string{
						"deployment": "Production",
						"instance":   "sg-etcd-browser",
					},
				},
				Spec: &guber.PodSpec{
					Containers: []*guber.Container{
						&guber.Container{
							Name:  "sg-etcd-browser",
							Image: "nikfoundas/etcd-viewer",
							Ports: []*guber.ContainerPort{
								&guber.ContainerPort{
									Name:          "frontend",
									ContainerPort: 8080,
									Protocol:      "TCP",
								},
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
