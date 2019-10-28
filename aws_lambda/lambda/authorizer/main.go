// @see
//   - https://github.com/awslabs/aws-apigateway-lambda-authorizer-blueprints/blob/master/blueprints/go/main.go
//   - https://github.com/aws/aws-lambda-go/blob/master/events/README_ApiGatewayCustomAuthorizer.md
package main

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, event events.APIGatewayCustomAuthorizerRequestTypeRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	log.Println("Client token: " + event.Headers["authorizationtoken"])
	log.Println("Method ARN: " + event.MethodArn)
	log.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaac")

	// @todo 2019-10-24
	//   - 着信トークンを検証しトークンに関連付けられたプリンシパルユーザー識別子を下記のいずれかの方法で生成します
	//     1. OAuthプロバイダーを呼び出す
	//     2. JWTトークンをインラインでデコードします
	//     3. 自己管理DBのルックアップ
	//   - トークンの状態におけるながれ
	//     - トークンが有効な場合
	//       - クライアントへのアクセスを許可または拒否するポリシーを生成する必要があります
	//         - アクセスが許可され、API Gatewayは呼び出されたメソッドで設定されたバックエンド統合を続行します
	//       - 失敗すると、401 Unauthorized応答をクライアントに送信できます。
	//         events.APIGatewayCustomAuthorizerResponse {}、errors.New（ "Unauthorized"）を返します
	//     - トークンが無効な場合
	//       - アクセスが拒否され、クライアントは403 Access Denied応答を受け取ります
	//   - 備考
	//     - この関数は、認識されたプリンシパルユーザー識別子に関連付けられたポリシーを生成する必要があります。ユー
	//       スケースに応じて、ポリシーをDBに保存するか、その場で生成します
	//     - ポリシーはデフォルトで5分間キャッシュされることに注意してください（承認者でTTLを構成可能）。そして、同
	//       じトークンで作成されたRestApi内の任意のメソッド/リソースへの後続の呼び出しに適用されます。
	principalID := "user|a1b2c3d4"

	methodArnTmp := strings.Split(event.MethodArn, ":")
	apiGatewayArnTmp := strings.Split(methodArnTmp[5], "/")

	resp := NewAuthorizerResponse(
		principalID,     // principalID
		methodArnTmp[4], // awsAccountID
	)
	resp.Region = methodArnTmp[3]
	resp.APIID = apiGatewayArnTmp[0]
	resp.Stage = apiGatewayArnTmp[1]
	// resp.DenyAllMethods()
	resp.AllowMethod(All, "/pets/*")

	// @note
	//   Add additional key-value pairs associated with the authenticated principal these are made available by APIGW
	//   like so: $context.authorizer.<key> additional context is cached
	resp.Context = map[string]interface{}{
		"stringKey":  "stringval",
		"numberKey":  123,
		"booleanKey": true,
	}

	return resp.APIGatewayCustomAuthorizerResponse, nil
}

func main() {
	lambda.Start(handleRequest)
}
