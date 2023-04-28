terraform {
  source = "git::https://github.com/nexient-llc/tf-aws-wrapper_module-memcached_cluster.git//.?ref=feature/init"
}

locals {
  inputs = yamldecode(file("${get_terragrunt_dir()}/inputs.yaml"))
}

inputs = {
  environment_number = local.inputs.environment_num
  resource_names_map = local.inputs.resource_names_map

  vpc_id                             = local.inputs.vpc_id
  private_subnets                    = local.inputs.private_subnets
  ingress_rules                      = local.inputs.ingress_rules
  egress_rules                       = local.inputs.egress_rules
  cluster_size                       = local.inputs.cluster_size
  instance_type                      = local.inputs.instance_type
  engine_version                     = local.inputs.engine_version
  az_mode                            = local.inputs.az_mode
  elasticache_parameter_group_family = local.inputs.elasticache_parameter_group_family

  tags = local.inputs.tags

}
