data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["al2023-ami-2023.*-arm64"]
  }

  filter {
    name   = "architecture"
    values = ["arm64"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  lifecycle {
    postcondition {
      condition     = self.architecture == "arm64"
      error_message = "The selected AMI must be ARM64 architecture."
    }

    postcondition {
      condition     = self.root_device_type == "ebs"
      error_message = "The selected AMI must be EBS-backed."
    }
  }
}

resource "aws_instance" "this" {
  ami           = data.aws_ami.amazon_linux.id
  instance_type = var.instance_type

  associate_public_ip_address = true

  root_block_device {
    encrypted   = true
    volume_type = "gp3"
  }

  metadata_options {
    http_endpoint               = "enabled"
    http_tokens                 = "required"
    http_put_response_hop_limit = 1
  }

  tags = {
    Name = "minimal-ec2-instance"
  }
}
