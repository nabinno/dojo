---
title: "Cloud ML Engine: Qwik Start"
tags: google-cloud-platform, google-cloud-ml-engine, machine-learning
url: https://www.qwiklabs.com/focuses/581
---

# Goal
- Create a TensorFlow training application and validate it locally
- Run your training job on a single worker instance in the cloud
- Run your training job as a distributed training job in the cloud
- Optimize your hyperparameters by using hyperparameter tuning
- Deploy a model to support prediction
- Request an online prediction and see the response
- Request a batch prediction

# Task
- [x] Setup and Requirements
- [x] Install TensorFlow
- [x] Clone the example repo
- [x] Develop and validate your training application locally
- [x] Install dependencies
- [x] Run a local training job
- [x] Use your trained model for prediction
- [x] Run your training job in the cloud
- [x] Run a single-instance trainer in the cloud
- [x] Deploy your model to support prediction
- [x] Test your Understanding

# Supplement
![]()

```uml
```

## Install TensorFlow
```sh
pip install --user --upgrade tensorflow
python -c "import tensorflow as tf; print('TensorFlow version {} is installed.'.format(tf.VERSION))"
```

## Clone the example repo
```sh
git clone https://github.com/GoogleCloudPlatform/cloudml-samples.git
cd cloudml-samples/census/estimator
```

## Develop and validate your training application locally
```sh
mkdir data
gsutil -m cp gs://cloud-samples-data/ml-engine/census/data/* data/
export TRAIN_DATA=$(pwd)/data/adult.data.csv
export EVAL_DATA=$(pwd)/data/adult.test.csv
head data/adult.data.csv
```

## Install dependencies
```sh
pip install --user -r ../requirements.txt
```

## Run a local training job
```sh
export MODEL_DIR=output
gcloud ai-platform local train \
    --module-name trainer.task \
    --package-path trainer/ \
    --job-dir $MODEL_DIR \
    -- \
    --train-files $TRAIN_DATA \
    --eval-files $EVAL_DATA \
    --train-steps 1000 \
    --eval-steps 100
tensorboard --logdir=$MODEL_DIR --port=8080
ls output/export/census/
gcloud ai-platform local predict --model-dir output/export/census/$(ls output/export/census/) --json-instances ../test.json
```

## Run your training job in the cloud
```sh
PROJECT_ID=$(gcloud config list project --format "value(core.project)")
BUCKET_NAME=${PROJECT_ID}-mlengine
echo $BUCKET_NAME
REGION=us-central1
gsutil mb -l $REGION gs://$BUCKET_NAME
gsutil cp -r data gs://$BUCKET_NAME/data
TRAIN_DATA=gs://$BUCKET_NAME/data/adult.data.csv
EVAL_DATA=gs://$BUCKET_NAME/data/adult.test.csv
gsutil cp ../test.json gs://$BUCKET_NAME/data/test.json
TEST_JSON=gs://$BUCKET_NAME/data/test.json
```

## Run a single-instance trainer in the cloud
```sh
JOB_NAME=census_single_1
OUTPUT_PATH=gs://$BUCKET_NAME/$JOB_NAME
gcloud ai-platform jobs submit training $JOB_NAME \
    --job-dir $OUTPUT_PATH \
    --runtime-version 1.10 \
    --module-name trainer.task \
    --package-path trainer/ \
    --region $REGION \
    -- \
    --train-files $TRAIN_DATA \
    --eval-files $EVAL_DATA \
    --train-steps 1000 \
    --eval-steps 100 \
    --verbosity DEBUG
gcloud ai-platform jobs stream-logs $JOB_NAME
gsutil ls -r $OUTPUT_PATH
tensorboard --logdir=$OUTPUT_PATH --port=8080
```

## Deploy your model to support prediction
```sh
MODEL_NAME=census
gcloud ai-platform models create $MODEL_NAME --regions=$REGION
gsutil ls -r $OUTPUT_PATH/export
MODEL_BINARIES=$OUTPUT_PATH/export/census/1569452350/
gcloud ai-platform versions create v1 \
    --model $MODEL_NAME \
    --origin $MODEL_BINARIES \
    --runtime-version 1.10
gcloud ai-platform models list
gcloud ai-platform predict \
    --model $MODEL_NAME \
    --version v1 \
    --json-instances ../test.json
```
