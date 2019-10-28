# AWS Lambda
## Dependencies installation
```sh
npm clean-install
go mod why
npm run install:sam
```

## Local invocation
```sh
npm run invoke -- {{lambda_resource}} {{lambda_action}}
``

## Deployment commands
```sh
npm run deploy
```

## Useful commands for CDK
- `npm run build` compile typescript to js
- `npm run watch` watch for changes and compile
- `npm run test` perform the jest unit tests
- `cdk deploy` deploy this stack to your default AWS account/region
- `cdk diff` compare deployed stack with current state
- `cdk synth` emits the synthesized CloudFormation template

## Useful commands for Go
- `mage lbd:dispatch` dispatch on local
- `mage lbd:test` or `ginkgo` testing
- `mage lbd:deploy` deploy
- `mage lbd:clean` clean
