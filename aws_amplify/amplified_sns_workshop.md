---
title: Amplified SNS Workshop
tags: aws-amplify
url: http://educationhub-d90022b0-792f-11ea-b337-93f676b35e41.s3-website-us-east-1.amazonaws.com/
---

# 1. 準備
## 1.2. インストール
```
npm install -g @aws-amplify/cli@4.16.1
amplify configure
```

# 3. MVPを作ろう！
## 3.1. Bootstrap
```
mkdir amplified-sns-workshop
cd amplified-sns-workshop

npx create-react-app boyaki
cd boyaki

amplify init
# Enter a name for the project boyaki
# Enter a name for the environment production
# Choose your default editor: Vim (via Terminal, Mac OS only)
# Choose the type of app that you’re building javascript
# What javascript framework are you using react
# Source Directory Path: src
# Distribution Directory Path: build
# Build Command: npm run-script build
# Start Command: npm run-script start
# Do you want to use an AWS profile? Yes
# Please choose the profile you want to use amplify-handson

npm start
```

## 3.2. 認証機能の追加
```
amplify add auth
# Do you want to use the default authentication and security configuration? Default configuration
# How do you want users to be able to sign in? Username
# Do you want to configure advanced settings? No, I am done.

amplify status

amplify push
```

### 認証機能のフロントエンドへの実装
```
npm install --save aws-amplify@2.2.7 aws-amplify-react@3.1.8
```

```js
// ./src/App.js
import React from 'react';

import Amplify from '@aws-amplify/core';
import awsmobile from './aws-exports';

import { withAuthenticator } from 'aws-amplify-react';

Amplify.configure(awsmobile);

function App() {
  return (
    <h1>
      Hello World!
    </h1>
  );
}

export default withAuthenticator(App, {
  signUpConfig: {
    hiddenDefaults: ['phone_number']
  }
});
```

## 3.3. Post機能: Back-End(1)
```
amplify add api
# Please select from one of the below mentioned services: GraphQL
# Provide API name: BoyakiGql
# Choose the default authorization type for the API: Amazon Cognito User Pool
# Do you want to configure advanced settings for the GraphQL API: No, I am done.
# Do you have an annotated GraphQL schema? No
# Do you want a guided schema creation? No
# Provide a custom type name: Post
```

### Post型の作成
```
# ./amplify/backend/api/BoyakiGql/schema.graphql
type Post
  @model (subscriptions: { level: public })
  @auth(rules: [
    {allow: owner, ownerField:"owner", provider: userPools, operations:[read, create]}
    {allow: private, provider: userPools, operations:[read]}
  ])
{
  type: String! # always set to 'post'. used in the SortByTimestamp GSI
  id: ID
  content: String!
  owner: String
  timestamp: AWSTimestamp!
}
```

### Amplify Mocking
```
amplify mock api
# Choose the code generation language target javascript
# Enter the file name pattern of graphql queries, mutations and subscriptions src/graphql/**/*.js
# Do you want to generate/update all possible GraphQL operations - queries, mutations and subscriptions Yes
# Enter maximum statement depth [increase from default if your schema is deeply nested] 3
```

設定を変えたい場合は `amplify update codegen` コマンドにより同様の設定を行うことが可能です。

**createPost**
```mutation
mutation MyMutation {
  createPost(input: {type: "Post", content: "test", timestamp: "1586947321"}) {
    content
    id
    owner
    timestamp
    type
  }
}
### response
{
  "data": {
    "createPost": {
      "content": "test",
      "id": "28cdbcb1-e52e-42bd-901f-d81360073a14",
      "owner": "user1",
      "timestamp": "1586947321",
      "type": "Post"
    }
  }
}
```

**listPosts**
```query
query MyQuery {
  listPosts {
    items {
      content
      id
      owner
      timestamp
      type
    }
  }
}
### response
{
  "data": {
    "listPosts": {
      "items": [
        {
          "content": "test",
          "id": "e2335c1e-a306-4962-9ec4-83265a09e131",
          "owner": "user1",
          "timestamp": "1586947322",
          "type": "Post"
        },
        {
          "content": "test",
          "id": "28cdbcb1-e52e-42bd-901f-d81360073a14",
          "owner": "user1",
          "timestamp": "1586947321",
          "type": "Post"
        },
        {
          "content": "test",
          "id": "1fcc5566-bd4c-4b3c-aefa-0d1e1db760a0",
          "owner": "user1",
          "timestamp": "1586947323",
          "type": "Post"
        },
        {
          "content": "test",
          "id": "10272ca2-a10b-4600-8b56-86ebd20bd43c",
          "owner": "user1",
          "timestamp": "1586947325",
          "type": "Post"
        }
      ]
    }
  }
}
```

## 3.4. Post機能: Back-end(2)
### Amazon DynamoDBのPartition KeyとSort Key
DynamoDBのクエリは、最大で二つのAttribute(AppSyncのフィールド)を使うのがよいとされます。 この二つのAttributeはPartition Key(PK)とSort Key(SK)と呼ばれます。 DynamoDBはPK単体をプライマリキー(Tableにおけるユニークな識別子)、あるいはPKとSKを組み合わせてプライマリキーとして利用することができます。

### 必要なクエリを考えてみよう
Postには自動でIDを振りたいですし、Post単体をgetPostQueryでidを指定して引っ張って来るためにも、作成したDynamoDB TableのPartition Key(PK)はidフィールドのままで良いでしょう。

ここで、クライアントアプリケーションがどのような形式でデータをフェッチしたいか考えてみましょう。

- 全てのつぶやきを時系列順にリスト
- 特定のユーザーによるつぶやきをリスト

```graphql
# ./amplify/backend/api/BoyakiGql/schema.graphql
type Post
  @model (subscriptions: { level: public })
  @auth(rules: [
    {allow: owner, ownerField:"owner", provider: userPools, operations:[read, create]}
    {allow: private, provider: userPools, operations:[read]}
	])
  @key(name: "SortByTimestamp", fields:["type", "timestamp"], queryField: "listPostsSortedByTimestamp")
  @key(name: "BySpecificOwner", fields:["owner", "timestamp"], queryField: "listPostsBySpecificOwner")
{
  type: String! # always set to 'post'. used in the SortByTimestamp GSI
  id: ID
  content: String!
  owner: String
  timestamp: AWSTimestamp!
}
```

### Amplify Mockによる挙動の確認
```
rm amplify/mock-data/dynamodb/fake_us-fake-1.db
amplify mock api
```

```query
query MyQuery {
  listPostsBySpecificOwner(owner: "user_2", sortDirection: ASC) {
    items {
      owner
      timestamp
    }
  }
  listPostsSortedByTimestamp(type: "post", sortDirection: DESC) {
    items {
      owner
      timestamp
    }
  }
}
### response
{
  "data": {
    "listPostsBySpecificOwner": {
      "items": [
        {
          "owner": "user_2",
          "timestamp": 1586947327
        },
        {
          "owner": "user_2",
          "timestamp": 1586947328
        }
      ]
    },
    "listPostsSortedByTimestamp": {
      "items": [
        {
          "owner": "user_2",
          "timestamp": 1586947328
        },
        {
          "owner": "user_2",
          "timestamp": 1586947327
        },
        {
          "owner": "user1",
          "timestamp": 1586947324
        },
        {
          "owner": "user1",
          "timestamp": 1586947323
        },
        {
          "owner": "user1",
          "timestamp": 1586947321
        }
      ]
    }
  }
}
```

## 3.5. Post機能: Front-end
- Sidebar.js: 左側のメニュー一覧
- 右側のPost一覧
    - AllPost.js: Global Timeline - すべてのユーザーのPostが表示される
    - PostsBySpecifiedUser.js: Profile - 特定のユーザーのPostが表示される
- App.js: ルーティング

```
npm install --save @material-ui/core @material-ui/icons moment react-router react-router-dom
```

```
mkdir src/containers
touch src/containers/Sidebar.js
touch src/containers/AllPosts.js
touch src/containers/PostsBySpecifiedUser.js

mkdir src/components
touch src/components/PostList.js

tree 
├── App.css
├── App.js
├── App.test.js
├── aws-exports.js
├── components
│   └── PostList.js
├── containers
│   ├── AllPosts.js
│   ├── PostsBySpecifiedUser.js
│   └── Sidebar.js
├── graphql
│   ├── mutations.js
│   ├── queries.js
│   ├── schema.json
│   └── subscriptions.js
├── index.css
├── index.js
├── logo.svg
├── serviceWorker.js
└── setupTests.js
```

### 動作確認
```
amplify push
```

# 4. ウェブサイトホスティング
## 手動デプロイ
```
amplify add hosting
# Select the plugin module to execute Hosting with Amplify Console (Managed hosting with custom domains, Continuous deployment)
# Choose a type Manual deployment

amplify publish
```

# 5. Follow/Timeline機能の実装
## 5.1. Follow機能: Back-end
### GraphQL APIの作成
```graphql
// ./amplify/backend/api/BoyakiGql/schema.graphql
type FollowRelationship
	@model
	@auth(rules: [
		{allow: owner, ownerField:"followerId", provider: userPools, operations:[read, create]},
		{allow: private, provider: userPools, operations:[read]}
	])
	@key(fields: ["followeeId", "followerId"])
{
	followeeId: ID!
	followerId: ID!
	timestamp: AWSTimestamp!
}
```

### Mockingによる動作確認
```
amplify mock api
```
**mutation**
```
mutation MyMutation {
  createFollowRelationship(input: {followeeId: "test_followee", followerId: "test_follower", timestamp: 1587032763}) {
    followeeId
    followerId
    timestamp
  }
}
```

**query**
```
query MyQuery {
  listFollowRelationships(followeeId: "test_followee") {
    items {
      followeeId
      followerId
      timestamp
    }
  }
}
```

## 5.2. Follow機能: Front-end
### PostsBySpecifiedUser
なぜ$ amplify pushしていないにも関わらず、schema.graphqlの変更が行われた後の挙動をGraphQL APIが行っているのでしょうか。 Amplify CLIでは、 $ amplify pushを実行すると、Amplify Framework(SDK)がリソースにアクセスするために必要なエンドポイント、IDといった情報をaws-exports.jsにexportします。 aws-exports.jsを使ってAmplifyの初期設定を行うと、リソース変更や追加をしても設定を書き換える必要がなくなり、非常に楽です。 そして$ amplify mockコマンドを使用してAmplify Mockingをしている間だけ、ローカルでたちあがったMock ServerのGraphQL Endpointを指すようにaws-exports.jsが書き換わります。

## 5.3. Timeline機能: Back-end
### Timeline APIの作成
```
// ./amplify/backend/api/BoyakiGql/schema.graphql
type Timeline 
	@model
	@auth(rules: [
      {allow: owner, ownerField: "userId", provider: userPools, operations:[read, create]},
      {allow: private, provider: iam ,operations:[create]},
	])
	@key(fields: ["userId", "timestamp"])
{
	userId: ID!
	timestamp: AWSTimestamp!
	postId: ID!
	post: Post @connection(fields: ["postId"])
}
```

### GraphQL APIの認証方法にIAMを追加
```
amplify update api
# Please select from one of the below mentioned services: GraphQL
# Choose the default authorization type for the API Amazon Cognito User Pool
# Do you want to configure advanced settings for the GraphQL API Yes, I want to make some additional changes.
# Configure additional auth types? Yes
# Choose the additional authorization types you want to configure for the API IAM
# Configure conflict detection? No
```

### Amplify Mockingによる動作確認
```
amplify mock api
```


## 5.4. Timeline機能: @function
### amplify add function
```
amplify add function
# Provide a friendly name for your resource to be used as a label for this category in the project: createPostAndTimeline
# Provide the AWS Lambda function name: createPostAndTimeline
# Choose the function template that you want to use: Hello world function
# Do you want to access other resources created in this project from your Lambda function? Yes
# Select the category api
# Select the operations you want to permit for BoyakiGql create, read
# Do you want to edit the local lambda function now? No
```

### createPostAndTimeline Mutationの作成
```
type Mutation
{
  createPostAndTimeline(
		content: String!
	): Post
    @function(name: "createPostAndTimeline-${env}")
    @auth(rules: [
      {allow: private, provider: userPools},
    ])
}
```

### 既存APIへのアクセス権の追加
```
cd ./amplify/backend/function/createPostAndTimeline/src
npm install --save aws-appsync graphql-tag node-fetch
cd ../../../../..
```

## 5.5. Timeline機能: Front-end
```
amplify publish
```

# 7. 複数メンバーでの開発
## 7.1. Staging 環境の構築
```sh
amplify env list

amplify env add
# Do you want to use an existing environment?: No
# Enter a name for the environment: staging
# Do you want to use an AWS profile?: Yes
# Please choose the profile you want to use amplify-handson

amplify env list

amplify env checkout production
amplify env list

amplify env checkout staging
amplify push
```

## 7.2. Githubと連携してCI/CD環境を構築する(2)
### Amplify Console と Github アカウントを紐付ける
```
amplify env checkout production
amplify status
amplify remove hosting
amplify push

amplify env checkout staging
amplify status
amplify push

amplify env checkout production
amplify hosting add
# ? Select the plugin module to execute (Use arrow keys) では、
# > Hosting with Amplify Console (Managed hosting with custom domains, Continuous deployment)
```
