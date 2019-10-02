---
title: "Cloud SQL with Terraform"
tags: google-cloud-platform, google-cloud-sql, terraform
url: https://www.qwiklabs.com/focuses/1215
---

# Goal
- Create a Cloud SQL instance
- Install the Cloud SQL Proxy
- Test connectivity with MySQL client using Cloud Shell

# Task
- [x] Setup
- [x] Cloud SQL
- [x] Download necessary files
- [x] Understand the code
- [x] Run Terraform
- [x] Cloud SQL Proxy
- [x] Installing the Cloud SQL Proxy
- [x] Test connection to the database
- [x] Test your Understanding

# Supplement
## Cloud SQL
```sh
terraform version
```

## Download necessary files
```sh
gsutil cp -r gs://spls/gsp234 .
cd ~/gsp234/
unzip gsp234.zip
```

## Understand the code
```sh
cat main.tf
```

## Run Terraform
```sh
terraform init
terraform plan -out=tfplan
terraform apply tfplan
```

## Installing the Cloud SQL Proxy
```sh
wget https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 -O cloud_sql_proxy
chmod +x cloud_sql_proxy
```

## Test connection to the database
```sh
export GOOGLE_PROJECT=$(gcloud config list --format "value(core.project)")
MYSQL_DB_NAME=$(terraform output instance_name)
MYSQL_CONN_NAME="${GOOGLE_PROJECT}:us-central1:${MYSQL_DB_NAME}"

./cloud_sql_proxy -instances=${MYSQL_CONN_NAME}=tcp:3306
```
```sh
cd ~/gsp234
MYSQL_PASSWORD=$(terraform output generated_user_password)
mysql -udefault -p$MYSQL_PASSWORD --host 127.0.0.1 default
```

## Reference
- https://cloud.google.com/sql/docs/mysql/configure-ip
- https://cloud.google.com/sql/docs/mysql/configure-ssl-instance
- https://learn.hashicorp.com/terraform/getting-started/outputs.html
