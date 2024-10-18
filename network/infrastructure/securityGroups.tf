resource "aws_security_group" "crawler_security_group" {
    name = "crawler_security_group"
    description = "Crawler Security Group"
    vpc_id = var.vpc_id
    ingress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = ["0.0.0.0/0"]
    }
}

//openapi swaggerui preview