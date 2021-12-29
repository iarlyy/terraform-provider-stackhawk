---
layout: "stackhawk"
page_title: "Provider: Stackhawk"
description: |-
  Use the Stackhawk provider to interact with the many resources supported by Stackhawk. You must configure the provider with the proper credentials before you can use it.
---

# Stackhawk Provider

## Example Usage

```terraform
terraform {
  required_providers {
    stackhawk = {
      source  = "iarlyy/stackhawk"
      version = "~> 0.1"
    }
  }
}

provider "stackhawk" {
  organization_id = "abcd0123-01ab-ab01-aabb-aabb9876wyz0"
}
```

## Authentication

### Static Credentials

Static credentials can be provided by adding an in-line `api_key` in the
provider block:

Usage:

```terraform
provider "stackhawk" {
  organization_id = "abcd1234-01ab-ab01-aabb-aabb9876abc0"
  api_token       = "hawk.abc..."
}
```

### Environment Variables

You can provide your api key via the `HAWK_API_KEY` environment variable

```terraform
provider "stackhawk" {
  organization_id = "abcd1234-01ab-ab01-aabb-aabb9876abc0"
}
```

Usage:

```sh
$ export HAWK_API_KEY="hawk.abc...."
$ terraform plan
```

## Argument Reference

The following arguments are supported in the Stackhawk `provider` block:

* `organization_id` - (Required) The UUID identifier of the organization that
will own this application.

* `api_key` - (Optional) This is the Stackhawk api key. It must be provided, but
  it can also be sourced from the `HAWK_API_KEY` environment variable.
