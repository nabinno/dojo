---
title: "Serverless Web Apps using Amazon DynamoDB - Part 1"
tags: aws, amazon-dynamodb, aws-lambda
url: https://amazon.qwiklabs.com/focuses/395?parent=catalog
---

# Goal
- Create an Amazon DynamoDB table
- Add items to your Amazon DynamoDB table
- Understand the structure of IAM roles and policies needed to access your table

# Task
- [x] Create Your DynamoDB Table
- [x] Add Items via Tree Method
- [x] Add Items via JSON
- [x] Review IAM Policies and Roles

# Supplement
## IAM
Grant access to the Scan, BatchWriteItem APIs under DynamoDB and GetObject, PutObject APIs under S3 in your account
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "dynamodb:Scan",
        "s3:GetObject",
        "s3:PutObject",
        "dynamodb:BatchWriteItem"
      ],
      "Resource": [
        "*"
      ],
      "Effect": "Allow"
    }
  ]
}
```

Grant access with query policy
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Condition": {
        "ForAllValues:StringEquals": {
          "dynamodb:Attributes": [
            "SuperHero",
            "MissionStatus",
            "Villain1",
            "Villain2",
            "Villain3"
          ]
        }
      },
      "Action": [
        "dynamodb:Query"
      ],
      "Resource": "*",
      "Effect": "Allow"
    }
  ]
}
```
