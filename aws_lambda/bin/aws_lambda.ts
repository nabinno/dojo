#!/usr/bin/env node
import "source-map-support/register";
import util = require("util");
import cdk = require("@aws-cdk/core");
import { AwsLambdaS3deployStack } from "../lib/s3deploy/s3deploy-stack";
import { AwsLambdaApigatewayStack } from "../lib/apigateway/apigateway-stack";

/**
 * Pre-set UID
 */
async function setFile() {
  const exec = util.promisify(require("child_process").exec);
  await exec(`mkdir -p tmp/cache && touch tmp/cache/uid`);
}
setFile();

/**
 * Init CDK Apps
 */
const app = new cdk.App();
new AwsLambdaS3deployStack(app, "AwsLambdaS3deployStack");
new AwsLambdaApigatewayStack(app, "AwsLambdaApigatewayStack");
