---
title: Introduction to Amazon Kinesis Firehose
tags: amazon-kinesis,amazon-kinesis-firehose,message-queueing-service
url: https://www.qwiklabs.com/focuses/274?catalog_rank=%7B%22rank%22%3A1%2C%22num_filters%22%3A0%2C%22has_search%22%3Atrue%7D&parent=catalog&search_id=4126127
---

# Goal
- Create an AWS Lambda data transformation function.
- Create an Amazon Kinesis Firehose delivery stream connected to an Elasticsearch cluster.
- Send data to the delivery stream.
- Visualize streaming data with Kibana.

# Task
- [x] Overview
- [x] Topics covered
- [x] Start Lab
- [x] Scenario
- [x] Introducing the Technologies
- [x] Task 1: Create your AWS Lambda Transformation Function
- [x] Task 2: Create your Amazon Kinesis Firehose
- [x] Task 3: Access your Elasticsearch Cluster
- [x] Task 4: Send Data to the Delivery Stream
- [x] Task 5: Visualize the Data in Kibana
- [x] Task 6: Monitor the Delivery Stream with Amazon CloudWatch
- [x] Conclusion
- [x] End Lab
- [x] Additional Resources

# Supplement
## Overview
![]()

## Task 1: Create your AWS Lambda Transformation Function
**AWS Lambda**

- Runtime: Python 2.7
- Handler: lambda_function.lambda_handler
- Timeout: 1 minute

```python
import base64
import json
from datetime import datetime

# Incoming Event
def lambda_handler(event, context):
    output = []
    now = datetime.utcnow().isoformat()

    # Loop through records in incoming Event
    for record in event['records']:

        # Extract message
        message = json.loads(base64.b64decode(record['data']))

        # Construct output
        data_field = {
                'timestamp': now,
                'ticker_symbol': message['ticker_symbol'],
                'price': message['price']
        }
        output_record = {
                'recordId': record['recordId'],
                'result': 'Ok',
                'data': base64.b64encode(json.dumps(data_field))
        }
        output.append(output_record)

    return {'records': output}
```

# Reference
## Amazon Kinesis Firehose
- https://www.youtube.com/watch?v=814aUb5n_Fk

## AWS Lambda
- https://www.youtube.com/watch?v=eOBq__h4OJ4
