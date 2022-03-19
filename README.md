# terraform-provider-stackhawk


## Build provider

Run the following command to build the provider

```shell
$ go build -o terraform-provider-stackhawk
```

## Test sample configuration

First, build and install the provider.

```shell
$ make install
```

Then, navigate to the `examples` directory.

```shell
$ cd examples
```

Set `organization_id`

The, run the following command to initialize the workspace and apply the sample configuration.

```shell
$ terraform init && terraform apply
```
