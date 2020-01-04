---
title: Dataflow: Qwik Start - Python
tags: google-cloud-dataflow,apache-beam,python
url: https://www.qwiklabs.com/focuses/1100?parent=catalog
---

# Goal
- Google Cloud Dataflow

# Task
- [x] Setup and Requirements
- [x] Create a Cloud Storage bucket
- [x] Check that your job succeeded
- [x] Test your Understanding

# Supplement
## Setup and Requirements
```sh
python --version
pip --version
sudo pip install -U pip
sudo pip install --upgrade virtualenv
virtualenv -p python2.7 env
source env/bin/activate
pip install apache-beam[gcp]

python -m apache_beam.examples.wordcount --output OUTPUT_FILE

BUCKET=gs://qwiklabs-2019-01-04
python \
  -m apache_beam.examples.wordcount \
  --project $DEVSHELL_PROJECT_ID \
  --runner DataflowRunner \
  --staging_location $BUCKET/staging \
  --temp_location $BUCKET/temp \
  --output $BUCKET/results/output
```

# References
- https://beam.apache.org/get-started/wordcount-example/
- http://shop.oreilly.com/product/0636920057628.do
