---
title: "Introduction to APIs in Google"
tags: google-cloud-platform, google-api, web-api
url: https://www.qwiklabs.com/focuses/3473
---

# Goal
- Google APIs
- API architecture
- HTTP protocol and methods
- Endpoints
- REST (Representational State Transfer) and RESTful APIs
- JSON (JavaScript Object Notation)
- API authentication services

# Task
- [x] Setup and Requirements
- [x] APIs - What and Why
- [x] APIs in GCP
- [x] API Library
- [x] API Dashboard
- [x] API Architecture
- [x] HTTP protocol and request methods
- [x] Endpoints
- [x] RESTful APIs
- [x] API Data Formats (JSON)
- [x] Creating a JSON File in the GCP Console
- [x] Authentication and Authorization
- [x] Authenticate and authorize the Cloud Storage JSON/REST API
- [x] Create a bucket with the Cloud Storage JSON/REST API
- [x] Upload a file using the Cloud Storage JSON/REST API

# Supplement
## Create a bucket with the Cloud Storage JSON/REST API
```sh
nano values.json
nano> {  "name": "my-bucket-2019-09-24",
    >    "location": "us",
    >    "storageClass": "multi_regional"
    > }

export OAUTH2_TOKEN="ya29.Il-MBww9P8CJjNQ1t3d_AeAuOjSqIj9tfyNJNyolK2a7acgU-c2BW65Br4b9OaJwXxxi81RwYLFkG3Urc3CoKXOkhXo5iDdx3zXiHjqOpZsIxmLg1K4Cr20BRW8r6VRAXQ"
export PROJECT_ID=qwiklabs-gcp-d6ed9ade391617ae
curl -X POST --data-binary @values.json -H "Authorization: Bearer $OAUTH2_TOKEN" -H "Content-Type: application/json" "https://www.googleapis.com/storage/v1/b?project=$PROJECT_ID"
```

## Upload a file using the Cloud Storage JSON/REST API
```sh
export OBJECT=$(raelpath demo-image.png)
export BUCKET_NAME=my-bucket-2019-09-24
curl -X POST --data-binary @$OBJECT -H "Authorization: Bearer $OAUTH2_TOKEN" -H "Content-Type: image/png" "https://www.googleapis.com/upload/storage/v1/b/$BUCKET_NAME/o?uploadType=media&name=demo-image"
```