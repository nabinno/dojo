---
title: "Creating a Data Transformation Pipeline with Cloud Dataprep"
tags: google-cloud-platform, google-cloud-dataprep, data-wrangling, data-pre-processing, google-cloud-bigquery, google-analytics
url: https://www.qwiklabs.com/focuses/4415
---

# Goal
- Connect BigQuery datasets to Cloud Dataprep
- Explore dataset quality with Cloud Dataprep
- Create a data transformation pipeline with Cloud Dataprep
- Schedule transformation jobs outputs to BigQuery

# Task
- [x] Setting up your development environment
- [x] Creating a BigQuery Dataset
- [x] Opening Cloud Dataprep
- [x] Connecting BigQuery data to Cloud Dataprep
- [x] Exploring ecommerce data fields with a UI
- [x] Cleaning the data
- [x] Enriching the data
- [x] Running and scheduling Cloud Dataprep jobs to BigQuery

# Supplement
## Creating a BigQuery Dataset
```sql
#standardSQL
 CREATE OR REPLACE TABLE ecommerce.all_sessions_raw_dataprep
 OPTIONS(
   description="Raw data from analyst team to ingest into Cloud Dataprep"
 ) AS
 SELECT * FROM `next-marketing-analytics.ecommerce.all_sessions_raw`
 WHERE date = '20170801'; # limiting to one day of data 56k rows for this lab
```

## Cleaning the data and enriching the data
- Change productSKU type to String
- Delete itemQuantity
- Delete itemRevenue
- Remove duplicate rows
- Delete rows where ISMISSING([totalTransactionRevenue])
- Keep rows where type == 'PAGE'
- Concatenate fullVisitorId, visitId separated by '-'
- Create eCommerceAction_label from 9 case conditions on eCommerceAction_type
- Create totalTransactionRevenue1 from DIVIDE(totalTransactionRevenue, 1000000)
- Change totalTransactionRevenue1 type to Decimal

## Reference
- https://support.google.com/analytics/answer/3437719
