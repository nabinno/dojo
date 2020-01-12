---
title: Deploying a Python Flask Web Application to App Engine Flexible
tags: google-cloud-app-engine,flask,python,software-deployment
url: https://www.qwiklabs.com/focuses/3339?parent=catalog
---

# Goal
- Familiarity with Python
- Familiarity with standard Linux text editors such as Vim, Emacs, or Nano
- Access to an image with a face

# Task
- [x] Overview
- [x] Setup and Requirements
- [x] Get the sample code
- [x] Authenticate API Requests
- [x] Testing the Application Locally
- [x] Exploring the Code
- [x] Deploying the App to App Engine Flexible
- [x] Congratulations!

# Supplement
## Setup and Requirements
```sh
git clone https://github.com/GoogleCloudPlatform/python-docs-samples.git
cd python-docs-samples/codelabs/flex_and_vision
```

## Get the sample code
```sh
export PROJECT_ID=qwiklabs-gcp-00-d16ea21f82a4
gcloud iam service-accounts create qwiklab \
  --display-name "My Qwiklab Service Account"
gcloud projects add-iam-policy-binding ${PROJECT_ID} \
  --member serviceAccount:qwiklab@${PROJECT_ID}.iam.gserviceaccount.com \
  --role roles/owner
gcloud iam service-accounts keys create ~/key.json \
  --iam-account qwiklab@${PROJECT_ID}.iam.gserviceaccount.com
export GOOGLE_APPLICATION_CREDENTIALS="/home/${USER}/key.json"
```

## Authenticate API Requests
```sh
## Starting your virtual environment
virtualenv -p python3 env
source env/bin/activate
pip install -r requirements.txt

## Creating an App Engine App
gcloud app create

## Creating a Storage Bucket
export CLOUD_STORAGE_BUCKET=${PROJECT_ID}
gsutil mb gs://${PROJECT_ID}

## Running the Application
python main.py
```

## Exploring the Code
**Sample Code Layout**
```
templates/
  homepage.html   /* HTML template that uses Jinja2 */
app.yaml          /* App Engine application configuration file */
main.py           /* Python Flask web application */
requirements.txt  /* List of dependencies for the project */
```

## Deploying the App to App Engine Flexible
```sh
cat <<EOF >app.yaml
runtime: python
env: flex
entrypoint: gunicorn -b :\$PORT main:app

runtime_config:
  python_version: 3

env_variables:
  CLOUD_STORAGE_BUCKET: ${PROJECT_ID}
EOF

gcloud app deploy

# GAE URL: https://${PROJECT_ID}.appspot.com
```

# References
- https://cloud.google.com/appengine/docs/flexible/python/configuring-your-app-with-app-yaml
- https://cloud.google.com/datastore/docs/concepts/entities
