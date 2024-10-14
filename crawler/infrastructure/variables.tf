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


data "aws_caller_identity" "current" {}