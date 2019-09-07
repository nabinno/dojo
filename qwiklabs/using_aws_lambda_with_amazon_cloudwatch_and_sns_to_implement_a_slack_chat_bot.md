---
title: "Using AWS Lambda with Amazon CloudWatch and SNS to Implement a Slack Chat Bot"
tags: aws, aws-lambda, amazon-cloudwatch, amazon-sns, slack
url: https://amazon.qwiklabs.com/focuses/1813?parent=catalog
---

# Goal
- Create a Slack chat bot using a Lambda blueprint
- Configure the bot with a Slack webhook to post messages to a Slack channel

# Task
- [x] Create Your Slack Account
- [x] Configure Incoming WebHooks For Slack
- [x] Create and Subscribe to an SNS Topic
- [x] Craete a Lambda Function
- [x] Test your Lambda function
- [x] Create a CloudWatch Alarm
- [x] Test your Lambda function

# Supplement
## Lambda
**IAM role**
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:*",
        "logs:*"
      ],
      "Resource": [
        "*"
      ],
      "Effect": "Allow"
    },
    {
      "Condition": {
        "StringNotEquals": {
          "ec2:InstanceType": [
            "t2.micro",
            "t1.micro"
          ]
        }
      },
      "Action": [
        "ec2:RunInstances"
      ],
      "Resource": "arn:aws:ec2:*:*:instance/*",
      "Effect": "Deny"
    },
    {
      "Action": [
        "ec2:*Spot*"
      ],
      "Resource": "*",
      "Effect": "Deny"
    }
  ]
}
```

**Function Slackfunction**
Version: Python 2.7
```python
from __future__ import print_function
import boto3
import json
import logging
from base64 import b64decode
from urllib2 import Request, urlopen, URLError, HTTPError

SLACK_CHANNEL = 'CHANNEL'
HOOK_URL = "HOOKURL"

logger = logging.getLogger()
logger.setLevel(logging.INFO)


def lambda_handler(event, context):
    logger.info("Event: " + str(event))
    message = event['Records'][0]['Sns']['Message']
    logger.info("Message: " + str(message))
    alarm_name = message['AlarmName']
    #old_state = message['OldStateValue']
    new_state = message['NewStateValue']
    reason = message['NewStateReason']
    slack_message = {
        'channel': SLACK_CHANNEL,
        'text': "%s state is now %s: %s" % (alarm_name, new_state, reason)
    }
    req = Request(HOOK_URL, json.dumps(slack_message))
    try:
        response = urlopen(req)
        response.read()
        logger.info("Message posted to %s", slack_message['channel'])
    except HTTPError as e:
        logger.error("Request failed: %d %s", e.code, e.reason)
    except URLError as e:
        logger.error("Server connection failed: %s", e.reason)
```

**Test**
Template: Amazon SNS Topic Notifiction
```json
{
  "Records": [
    {
      "EventVersion": "1.0",
      "EventSubscriptionArn": "arn:aws:sns:EXAMPLE",
      "EventSource": "aws:sns",
      "Sns": {
        "SignatureVersion": "1",
        "Timestamp": "1970-01-01T00:00:00.000Z",
        "Signature": "EXAMPLE",
        "SigningCertUrl": "EXAMPLE",
        "MessageId": "95df01b4-ee98-5cb9-9903-4c221d41eb5e",
        "Message": {
          "AlarmName": "SlackAlarm",
          "NewStateValue": "OK",
          "NewStateReason": "Threshold Crossed: 1 datapoint (0.0) was not greater than or equal to the threshold (1.0)."
        },
        "MessageAttributes": {
          "Test": {
            "Type": "String",
            "Value": "TestString"
          },
          "TestBinary": {
            "Type": "Binary",
            "Value": "TestBinary"
          }
        },
        "Type": "Notification",
        "UnsubscribeUrl": "EXAMPLE",
        "TopicArn": "arn:aws:sns:EXAMPLE",
        "Subject": "TestInvoke"
      }
    }
  ]
}
```
