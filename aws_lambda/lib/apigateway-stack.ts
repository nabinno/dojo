import cdk = require("@aws-cdk/core");
import lambda = require("@aws-cdk/aws-lambda");
import apigw = require("@aws-cdk/aws-apigateway");
import s3 = require("@aws-cdk/aws-s3");

export class AwsLambdaApigatewayStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    let stageName: string = "Prod";
    if (this.node.tryGetContext("stage") !== undefined) {
      stageName = this.node.tryGetContext("stage");
    }
    const stackName = this.stackName;

    // GET  /pets
    // GET  /pets/:id
    // POST /pets
    new apigw.LambdaRestApi(this, "gateway", {
      restApiName: `${stageName}-${stackName}`,
      handler: new lambda.Function(this, "lambda", {
        functionName: `${stageName}-${stackName}-Pets`,
        runtime: lambda.Runtime.GO_1_X,
        code: lambda.Code.fromBucket(
          s3.Bucket.fromBucketName(
            this,
            "s3bucket",
            `${stageName.toLowerCase()}-awslambdas3deploystack`
          ),
          "pets/pets.zip"
        ),
        handler: "pets",
        memorySize: 256,
        timeout: cdk.Duration.seconds(300),
        environment: {}
      })
    });
  }
}
