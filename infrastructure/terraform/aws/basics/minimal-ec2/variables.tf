variable "aws_region" {
  description = "AWS region to deploy resources"
  type        = string
  default     = "ap-northeast-1"
}

variable "instance_type" {
  description = "EC2 instance type"
  type        = string
  default     = "t4g.micro"
}

variable "instance_name" {
  description = "Name tag for the instance"
  type        = string
  default     = "minimal-instance"
}
