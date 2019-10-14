import cdk = require("@aws-cdk/core");
// import * as iam from "@aws-cdk/aws-iam";
import lambda = require("@aws-cdk/aws-lambda");

export class AwsLambdaStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    let stageName: string = "prod";
    if (this.node.tryGetContext("stage") !== undefined) {
      stageName = this.node.tryGetContext("stage");
    }
    const projectName: string = this.node.tryGetContext("project");

    new lambda.Function(this, "SampleIncomingWebhookApp", {
      functionName: `${projectName}-${this.stackName}-${stageName}`,
      runtime: lambda.Runtime.GO_1_X,
      code: lambda.Code.asset("./lambda"),
      handler: "main",
      memorySize: 256,
      timeout: cdk.Duration.seconds(10),
      environment: {}
    });
  }
}
