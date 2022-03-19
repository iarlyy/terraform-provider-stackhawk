terraform {
  required_providers {
    stackhawk = {
      source = "iarlyy/stackhawk"
    }
  }
}

provider "stackhawk" {
  organization_id = "abcd1234-01ab-ab01-aabb-aabb9876abc0"
}

resource "stackhawk_application" "foobar" {
  name       = "foobar"
  data_type  = "NONE"
  env        = "development"
  risk_level = "HIGH"
}


data "stackhawk_application" "foobar" {
  application_id = resource.stackhawk_application.foobar.id
}

output "id" {
  value = data.stackhawk_application.foobar.application_id
}

output "name" {
  value = data.stackhawk_application.foobar.name
}
