#!/usr/bin/env node
import "source-map-support/register";
import cdk = require("@aws-cdk/core");
import { AwsLambdaApigatewayStack } from "../lib/apigateway/apigateway-stack";
import { AwsLambdaS3deployStack } from "../lib/s3deploy/s3deploy-stack";

const app = new cdk.App();
new AwsLambdaApigatewayStack(app, "AwsLambdaApigatewayStack");
new AwsLambdaS3deployStack(app, "AwsLambdaS3deployStack");
