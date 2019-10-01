---
title: "Classify Text into Categories with the Natural Language API"
tags: google-cloud-platform, google-cloud-natural-language
url: https://www.qwiklabs.com/focuses/1749
---

# Goal
- Creating a Natural Language API request and calling the API with curl
- Using the NL API's text classification feature
- Using text classification to understand a dataset of news articles

# Task
- [x] Setup and Requirements
- [x] Confirm that the Cloud Natural Language API is enabled
- [x] Create an API Key
- [x] Classify a news article
- [x] Classifying a large text dataset
- [x] Creating a BigQuery table for our categorized text data
- [x] Classifying news data and storing the result in BigQuery
- [x] Analyzing categorized news data in BigQuery

# Supplement
![](classify_text_into_categories_with_the_natural_language_api.png)

```uml
skinparam monochrome true
skinparam backgroundColor #EEEEFF

actor User as U
actor "NLP" as N
participant Storage as S
participant BigQuery as B

U <-> S: |read|\ntext-classification-codelab
activate U
U -> N: |classify|
activate N
N -> U
deactivate N
U -> B: |insert_rows|\nnews_classification_dataset.article_data
```

## Create an API Key
```sh
export API_KEY=AIzaSyCTiFA2egBfi54ljKWkaUD2F9_67rklf0s
```

## Classify a news article
```sh
cat <<EOF >request.json
{
  "document":{
    "type":"PLAIN_TEXT",
    "content":"A Smoky Lobster Salad With a Tapa Twist. This spin on the Spanish pulpo a la gallega skips the octopus, but keeps the sea salt, olive oil, pimenton and boiled potatoes."
  }
}
EOF

curl "https://language.googleapis.com/v1/documents:classifyText?key=${API_KEY}" -s -X POST -H "Content-Type: application/json" --data-binary @request.json
```

## Classifying a large text dataset
```sh
gsutil cat gs://text-classification-codelab/bbc_dataset/entertainment/001.txt
```

## Classifying news data and storing the result in BigQuery
```sh
export PROJECT=$(gcloud config list project --format "value(core.project)")
gcloud iam service-accounts create my-account --display-name my-account
gcloud projects add-iam-policy-binding $PROJECT --member=serviceAccount:my-account@$PROJECT.iam.gserviceaccount.com --role=roles/bigquery.admin
gcloud iam service-accounts keys create key.json --iam-account=my-account@$PROJECT.iam.gserviceaccount.com
export GOOGLE_APPLICATION_CREDENTIALS=key.json

cat <<EOF >classify-text.py
from google.cloud import storage, language, bigquery

# Set up our GCS, NL, and BigQuery clients
storage_client = storage.Client()
nl_client = language.LanguageServiceClient()
# TODO: replace YOUR_PROJECT with your project id below
bq_client = bigquery.Client(project='qwiklabs-gcp-0f04638f9c7fb9ec')

dataset_ref = bq_client.dataset('news_classification_dataset')
dataset = bigquery.Dataset(dataset_ref)
table_ref = dataset.table('article_data') # Update this if you used a different table name
table = bq_client.get_table(table_ref)

# Send article text to the NL API's classifyText method
def classify_text(article):
        response = nl_client.classify_text(
                document=language.types.Document(
                        content=article,
                        type=language.enums.Document.Type.PLAIN_TEXT
                )
        )
        return response

rows_for_bq = []
files = storage_client.bucket('text-classification-codelab').list_blobs()
print("Got article files from GCS, sending them to the NL API (this will take ~2 minutes)...")

# Send files to the NL API and save the result to send to BigQuery
for file in files:
        if file.name.endswith('txt'):
                article_text = file.download_as_string()
                nl_response = classify_text(article_text)
                if len(nl_response.categories) > 0:
                        rows_for_bq.append((article_text, nl_response.categories[0].name, nl_response.categories[0].confidence))

print("Writing NL API article data to BigQuery...")
# Write article text + category data to BQ
errors = bq_client.insert_rows(table, rows_for_bq)
assert errors == []
EOF

python classify-text.py
```

```sql
SELECT * FROM `YOUR_PROJECT.news_classification_dataset.article_data`
```

## Analyzing categorized news data in BigQuery
```sql
SELECT
  category,
  COUNT(*) c
FROM
  `YOUR_PROJECT.news_classification_dataset.article_data`
GROUP BY
  category
ORDER BY
  c DESC
```

```sql
SELECT * FROM `YOUR_PROJECT.news_classification_dataset.article_data`
WHERE category = "/Arts & Entertainment/Music & Audio/Classical Music"
```

```sql
SELECT
  article_text,
  category
FROM `YOUR_PROJECT.news_classification_dataset.article_data`
WHERE cast(confidence as float64) > 0.9
```
