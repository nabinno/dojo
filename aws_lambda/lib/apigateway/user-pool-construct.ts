import cdk = require("@aws-cdk/core");
import cognito = require("@aws-cdk/aws-cognito");
import lambda = require("@aws-cdk/aws-lambda");
import { AwsLambdaApigatewayStack } from "./apigateway-stack";
import { Utils } from "../shared/utils";

export class UserPoolConstruct extends cdk.Construct {
  public arn: string;
  public id: string;
  public defaultChildId: string;
  public clientId: string;
  public domain: string;

  constructor(scope: AwsLambdaApigatewayStack, id: string) {
    super(scope, id);

    const supportedIdPs = ["COGNITO"];

    this.setUserPool(scope);

    this.setUserPoolClient(
      scope,
      this.setUserPoolIdentityProvider(scope, supportedIdPs),
      supportedIdPs
    );

    this.setUserPoolDomain(scope);
  }

  /**
   * User Pool
   * @desc Purpose: creates a user directory and allows federation from external IdPs
   * @see https://docs.aws.amazon.com/cdk/api/latest/docs/@aws-cdk_aws-cognito.CfnIdentityPool.html
   */
  private setUserPool(scope: AwsLambdaApigatewayStack) {
    const userPool = new cognito.UserPool(scope, `userPool`, {
      userPoolName: `${scope.stackName}-UserPool`,
      signInType: cognito.SignInType.EMAIL,
      autoVerifiedAttributes: [cognito.UserPoolAttribute.EMAIL],
      lambdaTriggers: {
        // @note
        //   Purpose: map from a custom attribute mapped from SAML, e.g. {..., "custom:groups":"[a,b,c]", ...}
        //            to cognito:groups claim, e.g. {..., "cognito:groups":["a","b","c"], ...}
        //            it can also optionally add roles and preferred_role claims
        // @see https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-lambda-pre-token-generation.html
        preTokenGeneration: new lambda.Function(scope, "preTokenGenerationLambda", {
          functionName: `${scope.stackName}-PreTokenGeneration`,
          runtime: scope.envNodeRuntime,
          handler: "index.handler",
          code: lambda.Code.fromAsset("lambda/pretokengeneration/dist/src"),
          environment: {
            GROUPS_ATTRIBUTE_CLAIM_NAME: scope.envGroupsAttributeClaimName
          }
        })
      }
    });
    this.arn = userPool.userPoolArn;
    this.id = userPool.userPoolId;

    // any properties that are not part of the high level construct can be added using this method
    const defaultChild = userPool.node.defaultChild as cognito.CfnUserPool;
    defaultChild.schema = [
      {
        name: scope.envGroupsAttributeName,
        attributeDataType: "String",
        mutable: true,
        required: false,
        stringAttributeConstraints: {
          maxLength: "2000"
        }
      }
    ];
    this.defaultChildId = defaultChild.ref;
  }

  /**
   * Identity Provider Settings
   * @note Purpose: define the external Identity Provider details, field mappings etc.
   * @see https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-saml-idp.html
   * mapping from IdP fields to Cognito attributes (key is cognito attribute, value is mapped field name)
   */
  private setUserPoolIdentityProvider(
    scope: AwsLambdaApigatewayStack,
    supportedIdPs: string[]
  ): cognito.CfnUserPoolIdentityProvider | undefined {
    let rc: cognito.CfnUserPoolIdentityProvider | undefined = undefined;

    if (scope.envIdentityProviderMetadataURLOrFile && scope.envIdentityProviderName) {
      rc = new cognito.CfnUserPoolIdentityProvider(scope, "CognitoIdP", {
        providerName: scope.envIdentityProviderName,
        providerDetails: Utils.isURL(scope.envIdentityProviderMetadataURLOrFile)
          ? {
              MetadataURL: scope.envIdentityProviderMetadataURLOrFile
            }
          : {
              MetadataFile: scope.envIdentityProviderMetadataURLOrFile
            },
        providerType: "SAML",
        attributeMapping: {
          email: "email",
          family_name: "lastName",
          given_name: "firstName",
          name: "firstName", // alias to given_name
          [scope.envGroupsAttributeClaimName]: "groups" //syntax for a dynamic key
        },
        userPoolId: this.id
      });

      supportedIdPs.push(scope.envIdentityProviderName);
    }

    return rc;
  }

  /**
   * User Pool Client - Cognito App Client
   * @note Purpose: each app needs an app client defined, where app specific details are set, such as redirect URIs
   * @see https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-client-apps.html
   */
  private setUserPoolClient(
    scope: AwsLambdaApigatewayStack,
    idP: cognito.CfnUserPoolIdentityProvider | undefined,
    supportedIdPs: string[]
  ) {
    const client = new cognito.CfnUserPoolClient(scope, "CognitoAppClient", {
      supportedIdentityProviders: supportedIdPs,
      clientName: "Web",
      allowedOAuthFlowsUserPoolClient: true,
      allowedOAuthFlows: ["code"],
      allowedOAuthScopes: ["phone", "email", "openid", "profile"],
      generateSecret: false,
      refreshTokenValidity: 1,
      callbackUrLs: [scope.envCallbackURL], // @todo add your app's prod URLs here
      logoutUrLs: [scope.envCallbackURL],
      userPoolId: this.id
    });

    // we want to make sure we do things in the right order
    if (idP) {
      client.node.addDependency(idP);
    }

    this.clientId = client.ref;
  }

  /**
   * User Pool Domain - Cognito Auth Domain
   * @note Purpose: creates / updates the custom subdomain for cognito's hosted UI
   * @see https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-assign-domain.html
   */
  private setUserPoolDomain(scope: AwsLambdaApigatewayStack) {
    const domain = new cognito.CfnUserPoolDomain(scope, "CognitoDomain", {
      domain: scope.envDomain,
      userPoolId: this.id
    });

    this.domain = domain.domain;
  }
}
