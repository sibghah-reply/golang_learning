resource "aws_kms_key" "crawler_url_queue" {
    description = "Crawler URL Queue"
    enable_key_rotation = true
}

resource "aws_elasticache_serverless_cache" "url_queue" {
    engine = "redis"
    name = "crawlerUrlQueue"
    cache_usage_limits {
        data_storage {
            maximum = 10
            unit = "GB"
        }
        ecpu_per_second {
           maximum = 5000
        }
    }
    daily_snapshot_time       = "09:00"
    description               = "Crawler URL Queue"
    kms_key_id                = aws_kms_key.crawler_url_queue.arn
    major_engine_version      = 7
    snapshot_retention_limit  = 1
    security_group_ids        = [var.security_group_id]
    subnet_ids                = [var.subnet_a_id, var.subnet_b_id, var.subnet_c_id] 

}