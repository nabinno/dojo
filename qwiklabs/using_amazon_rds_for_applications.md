---
title: "Using Amazon RDS for Applications"
tags: amazon-web-services, amazon-rds, database, drupal, php
url: https://www.qwiklabs.com/focuses/7627?parent=catalog
---

# Goal
- Launch an Amazon RDS database
- Transition an application to use the Amazon RDS database
- Change the Instance Type of an Amazon RDS database
- Configure an Amazon RDS database for High Availability

# Task
- [x] Lab Overview
- [x] Start Lab
- [x] Task 1: Access the Content Management System (CMS)
- [x] Task 2: Launch an Amazon RDS for MySQL Database
- [x] Task 3: Login to your Amazon EC2 Instance
- [x] Task 4: Backup the CMS Database
- [x] Task 5: Transition the Drupal CMS Instance to use the RDS database
- [x] Task 6: Change the RDS Instance Size
- [x] Conclusion
- [x] End Lab
- [x] Additional Resources

# Supplement
## Task 3: Login to your Amazon EC2 Instance
```sh
ssh -i qwikLABS-xxxx.pem ec2-user@xxxx
```

## Task 4: Backup the CMS Database
```sh
mysqldump -u root -padmin123 drupaldb >backup.sql
tail backup.sql
```

## Task 5: Transition the Drupal CMS Instance to use the RDS database
```sh
ENDPOINT=drupallab.ce9tex71soha.us-west-2.rds.amazonaws.com
mysql --user=root --password=admin123 --database=drupaldb --host=${ENDPOINT} <backup.sql

# Modify database host
sudo nano /var/www/html/sites/default/settings.php
```

# References
- https://aws.amazon.com/rds/
- https://www.drupal.org/

