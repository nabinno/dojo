#!/usr/bin/env node
import "source-map-support/register";
import util = require("util");
import cdk = require("@aws-cdk/core");
import { AwsLambdaS3deployStack } from "../lib/s3deploy/s3deploy-stack";
import { AwsLambdaApigatewayStack } from "../lib/apigateway/apigateway-stack";

/**
 * Init CDK Apps
 */
const app = new cdk.App();

new AwsLambdaS3deployStack(app, "AwsLambdaS3deployStack");

const samStack = new AwsLambdaApigatewayStack(app, "AwsLambdaApigatewayStack");
samStack.templateOptions.transforms = ["AWS::Serverless-2016-10-31"];
