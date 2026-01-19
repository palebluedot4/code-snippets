variable "aws_region" {
  description = "AWS region to deploy resources"
  type        = string
  default     = "ap-northeast-1"
}

variable "instance_type" {
  description = "EC2 instance type (must be an ARM64-compatible type, e.g., t4g.micro)"
  type        = string
  default     = "t4g.micro"

  validation {
    condition     = can(regex("^[a-z]+\\d+g[a-z]*\\.", var.instance_type))
    error_message = "The instance_type must be a Graviton-based (ARM64) instance to match the selected AMI (e.g., t4g.micro)"
  }
}

variable "instance_name" {
  description = "Name tag for the instance"
  type        = string
  default     = "minimal-instance"
}
