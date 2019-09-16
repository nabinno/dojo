---
title: "Ingesting New Datasets into BigQuery"
tags: google-cloud-platform, google-cloud-bigquery
url: https://www.qwiklabs.com/focuses/3692
---

# Goal


# Task
- [x] Create a new dataset to store tables
- [x] Ingest a new Dataset from a CSV
- [x] Ingest data from Google Cloud Storage
- [x] Ingest a new dataset from a Google Spreadsheet
- [x] Saving Data to Google Sheets
- [x] External table performance and data quality considerations

# Supplement
## CSV files
- gs://data-insights-course/exports/products.csv

## SQL example
```sql
#standardSQL
SELECT
  *
FROM
  ecommerce.products
ORDER BY
  stockLevel DESC
LIMIT  5
```

## Ingesting new dataset from Google Spreadsheets
**Source SQL**
```sql
#standardSQL
SELECT
  *,
  SAFE_DIVIDE(orderedQuantity,stockLevel) AS ratio
FROM
  ecommerce.products
WHERE
# include products that have been ordered and
# are 80% through their inventory
orderedQuantity > 0
AND SAFE_DIVIDE(orderedQuantity,stockLevel) >= .8
ORDER BY
  restockingLeadTime DESC
```

**Source Drive URL**
https://drive.google.com/open?id=<your spreadsheet id>

```sql
#standardSQL
SELECT * FROM ecommerce.products_comments WHERE comments IS NOT NULL
```

