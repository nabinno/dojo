---
title: "Exploring Your Ecommerce Dataset with SQL in Google BigQuery"
tags: google-cloud-platform, google-cloud-bigquery
url: https://www.qwiklabs.com/focuses/3618
---

# Goal
- Access an ecommerce dataset
- Look at the dataset metadata
- Remove duplicate entries
- Write and execute queries

# Task
- [x] Pin the Lab Project in BigQuery
- [x] Explore ecommerce data and identify duplicate records
- [x] Write basic SQL on ecommerce data

# Supplement
## Google Analytics dataset of Google Merchandise Store
- https://console.cloud.google.com/bigquery?p=data-to-insights&page=ecommerce
- [BigQuery]-[data-to-insights]-[ecommerce]-[all_sessions_raw]

**Select duplicate rows**
```sql
#standardSQL
SELECT COUNT(*) as num_duplicate_rows, *
FROM `data-to-insights.ecommerce.all_sessions_raw`
GROUP BY
fullVisitorId, channelGrouping, time, country, city, totalTransactionRevenue, transactions, timeOnSite, pageviews, sessionQualityDim, date, visitId, type, productRefundAmount, productQuantity, productPrice, productRevenue, productSKU, v2ProductName, v2ProductCategory, productVariant, currencyCode, itemQuantity, itemRevenue, transactionRevenue, transactionId, pageTitle, searchKeyword, pagePathLevel1, eCommerceAction_type, eCommerceAction_step, eCommerceAction_option
HAVING num_duplicate_rows > 1;
```

**Deduplicate rows in SQL**
- https://stackoverflow.com/questions/36675521/delete-duplicate-rows-from-a-bigquery-table
- https://bigquery.cloud.google.com/savedquery/133415875420:c00d759e2bf54e4dba116c83568ac2cb

```sql
#standardSQL
CREATE OR REPLACE TABLE `data-to-insights.ecommerce.all_sessions` AS
# Deduplicate rows in SQL
# https://stackoverflow.com/questions/36675521/delete-duplicate-rows-from-a-bigquery-table
SELECT k.*
FROM (
  SELECT ARRAY_AGG(x LIMIT 1)[OFFSET(0)] k
  FROM `data-to-insights.ecommerce.all_sessions_raw` x
  GROUP BY
  fullVisitorId,
visitId,
date,
time,
v2ProductName,
productSKU,
type,
eCommerceAction_type,
eCommerceAction_step,
eCommerceAction_option,
  transactionRevenue,
  transactionId
)
```

**Select**

```sql
#standardSQL
# schema: https://support.google.com/analytics/answer/3437719?hl=en
SELECT
fullVisitorId, # the unique visitor ID
visitId, # a visitor can have multiple visits
date, # session date stored as string YYYYMMDD
time, # time of the individual site hit  (can be 0 to many per visitor session)
v2ProductName, # not unique since a product can have variants like Color
productSKU, # unique for each product
type, # a visitor can visit Pages and/or can trigger Events (even at the same time)
eCommerceAction_type, # maps to eadd to cart', ecompleted checkout'
eCommerceAction_step,
eCommerceAction_option,
  transactionRevenue, # revenue of the order
  transactionId, # unique identifier for revenue bearing transaction
COUNT(*) as row_count
FROM
`data-to-insights.ecommerce.all_sessions`
GROUP BY 1,2,3 ,4, 5, 6, 7, 8, 9, 10,11,12
HAVING row_count > 1 # find duplicates
```

**Select unique visitors**
```sql
#standardSQL
SELECT
  COUNT(*) AS product_views,
  COUNT(DISTINCT fullVisitorId) AS unique_visitors
FROM `data-to-insights.ecommerce.all_sessions`;
```

```sql
#standardSQL
SELECT
  COUNT(DISTINCT fullVisitorId) AS unique_visitors,
  channelGrouping
FROM `data-to-insights.ecommerce.all_sessions`
GROUP BY 2
ORDER BY 2 DESC;
```
**Distict product names**
```sql
#standardSQL
SELECT
  (v2ProductName) AS ProductName
FROM `data-to-insights.ecommerce.all_sessions`
GROUP BY 1
ORDER BY 1
```

**Top 5 of product view**
```sql
#standardSQL
SELECT
  COUNT(*) AS product_views,
  (v2ProductName) AS ProductName
FROM `data-to-insights.ecommerce.all_sessions`
WHERE type = 'PAGE'
GROUP BY v2ProductName
ORDER BY product_views DESC
LIMIT 5;
```

```sql
#standardSQL
SELECT
  COUNT(*) AS product_views,
  COUNT(productQuantity) AS orders,
  SUM(productQuantity) AS quantity_product_ordered,
  v2ProductName
FROM `data-to-insights.ecommerce.all_sessions`
WHERE type = 'PAGE'
GROUP BY v2ProductName
ORDER BY product_views DESC
LIMIT 5;
```

```sql
#standardSQL
SELECT
  COUNT(*) AS product_views,
  COUNT(productQuantity) AS orders,
  SUM(productQuantity) AS quantity_product_ordered,
  SUM(productQuantity) / COUNT(productQuantity) AS avg_per_order,
  (v2ProductName) AS ProductName
FROM `data-to-insights.ecommerce.all_sessions`
WHERE type = 'PAGE'
GROUP BY v2ProductName
ORDER BY product_views DESC
LIMIT 5;
```
