---
title: "Serverless Architectures using Amazon CloudWatch Events with AWS Lambda"
tags: aws, aws-lambda, amazon-cloudwatch-events, amazon-cloudwatch, amazon-sns
url: https://amazon.qwiklabs.com/focuses/394?parent=catalog
---

# Goal
- Crate an AWS Lambda function
- Configure Amazon CloudWatch Events
- Configure an Amazon CloudWatch Alarm
- Configure Amazon SNS to send notifications

# Task
- [x] Create an AWS Lambda function
- [x] Create an Amazon CloudWatch Events Rule
- [x] Test The Amazon CloudWatch Events Rule and Lambda Function
- [x] Create a Website Checker Lambda Function
- [x] Configure Amazon SNS
- [x] Create an Amazon CloudWatch Alarm

# Supplement
## Lambda
**IAM**
```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:createLogGroup"
      ],
      "Resource": "arn:aws:logs:us-east-2:313925688633:*",
      "Effect": "Allow"
    },
    {
      "Action": [
        "logs:createLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:us-east-2:313925688633:log-group:*",
      "Effect": "Allow"
    }
  ]
}
```

**Function MonitorEC2**
```js
exports.handler = function(event, context) {
  console.log("MonitorEC2()");
  console.log("Here's the event:\n:"+JSON.stringify(event, null, 4));
  context.succeed("Ready!");
};
```

**Function lambda-canary**
```py
import os
from datetime import datetime
from urllib.request import Request, urlopen

SITE = os.environ['site']  # URL of the site to check, stored in the site environment variable
EXPECTED = os.environ['expected']  # String expected to be on the page, stored in the expected environment variable


def validate(res):
    '''Return False to trigger the canary

    Currently this simply checks whether the EXPECTED string is present.
    However, you could modify this to perform any number of arbitrary
    checks on the contents of SITE.
    '''
    return EXPECTED in res


def lambda_handler(event, context):
    print('Checking {} at {}...'.format(SITE, event['time']))
    try:
        req = Request(SITE, headers={'User-Agent': 'AWS Lambda'})
        if not validate(str(urlopen(req).read())):
        raise Exception('Validation failed')
    except:
        print('Check failed!')
                raise
    else:
        print('Check passed!')
                return event['time']
    finally:
    print('Check complete at {}'.format(str(datetime.now())))
```

## CloudWatch Events
**Event pattern**
```json
{
  "source": [
    "aws.ec2"
  ],
  "detail-type": [
    "EC2 Instance State-change Notification"
  ],
  "detail": {
    "state": [
      "running"
    ]
  }
}
```
