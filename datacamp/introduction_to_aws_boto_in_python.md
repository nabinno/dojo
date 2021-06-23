---
title: Introduction to AWS Boto in Python
tags: python,boto,boto3
url: https://www.datacamp.com/courses/introduction-to-aws-boto-in-python
---

# 1. Putting Files in the Cloud!
## Your first boto3 client
```python
# Generate the boto3 client for interacting with S3
s3 = boto3.client('s3', region_name='us-east-1', 
                        # Set up AWS credentials 
                        aws_access_key_id=AWS_KEY_ID,
                         aws_secret_access_key=AWS_SECRET)
# List the buckets
buckets = s3.list_buckets()

# Print the buckets
print(buckets)
```

## Multiple clients
```python
# Generate the boto3 client for interacting with S3 and SNS
s3 = boto3.client('s3', region_name='us-east-1',
                         aws_access_key_id=AWS_KEY_ID, 
                         aws_secret_access_key=AWS_SECRET)

sns = boto3.client('sns', region_name='us-east-1',
                         aws_access_key_id=AWS_KEY_ID, 
                         aws_secret_access_key=AWS_SECRET)

# List S3 buckets and SNS topics
buckets = s3.list_buckets()
topics = sns.list_topics()

# Print out the list of SNS topics
print(topics)
```

## Creating a bucket
```python
import boto3

# Create boto3 client to S3
s3 = boto3.client('s3', region_name='us-east-1', 
                         aws_access_key_id=AWS_KEY_ID, 
                         aws_secret_access_key=AWS_SECRET)

# Create the buckets
response_staging = s3.create_bucket(Bucket='gim-staging')
response_processed = s3.create_bucket(Bucket='gim-processed')
response_test = s3.create_bucket(Bucket='gim-test')

# Print out the response
print(response_staging)
```

## Listing buckets
```python
# Get the list_buckets response
response = s3.list_buckets()

# Iterate over Buckets from .list_buckets() response
for bucket in response['Buckets']:
  
  	# Print the Name for each bucket
    print(bucket['Name'])
```

## Deleting a bucket
```python
# Delete the gim-test bucket
s3.delete_bucket(Bucket='gim-test')

# Get the list_buckets response
response = s3.list_buckets()

# Print each Buckets Name
for bucket in response['Buckets']:
    print(bucket['Name'])
```

## Deleting multiple buckets
```python

```

## Uploading and retrieving files
```python

```

## Putting files in the cloud
```python

```

## Spring cleaning
```python

```






# 2. Sharing Files Securely
## Keeping objects secure
```python

```

## Making an object public
```python

```

## Uploading a public report
```python

```

## Making multiple files public
```python

```

## Accessing private objects in S3
```python

```

## Generating a presigned URL
```python

```

## Opening a private file
```python

```

## Sharing files through a website
```python

```

## Generate HTML table from Pandas
```python

```

## Upload an HTML file to S3
```python

```





# 3. Reporting and Notifying!
## SNS Topics
```python

```

## Creating a Topic
```python

```

## Creating multiple topics
```python

```

## Deleting multiple topics
```python

```

## SNS Subscriptions
```python

```

## Subscribing to topics
```python

```

## Creating multiple subscriptions
```python

```

## Deleting multiple subscriptions
```python

```

## Sending messages
```python

```

## Sending an alert
```python

```

## Sending a single SMS message
```python

```

## Case Study: Building a notification system
```python

```

## Creating multi-level topics
```python

```

## Different protocols per topic level
```python

```

## Sending multi-level alerts
```python

```




# 4. Pattern Rekognition
## Rekognizing patterns
```python

```

## Cat detector
```python

```

## Multiple cat detector
```python

```

## Parking sign reader
```python

```

## Comprehending text
```python

```

## Detecting language
```python

```

## Translating Get It Done requests
```python

```

## Getting request sentiment
```python

```

## Case Study: Scooting Around!
```python

```

## Scooter community sentiment
```python

```

## Scooter dispatch
```python

```

## Wrap up
```python

```

