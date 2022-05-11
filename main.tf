resource "random_string" "string" {
  length  = var.length
  number  = var.number
  special = var.special
}
