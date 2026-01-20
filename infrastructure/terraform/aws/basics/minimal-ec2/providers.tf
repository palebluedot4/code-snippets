provider "aws" {
  region = var.aws_region

  default_tags {
    tags = {
      Project   = "quark"
      Component = "minimal-ec2"
      ManagedBy = "terraform"
    }
  }
}
