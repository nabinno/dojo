import cdk = require("@aws-cdk/core");
import dynamodb = require("@aws-cdk/aws-dynamodb");
import { AwsLambdaApigatewayStack } from "./apigateway-stack";

/**
 * @desc Purpose: serverless, pay as you go, persistent storage for the demo app
 * @see https://docs.aws.amazon.com/cdk/api/latest/docs/aws-dynamodb-readme.html
 */
export class TableConstruct extends cdk.Construct {
  public readonly items: dynamodb.Table;
  public readonly users: dynamodb.Table;

  constructor(scope: AwsLambdaApigatewayStack, id: string) {
    super(scope, id);

    this.items = new dynamodb.Table(scope, "itemsTable", {
      tableName: `${scope.stackName}-Items`,
      billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
      serverSideEncryption: true,
      stream: dynamodb.StreamViewType.NEW_AND_OLD_IMAGES,
      partitionKey: { name: "id", type: dynamodb.AttributeType.STRING }
    });
    this.users = new dynamodb.Table(scope, "usersTable", {
      tableName: `${scope.stackName}-Users`,
      billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
      serverSideEncryption: true,
      stream: dynamodb.StreamViewType.NEW_AND_OLD_IMAGES,
      partitionKey: { name: "username", type: dynamodb.AttributeType.STRING },
      timeToLiveAttribute: "ttl"
    });
  }
}
