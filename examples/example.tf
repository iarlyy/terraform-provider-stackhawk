terraform {
  required_providers {
    stackhawk = {
      source = "iarlyy/stackhawk"
    }
  }
}

provider "stackhawk" {
  organization_id = "abcd1234-ab01-ab01-ab01-abcdef123456"
}

resource "stackhawk_application" "my_test_app" {
  name       = "app-test"
  data_type  = "NONE"
  env        = "stg"
  risk_level = "HIGH"
}


data "stackhawk_application" "my_test_app" {
  application_id = resource.stackhawk_application.my_test_app.id
}

output "id" {
  value = data.stackhawk_application.my_test_app.application_id
}

output "name" {
  value = data.stackhawk_application.my_test_app.name
}
