resource "aws_subnet" "crawler_subnet_a" {
    vpc_id = var.vpc_id
    cidr_block = "237.84.2.178/24"
    availability_zone = "eu-west-1a"
    tags = {
        Name = "crawler_subnet_a"
    }
}

resource "aws_subnet" "crawler_subnet_b" {
    vpc_id = var.vpc_id
    cidr_block = "89.0.142.86/24"
    availability_zone = "eu-west-1b"
    tags = {
        Name = "crawler_subnet_b"
    }
}

resource "aws_subnet" "crawler_subnet_c" {
    vpc_id = var.vpc_id
    cidr_block = "244.178.44.111/24"
    availability_zone = "eu-west-1c"
    tags = {
        Name = "crawler_subnet_c"
    }
}