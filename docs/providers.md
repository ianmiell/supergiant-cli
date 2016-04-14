# Supergiant Spacetime Provider Control
## `supergiant <verb> spacetime provider <flags>`

*Configuration Directory: ~/.supergiant*

The spacetime provider control tool allows the user to define cloud credentials to use while performing supergiant actions. These credentials can be used while creating spacetime clusters `supergiant create spacetime`, importing Kubernetes clusters into the spacetime database `supergiant create spacetime import`, and are consumed by the Supergiant core when deployed to your cluster `supergiant create core`. Spacetime maintains a separate provider database in your spacetime config, because often a user may have or want Kubernetes clusters to reside in multiple Regions, Zones, Accounts, or Clouds using unique credentials.

### Create spacetime provider.
#### `supergiant create spacetime provider <flags>`

Help: `supergiant create spacetime provider --help`

Flags:

*Required:*

`--name <Provider Name>`  - The name to associate with this provider. The name should be unique, and is used to reference your provider while performing other spacetime actions.

   *Currently only AWS is supported, more clouds coming soon!*

`--access-key <AWS Access Key>` - The AWS Access Key to assign to the provider.

`--secret-key <AWS Secret Key>` - The AWS Secret Key to assign to the provider.

*Optional*

`--provider-service <service>` - This flag notes the cloud provider the provider refers to. Defaults to "aws".

`--rebuild` - This flag tells spacetime that you would like to re-build or re-verify your credentials. (Requires the `--name` flag.)

Usage Example:

Create Provider:
```
##$ supergiant create spacetime provider
Please enter Spacetime Provider Name: my_provider
Please enter AWS Access Key: my_access_key
Please enter AWS Secret Key: my_secret_key
Verifying your credentials...
(spacetime will validate your credentials are valid.)......
```

### List spacetime provider.
#### `supergiant list spacetime provider`

Help: `supergiant list spacetime provider --help`

Flags: *This action has no flags.*

Usage Example:

```
##$ supergiant list spacetime provider
Provider	Type	Status 	  Access Key	   Secret Key
my_provider	aws	verified 	my_access_key	my_secret_key
##$
```

### Delete spacetime provider.
#### `supergiant delete spacetime provider <Provider Name>`

Help: `supergiant delete spacetime provider --help`

Flags: *This action has no flags.*

Usage Example:
```
##$ supergiant list spacetime provider
Provider	Type	Status 	  Access Key	Secret Key
my_provider	aws	verified 	my_access_key	my_secret_key
##$ supergiant delete spacetime provider my_provider
Success...
##$
```
