#!/usr/bin/env bash

set -e

if (($# != 2)); then
  echo "$0 lambda_resource lambda_action"
  exit 1
fi

if [ ! -f template.yaml ]; then
  cdk synth AwsLambdaApigatewayStack --no-staging >template.yaml
fi

lambda_resource="$1"
lambda_action="$2"
lambda_id=$(
  cat template.yaml |
    grep -1 "Type: AWS::Lambda::Function" |
    grep -E "${lambda_resource}Lambda.*:" |
    sed -e "s/ *\(.*\):/\1/g"
)

invoke() {
  sam local invoke ${lambda_id} \
    -v lambda/${lambda_resource} \
    --debug
}

case ${lambda_resource} in
  authorizer)
    sam local generate-event apigateway authorizer | invoke
    ;;
  *)
    case ${lambda_action} in
      list)
        sam local generate-event apigateway aws-proxy \
          --body {"authorizationToken":"test123"} \
          --method GET \
          --resource / \
          --path ${lambda_resource} | invoke
        ;;
      read)
        sam local generate-event apigateway aws-proxy \
          --body {"authorizationToken":"test123"} \
          --method GET \
          --resource / \
          --path ${lambda_resource}/1 | invoke
        ;;
      create)
        # @todo 2019-10-28
        sam local generate-event apigateway aws-proxy \
          --body {"authorizationToken":"test123"} \
          --method POST \
          --resource / \
          --path ${lambda_resource} | invoke
        ;;
    esac
    ;;
esac
