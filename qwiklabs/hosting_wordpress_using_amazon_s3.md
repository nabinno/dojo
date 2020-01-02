---
title: Hosting WordPress Using Amazon S3
tags: wordpress,amazon-s3
url: https://www.qwiklabs.com/focuses/347?parent=catalog
---

# Goal
- Configuring WordPress on Amazon EC2.
- Exporting WordPress to static files.
- Copying static files to an Amazon S3 static website.
- Creating a script to send your Wordpress changes to Amazon S3.

# Task
- [x] Overview
- [x] Start Lab
- [x] Task 1: Configure WordPress on Amazon EC2
- [x] Task 2: Create an Amazon S3 static website
- [x] Task 3: Log in to your Amazon EC2 instance
- [x] Task 4: Generate a static version of WordPress
- [x] Task 5: Uploading static WordPress pages to Amazon S3
- [x] Task 6: Using scripts to upload changes to Amazon S3
- [x] Conclusion
- [x] Some things to consider
- [x] End Lab

# Supplement
## Task 2: Create an Amazon S3 static website
```
http://wordpress-ni-2019-01-02.s3-website-us-west-2.amazonaws.com
```

## Task 3: Log in to your Amazon EC2 instance
```
ssh -i .\qwikLABS-L73-10926061.pem ec2-user@34.216.26.208
```

## Task 4: Generate a static version of WordPress
```sh
sudo sed -i.bak -e 's/AllowOverride None/AllowOverride All/g' /etc/httpd/conf/httpd.conf;
sudo service httpd restart
(
  cd /var/www/html/wordpress;
  sudo wget https://us-west-2-aws-training.s3.amazonaws.com/awsu-spl/spl-39/scripts/wpstatic.sh;
  sudo /bin/sh wpstatic.sh -a;
)
sudo chown -R apache:apache /var/www/html/wordpress
```

## Task 5: Uploading static WordPress pages to Amazon S3
```sh
AZ=`curl --silent http://169.254.169.254/latest/meta-data/placement/availability-zone/`
REGION=${AZ::-1}
BUCKET=`aws s3api list-buckets --query "Buckets[?starts_with(Name, 'wordpress-')].Name | [0]" --output text`
aws s3 sync --acl public-read /var/www/html/wordpress/wordpress-static s3://$BUCKET
```

## Task 6: Using scripts to upload changes to Amazon S3
```sh
cat <<EOF >$HOME/wordpress-to-s3.sh;
cd /var/www/html/wordpress;
sudo rm -rf wordpress-static;
sudo /bin/sh wpstatic.sh -a;
aws s3 sync --acl public-read --delete /var/www/html/wordpress/wordpress-static s3://$BUCKET
EOF
chmod 0755 $HOME/wordpress-to-s3.sh;

crontab -l >/tmp/mycron
echo "* * * * * /home/ec2-user/wordpress-to-s3.sh" >>/tmp/mycron
crontab /tmp/mycron
rm /tmp/mycron
```

# References
- https://github.com/chnm/WP-Static
- https://wordpress.org/plugins/static-html-output-plugin/
