variable "ec2_key_name" {
  description = "AWS EC2 Key name for SSH access"
  type        = string
  default     = "coreint-ssh-key"
}

variable "prefix" {
  description = "Prefix for resources created by this module"
  type        = string
  default     = "tf-aws-bigip"
}
variable "region" {}

variable "cidr" {
  description = "aws VPC CIDR"
  type        = string
  default     = "10.0.0.0/16"
}


variable "availabilityZones" {
  description = "If you want the VM placed in an AWS Availability Zone, and the AWS region you are deploying to supports it, specify the numbers of the existing Availability Zone you want to use."
  type        = list(any)
  default     = ["us-east-1a", "us-east-1b"]
}
variable "AllowedIPs" {}

variable "instance_count" {
  description = "Number of Bigip instances to create( From terraform 0.13, module supports count feature to spin mutliple instances )"
  type        = number
  default     = 1
}

variable "f5_ami_search_name" {
  description = "BIG-IP AMI name to search for"
  type        = string
  default     = "F5 BIGIP-16* PAYG-Good 25Mbps*"
  #default     = "F5 BIGIP-16.1.2.1-0.0.10 PAYG-Good 25Mbps-211222200303-8e1217d4-a046-4cdf-894e-e38175bae37f"
}

variable "ec2_instance_type" {
  description = "AWS EC2 instance type"
  type        = string
  default     = "t2.medium"
}
