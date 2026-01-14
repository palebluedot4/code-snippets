provider "aws" {
  region = var.aws_region

  default_tags {
    tags = {
      Project   = "code-snippets"
      Component = "minimal-ec2"
      ManagedBy = "Terraform"
    }
  }
}
