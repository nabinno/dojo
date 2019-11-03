import cdk = require("@aws-cdk/core");
import apigw = require("@aws-cdk/aws-apigateway");
import lambda = require("@aws-cdk/aws-lambda");
import s3 = require("@aws-cdk/aws-s3");
import iam = require("@aws-cdk/aws-iam");
import { LambdaAuthorizerConstruct } from "./lambda-authorizer-construct";
import { CognitoAuthorizerConstruct } from "./cognito-authorizer-construct";
import { UserPoolConstruct } from "./user-pool-construct";
import { TableConstruct } from "./table-construct";
import { PetsLambdaConstruct } from "./pets-lambda-construct";
import { Utils } from "../shared/utils";
import { URL } from "url";

export class AwsLambdaApigatewayStack extends cdk.Stack {
  public readonly stageName: string;
  public readonly api: apigw.RestApi;
  public readonly apiLambdaPermission: lambda.Permission;
  public readonly apiRequestValidator: apigw.RequestValidator;
  public readonly apiLambdaAuthorizer: LambdaAuthorizerConstruct;
  public readonly apiCognitoAuthorizer: CognitoAuthorizerConstruct;
  public readonly userPool: UserPoolConstruct;
  public readonly table: TableConstruct;
  public readonly lambdaBucket: s3.IBucket;
  public readonly petsLambda: PetsLambdaConstruct;
  public envDomain: string;
  public envIdentityProviderMetadataURLOrFile: string;
  public envAppUrl: string;
  public envCallbackURL: string;
  public envAuthorizationHeaderName: string;
  public envGroupsAttributeName: string;
  public envGroupsAttributeClaimName: string;
  public envRoleNameOfAdmins: string;
  public envRoleNameOfUsers: string;
  public envNodeRuntime: lambda.Runtime;

  constructor(scope: cdk.Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    this.stageName = "prod";
    if (this.node.tryGetContext("stage") !== undefined) {
      this.stageName = this.node.tryGetContext("stage");
    }
    this.loadEnvironmentVariables();

    this.api = new apigw.RestApi(this, "gateway", { restApiName: this.stackName });
    this.apiRequestValidator = new apigw.RequestValidator(this, "requestvalidator", {
      restApi: this.api,
      requestValidatorName: "validateRequestBodyAndParameters",
      validateRequestBody: true,
      validateRequestParameters: false
    });
    this.apiLambdaPermission = {
      principal: new iam.ServicePrincipal("apigateway.amazonaws.com"),
      sourceArn: `arn:aws:execute-api:${this.region}:${this.account}:${this.api.restApiId}/*`
    };
    this.lambdaBucket = s3.Bucket.fromBucketName(
      this,
      "s3bucket",
      `awslambdas3deploystack-${this.stageName}`
    );

    this.userPool = new UserPoolConstruct(this, "userPoolConstruct");
    this.table = new TableConstruct(this, "tableConstruct");

    this.apiLambdaAuthorizer = new LambdaAuthorizerConstruct(this, "lambdaAuthorizerConstruct");
    this.apiCognitoAuthorizer = new CognitoAuthorizerConstruct(
      this,
      "cognitoAuthorizerConstruct"
    );

    this.petsLambda = new PetsLambdaConstruct(this, "petsLambdaConstruct");

    this.outputStack();
  }

  private loadEnvironmentVariables() {
    this.envDomain = Utils.getEnv("COGNITO_DOMAIN_NAME", `2019-10-31-${this.stageName}`);
    this.envIdentityProviderMetadataURLOrFile = Utils.getEnv("IDENTITY_PROVIDER_METADATA", "");
    this.envAppUrl = Utils.getEnv("APP_URL", "http://localhost:3000");
    new URL(this.envAppUrl); // validate URL (throws if invalid URL
    this.envCallbackURL = `${this.envAppUrl}/idpcallback`;
    this.envAuthorizationHeaderName = "Authorization";
    this.envGroupsAttributeName = Utils.getEnv("GROUPS_ATTRIBUTE_NAME", "groups");
    this.envGroupsAttributeClaimName = `custom:${this.envGroupsAttributeName}`;
    this.envRoleNameOfAdmins = Utils.getEnv("ROLE_NAME_OF_ADMINS", "pet-app-admins");
    this.envRoleNameOfUsers = Utils.getEnv("ROLE_NAME_OF_USERS", "pet-app-users");
    this.envNodeRuntime = lambda.Runtime.NODEJS_10_X;
  }

  /**
   * @desc Publish the custom resource output
   */
  private outputStack() {
    new cdk.CfnOutput(this, "apiUrlOutput", {
      description: "API URL",
      value: this.api.url
    });
    new cdk.CfnOutput(this, "userPoolIdOutput", {
      description: "UserPool ID",
      value: this.userPool.id
    });
    new cdk.CfnOutput(this, "userPoolClientIdOutput", {
      description: "UserPool Client ID (App Client ID)",
      value: this.userPool.clientId
    });
    new cdk.CfnOutput(this, "regionOutput", {
      description: "Region",
      value: this.region
    });
    new cdk.CfnOutput(this, "userPoolDomainOutput", {
      description: "UserPool Domain (Cognito Domain)",
      value: this.userPool.domain
    });
    new cdk.CfnOutput(this, "petsLambdaFunctionNameOutput", {
      description: "Pets Lambda Function Name",
      value: this.petsLambda.functionName
    });
  }
}
