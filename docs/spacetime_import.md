# Supergiant Spacetime Import Control
## `supergiant create spacetime import <flags>`
*Prerequisite: You must have created a Spacetime Provider with `supergiant create spacetime provider`*

Import control allows the user to import an already existing Kubernetes cluster into spacetime. This cluster will not be managed by Terraform, but you will be able to deploy the Supergiant core, and manipulate your spacetime instance using the CLI.

### Import spacetime.
#### `supergiant create spacetime import<flags>`

Help: `supergiant create spacetime import --help`

Flags:

*required:*

`--ip <IP Address>` - The IP address for the API/Master of your Kubernetes cluster.

`--user <User Name>` - The user name credential for your Kubernetes cluster.

`--pass <Password>` - The password credential for your Kubernetes cluster.

`--name <Cluster Name>` - The name you would like to assign your imported cluster. This can be anything you like. This just needs to be unique.

`--region <Region>` - The Region your Kubernetes cluster is located. Used by the core.

`--az <Availability Zone>` - The Availability Zone where your Kubernetes cluster is located. Used by the core.

`--provider <Provider Name>` - The provider name you would like to use with your Kubernetes cluster.

### Import spacetime (Delete).
You can delete a imported Kubernetes cluster with `supergiant delete spacetime` Your imported cluster will simply get removed from the local spacetime database because it is "Not Managed" by spacetime.
