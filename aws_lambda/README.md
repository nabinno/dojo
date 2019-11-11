# AWS Lambda
## Dependencies installation
```sh
npm clean-install
go mod build
npm run install:sam
```

## Local invocation
```sh
npm run invoke -- {{lambda_resource}} {{lambda_action}}
```

## Deployment commands
```sh
npm run deploy
```

## Operation check
```sh
npm run cognito:curl -- \
  -v \
  -X GET \
  https://foo.execute-api.ap-northeast-1.amazonaws.com/prod/pets
```

## Useful commands for CDK
- `npm run build` compile typescript to js
- `npm run watch` watch for changes and compile
- `npm run test` perform the jest unit tests
- `cdk deploy` deploy this stack to your default AWS account/region
- `cdk diff` compare deployed stack with current state
- `cdk synth` emits the synthesized CloudFormation template
