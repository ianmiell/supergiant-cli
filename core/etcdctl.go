package sgcore

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/supergiant/guber"
	"github.com/supergiant/supergiant-cli/spacetime"
)

func initETCD(c guber.Client, p *spacetime.ProviderConfig, k *spacetime.Kube) error {

	// make an EBS folume for our data dir.
	token := ""

	creds := credentials.NewStaticCredentials(p.AccessKey, p.SecretKey, token)
	_, err := creds.Get()
	if err != nil {
		fmt.Println("ERROR: AWS Credentials Install Failed...", err)
	}
	fmt.Println("INFO: AWS Credentials Installed From Provider DB.")

	awsConf := aws.NewConfig().WithRegion(k.Region).WithCredentials(creds)

	EC2 := ec2.New(session.New(), awsConf)

	vtype := "gp2"
	vsize := int64(20)
	volInput := &ec2.CreateVolumeInput{
		AvailabilityZone: aws.String(k.AZ),
		VolumeType:       &vtype,
		Size:             &vsize,
	}

	fmt.Println("Creating ETCD Data Volume.")
	volume, err := EC2.CreateVolume(volInput)
	if err != nil {
		return err
	}

	service := &guber.Service{
		Metadata: &guber.Metadata{
			Name: "etcd",
		},
		Spec: &guber.ServiceSpec{
			Type: "ClusterIP",
			Selector: map[string]string{
				"instance": "etcd",
			},
			Ports: []*guber.ServicePort{
				&guber.ServicePort{
					Name:       "client",
					Port:       2379,
					Protocol:   "TCP",
					TargetPort: 2379,
				},
				&guber.ServicePort{
					Name:       "server",
					Port:       2380,
					Protocol:   "TCP",
					TargetPort: 2380,
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
			Name: "etcd",
		},
		Spec: &guber.ReplicationControllerSpec{
			Selector: map[string]string{
				"instance": "etcd",
			},
			Replicas: 1,
			Template: &guber.PodTemplate{
				Metadata: &guber.Metadata{
					Name: "etcd", // pod base name is same as RC
					Labels: map[string]string{
						"deployment": "Production", // for Service
						"instance":   "etcd",       // for RC
					},
				},
				Spec: &guber.PodSpec{
					Volumes: []*guber.Volume{
						&guber.Volume{
							Name: "data",
							AwsElasticBlockStore: &guber.AwsElasticBlockStore{
								VolumeID: *volume.VolumeId,
								FSType:   "ext4",
							},
						},
					},
					Containers: []*guber.Container{
						&guber.Container{
							Name:  "etcd",
							Image: "quay.io/coreos/etcd:latest",
							VolumeMounts: []*guber.VolumeMount{
								&guber.VolumeMount{
									Name:      "data",
									MountPath: "/data",
								},
							},
							Command: []string{
								"/etcd",
								"--name",
								"etcd",
								"--data-dir=/data",
								"--initial-advertise-peer-urls",
								"http://etcd:2380",
								"--listen-peer-urls",
								"http://0.0.0.0:2380",
								"--listen-client-urls",
								"http://0.0.0.0:2379",
								"--advertise-client-urls",
								"http://etcd:2379",
								"--initial-cluster",
								"etcd=http://etcd:2380",
								"--initial-cluster-state",
								"new",
							},
							Ports: []*guber.ContainerPort{
								&guber.ContainerPort{
									Name:          "client",
									ContainerPort: 2379,
									Protocol:      "TCP",
								},
								&guber.ContainerPort{
									Name:          "server",
									ContainerPort: 2380,
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
