terraform {
  required_version = ">= 1.7.0"
}

variable "environment" {
  type    = string
  default = "dev"
}

output "service_name" {
  value = "giftbox-${var.environment}"
}
