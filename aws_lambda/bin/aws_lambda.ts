#!/usr/bin/env node
import "source-map-support/register";
import cdk = require("@aws-cdk/core");
import { AwsLambdaStack } from "../lib/aws_lambda-stack";
import util = require("util");

const exec = util.promisify(require("child_process").exec);

async function execute() {
  await exec(
    `GOOS=linux
GOARCH=amd64
go get -v -t -d ./lambda/...
go build -o ./lambda/pets/main ./lambda/pets/**.go`
  );

  const app = new cdk.App();
  new AwsLambdaStack(app, "AwsLambdaStack");

  await exec("rm ./lambda/pets/main");
}

execute();
