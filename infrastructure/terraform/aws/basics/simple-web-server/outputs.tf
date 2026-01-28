output "instance_id" {
  description = "The unique identifier of the EC2 instance."
  value       = aws_instance.this.id
}

output "instance_arn" {
  description = "The ARN of the EC2 instance."
  value       = aws_instance.this.arn
}

output "instance_public_url" {
  description = "The HTTP URL for accessing the web server."
  value       = "http://${aws_instance.this.public_ip}:${var.server_port}/"
}
