---
title: Dataproc: Qwik Start - Command Line
tags: google-cloud-dataproc,apache-spark
url: https://www.qwiklabs.com/focuses/585?parent=catalog
---

# Goal
- Google Cloud Dataproc

# Task
- [x] Overview
- [x] Setup and Requirements
- [x] Create a cluster
- [x] Submit a job
- [x] Update a cluster
- [x] Test your Understanding
- [x] Congratulations!

# Supplement
## Create a cluster
```sh
gcloud dataproc clusters create example-cluster
```

## Submit a job
```sh
gcloud dataproc jobs submit spark \
  --cluster example-cluster \
  --class org.apache.spark.examples.SparkPi \
  --jars file:///usr/lib/spark/examples/jars/spark-examples.jar \
  -- 1000
```

## Update a cluster
```sh
gcloud dataproc clusters update example-cluster --num-workers 4
gcloud dataproc clusters update example-cluster --num-workers 2
```

# References
- https://cloud.google.com/sdk/gcloud/reference/dataproc/jobs/submit/spark
