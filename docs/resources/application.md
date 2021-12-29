---
subcategory: "Application"
layout: "stackhawk"
page_title: "Stackhawk: stackhawk_application"
description: |-
  Provides an application resource.
---

# Resource: stackhawk_application

## Example Usage

```terraform
resource "stackhawk_application" "test" {
  name       = "test"
  data_type  = "NONE"
  env        = "stg"
  risk_level = "HIGH"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the application.
* `data_type` - (Required) The type of data of this application. Supported values are: `NONE, PII, PCI, FIN, PKI, HIPAA, FERPA`
* `env` - (Required) The name of the initial environment for this new application.
* `risk_level` - (Required) The risk leve of the application. Supported values are: `LOW, MEDIUM, HIGH, CRITICAL`

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` or `application_id` - The id of the application
* `env_id` - The id of the application environment

## Import

Applications can be imported using `application id`

```
$ terraform import stackhawk_application.test abcd0123-01ab-ab01-aabb-aabb9876wyz0

```
