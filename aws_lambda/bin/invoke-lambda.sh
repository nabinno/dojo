#!/usr/bin/env bash

if (($# != 1)); then
  echo "$0 lambda_name"
  exit 1
fi

if [ ! -f template.yaml ]; then
  cdk synth AwsLambdaApigatewayStack --no-staging >template.yaml
fi

lambda_name="$1"
lambda_id=$(
  cat template.yaml |
    grep -1 "Type: AWS::Lambda::Function" |
    grep -E "${lambda_name}Lambda.*:" |
    sed -e "s/ *\(.*\):/\1/g"
)

sam local invoke ${lambda_id} \
  -v lambda/${lambda_name} \
  --no-event \
  --debug
