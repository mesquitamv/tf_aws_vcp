# Module TF_AWS_VPC

Module responsible for creation of VPC on AWS using Terraform

## How to use

You need first clone this repo in you project or use this as module in your git repo.

For creation, you just need pass some variable:

```bash
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
```

And after creation the __VPC ID__ and __VPC CIDR__ will be output for usage by others modules.

## Testing
You can test this module using __Terratest__ (for more information [here](https://terratest.gruntwork.io/)) executing test located on __repo_root/test/aws_vpc__:

```bash
go test
```