output "vpc_id" {
  value = aws_vpc.vpc.id
}

## Terratest validation
output "vpc_cidr_block" {
  value = aws_vpc.vpc.cidr_block
}
