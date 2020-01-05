---
title: "Introduction to Amazon Aurora"
tags: amazon-web-services, amazon-aurora, mysql
url: https://www.qwiklabs.com/focuses/7867?parent=catalog
---

# Goal
- Create an Amazon Aurora instance
- Connect a pre-created Amazon RDS for MySQL instance
- Connect a pre-created Amazon EC2 that has MySQL Workbench installed
- Load data into the Amazon RDS for MySQL instance from a dump file
- Load data into the Amazon Aurora instance from the same dump file
- Query the data in the Amazon Aurora instance
- Query the data in the Amazon RDS for MySQL instance and compare the results

# Task
- [x] Overview
- [x] Prerequisites
- [x] Start Lab
- [x] Introducing the Technologies
- [x] Task 1: Create an Amazon Aurora Instance
- [x] Task 2: Connect to Amazon EC2 Windows Instance
- [x] Task 3: Connect to Your Amazon RDS Instance
- [x] Task 4: Import a SQL Dump File Into Your Databases
- [x] Task 5: Query the Databases
- [x] End Lab
- [x] Conclusion
- [x] Additional Resources

# Supplement
## Task 3: Connect to Your Amazon RDS Instance
```
mysql.c6n2tc3vg6gc.us-east-1.rds.amazonaws.com
aurora.cluster-c6n2tc3vg6gc.us-east-1.rds.amazonaws.com
aurora.cluster-ro-c6n2tc3vg6gc.us-east-1.rds.amazonaws.com
```

## Task 4: Import a SQL Dump File Into Your Databases
```powershell
Invoke-WebRequest https://s3-us-west-2.amazonaws.com/aws-tc-largeobjects/SPLs/sharedDatabases/world.sql -OutFile c:\\Users\\Administrator\\Desktop\\world.sql
```

## Task 5: Query the Databases
```sql
SELECT * 
FROM world.city
WHERE Population > 10000
ORDER BY CountryCode;
```

# Reference
- https://aws.amazon.com/rds/details/multi-az/
- https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.WorkingWithRDSInstanceinaVPC.html
- https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_Introduction.html
- https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.Encryption.html
- https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Monitoring.OS.html
- https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-key-pairs.html

