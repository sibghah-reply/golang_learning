terraform {
    required_providers {
      aws = {
        source = "hashicorp/aws"
        version = "~> 4.0"
      }
    }
}

provider "aws" {
    region = var.region
    allowed_account_ids = [var.aws_account_id]
    assume_role {
        role_arn = "arn:aws:iam::${var.aws_account_id}:role/RoleForTerraform"
    }
    default_tags {
        tags = {
            EnvironmentUse        = var.environment_name
            EnvironmentGroup      = "PersonalProjectCrawlerScraper"
            SupportTeam           = "si.khan@reply.com"
        }
    }
}