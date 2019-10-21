import cdk = require("@aws-cdk/core");
import s3 = require("@aws-cdk/aws-s3");
import s3deploy = require("@aws-cdk/aws-s3-deployment");

export class AwsLambdaS3deployStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    let stageName: string = "prod";
    if (this.node.tryGetContext("stage") !== undefined) {
      stageName = this.node.tryGetContext("stage");
    }
    const stackName = this.stackName.toLowerCase();

    new s3deploy.BucketDeployment(this, "s3deploy", {
      sources: [s3deploy.Source.asset("lambda")],
      destinationBucket: new s3.Bucket(this, "bucket", {
        bucketName: `${stageName}-${stackName}`
      })
    });
  }
}
