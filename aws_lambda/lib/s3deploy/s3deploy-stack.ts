import cdk = require("@aws-cdk/core");
import s3 = require("@aws-cdk/aws-s3");
import s3deploy = require("@aws-cdk/aws-s3-deployment");
import generateUid = require("nanoid/generate");
import secretsmanager = require("@aws-cdk/aws-secretsmanager");

export class AwsLambdaS3deployStack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    let stageName: string = "prod";
    if (this.node.tryGetContext("stage") !== undefined) {
      stageName = this.node.tryGetContext("stage");
    }
    const stackName = this.stackName.toLowerCase();

    // const uid = process.env.CDK_UID;
    // // const uid = generateUid("1234567890abcdefghijklmnopqrstuvwxyz", 13);
    // // new secretsmanager.Secret(this, "s3deploySecretManager", {
    // //   secretName: `${stackName}-${stageName}`,
    // //   generateSecretString: {
    // //     secretStringTemplate: JSON.stringify({ uid: uid }),
    // //     generateStringKey: "secret"
    // //   }
    // // });

    new s3deploy.BucketDeployment(this, "s3deploy", {
      sources: [s3deploy.Source.asset("lambda")],
      destinationBucket: new s3.Bucket(this, "bucket", {
        bucketName: `${stackName}-${stageName}`
        // bucketName: `${stackName}-${stageName}-${uid}`
      })
    });
  }

  public capitalize(s: string) {
    return s.charAt(0) + s.slice(1).toLowerCase;
  }
}
