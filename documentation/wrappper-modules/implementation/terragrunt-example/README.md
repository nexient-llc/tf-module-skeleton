# memcached-cluster
This sample repository will contains terragrunt configuration files to provision  `memcached-cluster`. It refers to the terraform module https://github.com/nexient-llc/tf-aws-wrapper_module-memcached_cluster.git

# Pre-requisites
1. Configure the profiles in `accounts.json`. The AWS profiles by this names should exist in `~/.aws/config` directory on your laptop
```
{
  "perf": "demo-perf-admin",
  "preqa": "demo-preqa-admin",
  "prod": "demo-prod-admin",
  "qa": "demo-qa-admin",
  "sandbox": "demo-sandbox-admin",
  "uat": "demo-uat-admin"
}
```
2. Make sure `terragrunt` in installed on your machine. This directory supports `asdf`. Check the `.tool-versions` file for more details.

3. Configure `S3 Backend` in root `terragrunt.hcl`. Terragrunt is capable to create s3 buckets and dynamodb tables, if not already present
```
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
```
4. Verify the `account.hcl` in each enironment. The `account_name` should match the entry in `accounts.json`

5. Verify the `region.hcl` for the correct AWS Region.

6. Edit the `inputs.yaml` in each environment. For example: `env/sandbox/us-east-2/000/inputs.yaml`

7. Run terragrunt

```
cd env/sandbox/us-east-2/000
terragrunt init
terragrunt plan
terragrunt apply

# to check output
terragrunt output
```

# Directory Structure
The directory structure of the terragrunt configuration looks like as shown below
```
.
├── common
│   └── memcached_cluster.hcl
├── env
│   ├── perf
│   │   ├── us-east-2
│   │   │   ├── 000
│   │   │   │   ├── inputs.yaml
│   │   │   │   └── terragrunt.hcl
│   │   │   └── region.hcl
│   │   └── account.hcl
│   ├── preqa
│   │   ├── us-east-2
│   │   │   ├── 000
│   │   │   │   ├── inputs.yaml
│   │   │   │   └── terragrunt.hcl
│   │   │   └── region.hcl
│   │   └── account.hcl
│   ├── prod
│   │   ├── us-east-2
│   │   │   ├── 000
│   │   │   │   ├── inputs.yaml
│   │   │   │   └── terragrunt.hcl
│   │   │   └── region.hcl
│   │   └── account.hcl
│   ├── qa
│   │   ├── us-east-2
│   │   │   ├── 000
│   │   │   │   ├── inputs.yaml
│   │   │   │   └── terragrunt.hcl
│   │   │   └── region.hcl
│   │   └── account.hcl
│   ├── sandbox
│   │   ├── us-east-2
│   │   │   ├── 000
│   │   │   │   ├── inputs.yaml
│   │   │   │   └── terragrunt.hcl
│   │   │   └── region.hcl
│   │   └── account.hcl
│   └── uat
│       ├── us-east-2
│       │   ├── 000
│       │   │   ├── inputs.yaml
│       │   │   └── terragrunt.hcl
│       │   └── region.hcl
│       └── account.hcl
├── README.md
├── accounts.json
├── common.hcl
├── inputs.yaml
└── terragrunt.hcl
```
