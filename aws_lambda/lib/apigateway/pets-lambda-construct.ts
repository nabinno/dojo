import fs = require("fs");
import cdk = require("@aws-cdk/core");
import apigw = require("@aws-cdk/aws-apigateway");
import lambda = require("@aws-cdk/aws-lambda");
import iam = require("@aws-cdk/aws-iam");
import { AwsLambdaApigatewayStack } from "./apigateway-stack";

export class PetsLambdaConstruct extends cdk.Construct {
  public readonly functionName: string;
  private readonly uri: string;
  private readonly lambdaAuthorizerId: string;
  private readonly cognitoAuthorizerId: string;

  constructor(scope: AwsLambdaApigatewayStack, id: string) {
    super(scope, id);

    let uid: string = "";
    if (fs.existsSync("tmp/cache/uid-pets")) {
      uid = fs
        .readFileSync("tmp/cache/uid-pets")
        .toString()
        .replace(/^.+?:(.+?)\n$/, "$1");
    }

    const petsFn = new lambda.Function(scope, "petsLambda", {
      functionName: `${scope.stackName}-Pets`,
      runtime: lambda.Runtime.GO_1_X,
      code: lambda.Code.fromBucket(scope.lambdaBucket, `pets/_build/pets-${uid}.zip`),
      handler: `_build/pets-${uid}`,
      memorySize: 128,
      timeout: cdk.Duration.seconds(30),
      environment: {
        REGION: scope.region,
        ITEMS_TABLE_NAME: scope.table.items.tableName,
        USERS_TABLE_NAME: scope.table.users.tableName,
        ALLOWED_ORIGIN: scope.envAppUrl,
        AUTHORIZATION_HEADER_NAME: scope.envAuthorizationHeaderName,
        USER_POOL_ID: scope.userPool.defaultChildId,
        ROLE_NAME_OF_ADMINS: scope.envRoleNameOfAdmins,
        ROLE_NAME_OF_USERS: scope.envRoleNameOfUsers
      }
    });
    // grant apigateway access to lambda
    petsFn.addPermission(`${scope.stackName}-Pets`, scope.apiLambdaPermission);
    // grant lambda access to tables
    scope.table.items.grantReadWriteData(petsFn.role!);
    scope.table.users.grantReadWriteData(petsFn.role!);
    // grant lambda access to cognito
    petsFn.addToRolePolicy(
      new iam.PolicyStatement({
        resources: [scope.userPool.arn],
        actions: ["cognito-idp:AdminUserGlobalSignOut", "cognito-idp:AdminGetUser"]
      })
    );

    this.functionName = petsFn.functionName;
    this.uri = `arn:aws:apigateway:${scope.region}:lambda:path/2015-03-31/functions/${petsFn.functionArn}/invocations`;
    this.lambdaAuthorizerId = scope.apiLambdaAuthorizer.id;
    this.cognitoAuthorizerId = scope.apiCognitoAuthorizer.id;
    const petsRc = scope.api.root.addResource("pets");
    const petRc = petsRc.addResource("{id}");
    const lambdaRc = scope.api.root.addResource("pets-lambda");
    const noneRc = scope.api.root.addResource("pets-none");

    // GET /pets
    // @note the path require the cognito authorizer (validates the JWT and passes it to the lambda)
    this.addMethodWithCognitoAuthorizer(petsRc, "GET");

    // GET /pets/:id
    this.addMethodWithCognitoAuthorizer(petRc, "GET");
    // this.addMethodWithCognitoAuthorizer(pet, "GET", { authorizationType: apigw.AuthorizationType.NONE });

    // POST /pets
    this.addMethodWithCognitoAuthorizer(petsRc, "POST", {
      requestValidator: scope.apiRequestValidator,
      requestModels: {
        "application/json": new apigw.Model(scope, "petModel", {
          restApi: scope.api,
          modelName: "petModel",
          contentType: "application/json",
          schema: {
            type: apigw.JsonSchemaType.OBJECT,
            properties: {
              breed: { type: apigw.JsonSchemaType.STRING },
              name: { type: apigw.JsonSchemaType.STRING },
              dateOfBirth: { type: apigw.JsonSchemaType.STRING }
            },
            required: ["breed", "name", "dateOfBirth"]
          }
        })
      }
    });

    // GET /lambda
    this.addMethodWithLambdaAuthorizer(lambdaRc, "GET");

    // GET /none
    this.addMethodWithoutAuthorizer(noneRc, "GET");
  }

  private addMethodWithCognitoAuthorizer(
    resource: apigw.Resource,
    method: string,
    methodOptions?: apigw.MethodOptions
  ) {
    resource.addMethod(
      method,
      new apigw.Integration({
        type: apigw.IntegrationType.AWS_PROXY,
        uri: this.uri,
        integrationHttpMethod: "POST"
      }),
      {
        authorizationType: apigw.AuthorizationType.COGNITO,
        authorizer: { authorizerId: this.cognitoAuthorizerId },
        ...methodOptions
      }
    );

    // // API Gateway now expects an access token instead of an ID token
    // (apigwMethod.node.defaultChild as apigw.CfnMethod).authorizationScopes = ["openid"];
  }

  private addMethodWithLambdaAuthorizer(
    resource: apigw.Resource,
    method: string,
    methodOptions?: apigw.MethodOptions
  ) {
    resource.addMethod(
      method,
      new apigw.Integration({
        type: apigw.IntegrationType.AWS_PROXY,
        uri: this.uri,
        integrationHttpMethod: "POST"
      }),
      {
        authorizationType: apigw.AuthorizationType.CUSTOM,
        authorizer: { authorizerId: this.lambdaAuthorizerId },
        ...methodOptions
      }
    );
  }

  private addMethodWithoutAuthorizer(
    resource: apigw.Resource,
    method: string,
    methodOptions?: apigw.MethodOptions
  ) {
    resource.addMethod(
      method,
      new apigw.Integration({
        type: apigw.IntegrationType.AWS_PROXY,
        uri: this.uri,
        integrationHttpMethod: "POST"
      }),
      {
        authorizationType: apigw.AuthorizationType.NONE,
        ...methodOptions
      }
    );
  }
}
