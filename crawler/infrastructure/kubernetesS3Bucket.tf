resource "aws_s3_bucket" "kubernetes_s3_bucket" {
    bucket = "kubernetes-crawling-bucket"
}

resource "aws_s3_bucket_server_side_encryption_configuration" "kubernetes_s3_bucket" {
    bucket = aws_s3_bucket.kubernetes_s3_bucket.id

    rule {
        bucket_key_enabled = true
    }
}