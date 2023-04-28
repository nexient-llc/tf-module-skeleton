include "root" {
  path = find_in_parent_folders()
}

include "common" {
  path = "${get_repo_root()}/common/memcached_cluster.hcl"
}

inputs = {
  # Override inputs go here
}
