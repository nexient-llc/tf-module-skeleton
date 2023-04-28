locals {
  default_region = "us-east-2"

  accounts = jsondecode(file("accounts.json"))
}
