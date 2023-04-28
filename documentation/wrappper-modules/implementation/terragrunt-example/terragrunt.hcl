
locals {

  # Inputs that can be shared across all the child modules
  inputs = yamldecode(file("inputs.yaml"))

  # Loads the accounts.json
  common_vars = read_terragrunt_config("${get_parent_terragrunt_dir()}/common.hcl")

  # Loads the account related details like account name, id etc.
  account_vars = read_terragrunt_config(find_in_parent_folders("account.hcl"))

  # Loads the aws region information
  region_vars = read_terragrunt_config(find_in_parent_folders("region.hcl"))

  account_name = local.account_vars.locals.account_name
  region       = local.region_vars.locals.aws_region
  accounts     = local.common_vars.locals.accounts

  # AWS Profile name
  profile_name = local.accounts[local.account_name]

  relative_path      = path_relative_to_include()
  environment_number = basename(local.relative_path)
}

# Generate the AWS provider settings
generate "provider" {

  path      = "provider.tf"
  if_exists = "overwrite"
  contents  = <<EOF
provider "aws" {
  profile = "${local.profile_name}"
  region  = "${local.region}"
}
EOF

}

# Generates the config file for s3 backend
remote_state {
  backend = "s3"
  generate = {
    path      = "backend.tf"
    if_exists = "overwrite"
  }
  config = {
    profile        = "${local.profile_name}"
    bucket         = "demo-cache-${local.region}-${local.account_name}-${local.environment_number}-tfstate-000"
    key            = "terraform.tfstate"
    region         = "${local.region}"
    encrypt        = true
    dynamodb_table = "demo-cache-${local.region}-${local.account_name}-${local.environment_number}-tflocks-000"
  }
}

inputs = {
  naming_prefix = local.inputs.naming_prefix
  environment   = local.account_name
  profile_name  = local.profile_name
  region        = local.region
}
