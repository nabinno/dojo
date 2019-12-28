---
title: Analyze Big Data with Hadoop
tags: apache-hadoop
url: https://www.qwiklabs.com/focuses/3660?parent=catalog
---

# Goal
- Launch a fully functional Hadoop cluster using Amazon EMR
- Define the schema and create a table for sample log data stored in Amazon S3
- Analyze the data using a HiveQL script and write the results back to Amazon S3
- Download and view the results on your computer

# Task
- [x] Overview
- [x] Start Lab
- [x] Task 1: Create an Amazon S3 bucket
- [x] Task 2: Launch an Amazon EMR cluster
- [x] Task 3: Process Your Sample Data by Running a Hive Script
- [x] Task 4: View the Results
- [x] Task 5: Terminate your Amazon EMR Cluster
- [x] Conclusion
- [x] End Lab
- [x] Additional Resources

# Supplementr
## Task 1: Create an Amazon S3 bucket
```
s3://hadoop-2019-12-28/
```

## Task 3: Process Your Sample Data by Running a Hive Script
```
hive-script \
  --run-hive-script \
  --args \
  -f s3://us-west-2.elasticmapreduce.samples/cloudfront/code/Hive_CloudFront.q \
  -d INPUT=s3://us-west-2.elasticmapreduce.samples \
  -d OUTPUT=s3://hadoop-2019-12-28/ \
  -hiveconf hive.support.sql11.reserved.keywords=false
```

### Create the table
```sql
CREATE EXTERNAL TABLE IF NOT EXISTS cloudfront_logs (
  DateObject Date,
  Time STRING,
  Location STRING,
  Bytes INT,
  RequestIP STRING,
  Method STRING,
  Host STRING,
  Uri STRING,
  Status INT,
  Referrer STRING,
  OS String,
  Browser String,
  BrowserVersion String
);
```

### Parse the log files using the RegEx SerDe
```sql
ROW FORMAT SERDE 'org.apache.hadoop.hive.serde2.RegexSerDe'
WITH SERDEPROPERTIES (
  "input.regex" = "^(?!#)([^ ]+)\\s+([^ ]+)\\s+([^ ]+)\\s+([^ ]+)\\s+([^ ]+)\\s+([^ ]+)\\s+([^ ]+)\\s+([^ ]+)\\s+([^ ]+)\\s+([^ ]+)\\s+[^\(]+[\(](https://s3-us-west-2.amazonaws.com/us-west-2-aws-training/awsu-spl/spl-166/instructions/en_us/[^\;]+).*\%20([^\/]+)[\/](https://s3-us-west-2.amazonaws.com/us-west-2-aws-training/awsu-spl/spl-166/instructions/en_us/.*)$"
) LOCATION '${INPUT}/cloudfront/data/';
```

### Calculate requests by operating system
```sql
INSERT OVERWRITE DIRECTORY '${OUTPUT}/os_requests/'
SELECT
  os,
  COUNT(*) count
FROM cloudfront_logs
WHERE dateobject BETWEEN '2014-07-05' AND '2014-08-05'
GROUP BY os;
```

## Task 4: View the Results
```
Android855
Linux813
MacOS852
OSX799
Windows883
iOS794
```

# References
- https://aws.amazon.com/jp/big-data/datalakes-and-analytics/
