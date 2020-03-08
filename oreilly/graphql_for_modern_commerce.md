---
title: GraphQL for Modern Commerce
author: Kelly Goetsch
tags: graphql
url: https://learning.oreilly.com/library/view/graphql-for-modern/9781492056850/
---

# 1. Introducing GraphQL
## Commerce Requires More Than REST
**REST APIの課題**

REST APIは優れていますが、すべての変換テクノロジーと同様に、新たな課題がなければ利益は得られません。
- 仕様の欠如
- 呼び出すAPIがわからない
- データのオーバーフェッチ
- データのアンダーフェッチ

## What Is GraphQL?
```
query {
	product(id: "94695736", locale: "en_US") {
		brandIconURL
		name
		description
		price
		ableToSell
		averageReview
		numberOfReviews
		maxAllowableQty
		images {
			url
			altText
		}
	}
}
```

**GraphQLの利点**
- より簡単なフロントエンド開発
- パフォーマンスの向上
- 維持するコードが少ない

**GraphQLの欠点**
- セキュリティ
- 複数のGraphQLエンドポイントとスキーマを組み合わせることが難しい

## GraphQL Compared to REST APIs
![](https://learning.oreilly.com/library/view/graphql-for-modern/9781492056850/assets/gqmc_0107.png)

# 2. The GraphQL Specification
## Introducing the GraphQL Specification
## GraphQL Specification Governance and History
## Principles of the GraphQL Specification
## GraphQL Terminology
## GraphQL Operations
## Final Thoughts

# 3. GraphQL Clients
## Low-Level Networking
## Batching
## Authentication
## Caching
## Language-Specific Bindings
## Frontend Framework Integration
## Final Thoughts

# 4. GraphQL Servers
## Building a Type Schema
## HTTP Request Handling
## Parsing Queries
## Validating
## Executing Queries
## Server Implementations
## Monitoring
## Testing
## Security
## Merging Schemas
## Final Thoughts
