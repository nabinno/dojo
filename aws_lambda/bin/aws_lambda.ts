#!/usr/bin/env node
import "source-map-support/register";
import cdk = require("@aws-cdk/core");
import { AwsLambdaStack } from "../lib/aws_lambda-stack";

const app = new cdk.App();
new AwsLambdaStack(app, "AwsLambdaStack");
