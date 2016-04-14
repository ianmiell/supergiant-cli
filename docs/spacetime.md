# Supergiant Spacetime Control
## `supergiant <verb> spacetime`
*Prerequisite: You must have created a Spacetime Provider with `supergiant create spacetime provider`*

*Configuration Directory: ~/.supergiant*

Spacetime actions control Kubernetes clusters on which your supergiant core will run. The spacetime controller on it's own is a useful way to manage multiple Kubernetes clusters across multiple environments.

### Create spacetime.
#### `supergiant create spacetime <flags>`

Help: `supergiant create spacetime help`

Flags: If any required flags are missed, the spacetime action will prompt the user for missing info.

*Required:*

`--name <name>` - Name you wish to assign to your Kubernetes cluster.

`--user <user name>` - Username you would like to assign to your Kubernetes cluster api.

`--pass <password>` - Password you would like to assign to your kubernetes cluster api.

*Flags with default options.*

`--version <Kubernetes version>` - The version of Kubernetes you would like to use with your cluster. Defaults to the latest supported version.

`--provider <Provider Name>` - Name of the provider you would like to use to launch your Kubernetes cluster. This provider name must exist in the spacetime provider database or the Kubernetes launch will fail. Defaults to a provider named "aws".

`--region <AWS Region>` - The AWS Region that spacetime will use when launching your Kubernetes cluster. This defaults to "us-east-1".

`--avail-zone <AWS Availability Zone>` - The AWS Availability zone that spacetime will use when launching your Kubernetes cluster. This defaults to "us-east-1b".

`--retry <Cluster Name>` - This option retries a filed Kubernetes install. This option does not require other flags. You only need to specify the name of the cluster you wish to rebuild as the first argument for the flag.

Usage Example:

Creating Cluster: This will produce a large amount of output from Terraform.
```
##: $ supergiant create spacetime --name my_first_cluster --provider my_provider_name
Please enter Kube UserName: myuser
Please enter Kube Password: mypassword
template_file.master_user_data: Creating...
  rendered:               "" => "<computed>"
  (Terraform output)....  
```
Cluster Verification: Once spacetime can see that the cluster has been launched, it will start polling the Kubernetes API to verify that everything is working correctly. The verification usually takes around 5 minutes to complete. The verification utility will timeout after ~15 minutes, marking your cluster as failed and assuming verification did not pass.
```
............
aws_autoscaling_group.kubernetes-minions-large: Creation complete

Apply complete! Resources: 18 added, 0 changed, 0 destroyed.

The state of your infrastructure has been saved to the path
below. This state is required to modify and destroy your
infrastructure, so keep it safe. To inspect the complete state
use the `terraform show` command.

State path: terraform.tfstate
Terraform build completed successfully...
Loading cluster informations from terraform...
Checking that Kubernetes is up and running properly...
Check for life poll... 0
Check for life poll... 1
Check for life poll... 2
Check for life poll... 3
Check for life poll... 4
............
```
Cluster Launch Success: Once your cluster is verified spacetime will notify you that all is well.
```
............
Check for life poll... 27
Check for life poll... 28
Check for life poll... 29
Check for life poll... 30
Kubernetes verified...
Success...
##: $
```

### List spacetime.
#### `supergiant list spacetime`

Help: `supergiant list spacetime help`

Flags: *This action has no flags.*

Usage Example:

```
##: $ supergiant list spacetime
Name		        Provider 	AZ		IP		Status	Core Status	Managed	Selected
my_first_cluster    aws 		us-east-1b	1.2.3.4	Active	false		true	true
##: $
```

### Delete spacetime.
#### `supergiant delete spacetime <Cluster Name>`

Help: `supergiant delete spacetime help`

Flags: *This action has no flags.*

Usage Example:

Delete:
```
##: $ supergiant delete spacetime my_first_cluster
template_file.master_user_data: Refreshing state... (ID: xxxxxxxx)
aws_vpc.kube_cluster: Refreshing state... (ID: xxxxxx)
aws_internet_gateway.default: Refreshing state... (ID: xxxxxxx)
aws_subnet.public: Refreshing state... (ID: xxxxxxx)
aws_security_group.kubernetes_sg: Refreshing state... (ID: xxxxxxx)
(Terraform output)....  
```
Success:
```
...........
template_file.master_user_data: Destruction complete
aws_security_group.kubernetes_sg: Destruction complete
aws_subnet.public: Destruction complete
aws_vpc.kube_cluster: Destroying...
aws_vpc.kube_cluster: Destruction complete

Apply complete! Resources: 0 added, 0 changed, 18 destroyed.
Success...
##: $
```
