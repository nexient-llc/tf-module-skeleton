output "string" {
  value = format("%s🍰%s", module.cake_prefix.string, module.cake_suffix.string)
}
