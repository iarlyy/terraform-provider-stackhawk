---
subcategory: "Application"
layout: "stackhawk"
page_title: "Stackhawk: stackhawk_application"
description: |-
  Get information on an Application
---

# Data Source: stackhawk_application

Use this data source to get the name and other properties of an application.

## Example Usage

```terraform
data "stackhawk_application" "test" {
  application_id = "abcd0123-01ab-ab01-aabb-aabb9876wyz0"
}
```

## Argument Reference

* `application_id` - (Required) The id of the application.

## Attributes Reference

* `name` - The name of the application
