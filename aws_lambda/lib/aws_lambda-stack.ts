import cdk = require("@aws-cdk/core");
// import * as iam from "@aws-cdk/aws-iam";
import lambda = require("@aws-cdk/aws-lambda");
import apigw = require("@aws-cdk/aws-apigateway");
import s3 = require("@aws-cdk/aws-s3");
import s3deploy = require("@aws-cdk/aws-s3-deployment");

export class AwsLambdaStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const bucket = new s3.Bucket(this, "lambda");

    // @todo 2019-10-15
    new s3deploy.BucketDeployment(this, "DeployToLambdaBucket", {
      sources: [s3deploy.Source.asset("lambda/pets")],
      destinationBucket: bucket
    });

    // GET  /pets
    // GET  /pets/:id
    // POST /pets
    new apigw.LambdaRestApi(this, "apigw-pets", {
      handler: new lambda.Function(this, "lambda-pets", {
        runtime: lambda.Runtime.GO_1_X,
        code: lambda.Code.fromBucket(
          bucket,
          // s3.Bucket.fromBucketName(this, "lambda", "lambda"),
          this.stackName
        ),
        handler: "main",
        memorySize: 256,
        timeout: cdk.Duration.seconds(300),
        environment: {}
      })
    });
  }
}
