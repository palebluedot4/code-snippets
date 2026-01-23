variable "environment" {
  description = "Deployment environment identifier (dev, stg, uat, prod)."
  type        = string
  default     = "dev"

  validation {
    condition     = contains(["dev", "stg", "uat", "prod"], var.environment)
    error_message = "Invalid environment. Must be one of: dev, stg, uat, prod."
  }
}

variable "common_tags" {
  description = "Common metadata tags for all resources."
  type        = map(string)
  default = {
    Project   = "quark"
    Owner     = "palebluedot4"
    ManagedBy = "terraform"
  }
}

variable "aws_region" {
  description = "Target AWS region for resource deployment."
  type        = string
  default     = "ap-northeast-1"
}

variable "instance_type" {
  description = "EC2 instance type (must be a Graviton/ARM64 family)."
  type        = string
  default     = "t4g.micro"

  validation {
    condition     = can(regex("^[a-z]+[0-9]+g[a-z]*\\.", var.instance_type))
    error_message = "The instance_type must belong to a Graviton (ARM64) family (e.g., t4g.micro, c8g.large)."
  }
}
