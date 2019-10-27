#!/usr/bin/env bash

if (($# != 1)); then
  echo "$0 lambda_name"
  exit 1
fi

lambda_name="$1"
lambda_id=$(
  cdk synth AwsLambdaApigatewayStack --no-staging |
    grep -1 "Type: AWS::Lambda::Function" |
    grep -E "${lambda_name}Lambda.*:" |
    sed -e "s/ *\(.*\):/\1/g"
)

sam local invoke ${lambda_id} \
  --no-event
