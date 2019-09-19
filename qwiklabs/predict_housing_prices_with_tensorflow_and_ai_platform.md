---
title: "Predict Housing Prices with Tensorflow and AI Platform"
tags: google-cloud-platform, tensorflow, analytics
url: https://www.qwiklabs.com/focuses/3644
---

# Goal
- Predict Housing Prices with Tensorflow and AI Platform

# Task
- [x] Datalab setup
- [x] Download lab notebook
- [x] Open and execute the housing prices notebook

# Supplement
## Datalab setup
```sh
gcloud config set core/project qwiklabs-gcp-ee2430defe5a5a88
gcloud config set compute/zone us-central1-f
gsutil mb -c multi_regional -l us gs://qwiklabs-gcp-ee2430defe5a5a88-bucket
datalab create my-datalab --machine-type n1-standard-4
```

## Download lab notebook
```sh
!git clone https://github.com/vijaykyr/tensorflow_teaching_examples.git housing_prices
Cloning into 'housing_prices'...
remote: Enumerating objects: 160, done.
remote: Total 160 (delta 0), reused 0 (delta 0), pack-reused 160
Receiving objects: 100% (160/160), 292.25 KiB | 0 bytes/s, done.
Resolving deltas: 100% (70/70), done.
Checking connectivity... done.
```

## Open and execute the housing prices notebook
- https://github.com/vijaykyr/tensorflow_teaching_examples/blob/master/housing_prices/cloud-ml-housing-prices.ipynb

**Reference**
- https://www.tensorflow.org/api_docs/python/tf/feature_column
- https://www.tensorflow.org/guide/custom_estimators
    - https://github.com/GoogleCloudPlatform/training-data-analyst/blob/master/blogs/timeseries/rnn_cloudmle.ipynb
- https://cloud.google.com/ml-engine/docs/training-overview#job_configuration_parameters
- https://cloud.google.com/ml-engine/docs/using-hyperparameter-tuning

**Trouble shooting**
- https://stackoverflow.com/questions/47068709/your-cpu-supports-instructions-that-this-tensorflow-binary-was-not-compiled-to-u
