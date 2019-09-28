---
title: "Introduction to Amazon EC2"
tags: amazon-web-services, amazon-ec2
url: https://www.qwiklabs.com/focuses/7783
---

# Goal
- Launch a web server with termination protection enabled
- Monitor Your EC2 instance
- Modify the security group that your web server is using to allow HTTP access
- Resize your Amazon EC2 instance to scale
- Explore EC2 limits
- Test termination protection
- Terminate your EC2 instance

# Task
- [x] Task 1: Launch Your Amazon EC2 Instance
- [x] Task 2: Monitor Your Instance
- [x] Task 3: Update Your Security Group and Access the Web Server
- [x] Task 4: Resize Your Instance: Instance Type and EBS Volume
- [x] Task 5: Explore EC2 Limits
- [x] Task 6: Test Termination Protection

# Supplement
## Task 1: Launch Your Amazon EC2 Instance
```sh
#!/bin/bash
yum -y install httpd
systemctl enable httpd
systemctl start httpd
echo '<html><h1>Hello From Your Web Server!</h1></html>' > /var/www/html/index.html
```

## Reference
- https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/LaunchingAndUsingInstances.html
- https://aws.amazon.com/jp/ec2/instance-types/
- https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/AMIs.html
- https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/user-data.html
- https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/RootDeviceStorage.html
- https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/Using_Tags.html
- https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/using-network-security.html
- https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ec2-key-pairs.html
- https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/instance-console.html
- https://docs.aws.amazon.com/ja_jp/AmazonCloudWatch/latest/monitoring/aws-services-cloudwatch-metrics.html
- https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ec2-instance-resize.html
- https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/Stop_Start.html
- https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/ec2-resource-limits.html
- https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/terminating-instances.html
- https://docs.aws.amazon.com/ja_jp/AWSEC2/latest/UserGuide/monitoring-system-instance-status-check.html?icmpid=docs_ec2_console
