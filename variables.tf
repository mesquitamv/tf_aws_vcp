# Region on AWS
variable "aws_region" {
  type        = string
  description = "AWS Region"
  default     = "us-east-2"
}

# CIDR of new VPC
variable "vpc_cidr" {
  type        = string
  description = "CIDR block for VPC"
  default     = "10.0.0.0/16"
}

# Map With tags that you want/need
variable "tags" {
  default = {
    environment = "test"
  }
}