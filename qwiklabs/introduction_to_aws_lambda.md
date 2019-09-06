---
title: "Introduction to AWS Lambda"
tags: aws, aws-lambda
url: https://amazon.qwiklabs.com/focuses/6431?parent=catalog
---

# Goal
- Create an AWS Lambda function
- Configure an Amazon S3 bucket as a Lambda Event Source
- Trigger a Lambda function by uploading an object to Amazon S3
- Monitor AWS Lambda S3 functions through Amazon CloudWatch Log

# Task
- [x] Create the Amazon S3 Buckets
- [x] Create an AWS Lambda Function
- [x] Test Your Function
- [x] Monitoring and Logging

# Constraints
- [AWS Lambda Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html) - concurrent executions, layer storage, memory allocation, timeout, environment variables, and so on
- [Does not support FIFO queue](https://dev.classmethod.jp/etc/aws-lambda-support-sqs-event-source/)
- [If AWS Lambda fails to start due to the limit on the number of concurrent executions, the number of SQS messages that have been triggered is counted](https://qiita.com/shibataka000/items/381b607b2d9a70a97a2b)
