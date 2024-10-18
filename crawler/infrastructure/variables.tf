locals {
    account_id = data.aws_caller_identity.current.account_id
}

variable "environment_name" {
    type = string
}

variable "aws_account_id" {
    type = string
}

variable "region" {
    type = string
}

variable "subnet_a_id" {
    type = string
}

variable "subnet_b_id" {
    type = string
}

variable "subnet_c_id" {
    type = string
}

variable "security_group_id" {
    type = string
}




data "aws_caller_identity" "current" {}