---
title: "Entity and Sentiment Analysis with the Natural Language API"
tags: google-cloud-platform, google-cloud-natural-language
url: https://www.qwiklabs.com/focuses/1843
---

# Goal
- Creating a Natural Language API request and calling the API with curl
- Extracting entities and running sentiment analysis on text with the Natural Language API
- Performing linguistic analysis on text with the Natural Language API
- Creating a Natural Language API request in a different language

# Task
- [x] Setup and Requirements
- [x] Create an API Key
- [x] Make an Entity Analysis Request
- [x] Call the Natural Language API
- [x] Sentiment analysis with the Natural Language API
- [x] Analyzing entity sentiment
- [x] Analyzing syntax and parts of speech
- [x] Multilingual natural language processing

# Supplement
## Create an API Key
```sh
export API_KEY=AIzaSyAn2pjtkOpZiWWY4elB2Z5zUpbtukIVr5E
```

## Make an Entity Analysis Request
```sh
cat <<EOF >request.json
{
  "document":{
    "type":"PLAIN_TEXT",
    "content":"Joanne Rowling, who writes under the pen names J. K. Rowling and Robert Galbraith, is a British novelist and screenwriter who wrote the Harry Potter fantasy series."
  },
  "encodingType":"UTF8"
}
EOF
```

## Call the Natural Language API
```sh
curl "https://language.googleapis.com/v1/documents:analyzeEntities?key=${API_KEY}" -s -X POST -H "Content-Type: application/json" --data-binary @request.json > result.json
cat result.json
```

## Sentiment analysis with the Natural Language API
```sh
cat <<EOF >request.json
 {
  "document":{
    "type":"PLAIN_TEXT",
    "content":"Harry Potter is the best book. I think everyone should read it."
  },
  "encodingType": "UTF8"
}
EOF

curl "https://language.googleapis.com/v1/documents:analyzeSentiment?key=${API_KEY}" -s -X POST -H "Content-Type: application/json" --data-binary @request.json
```

## Analyzing entity sentiment
```sh
cat <<EOF >request.json
 {
  "document":{
    "type":"PLAIN_TEXT",
    "content":"I liked the sushi but the service was terrible."
  },
  "encodingType": "UTF8"
}
EOF

curl "https://language.googleapis.com/v1/documents:analyzeEntitySentiment?key=${API_KEY}" -s -X POST -H "Content-Type: application/json" --data-binary @request.json
```

## Analyzing syntax and parts of speech
```sh
cat <<EOF >request.json
{
  "document":{
    "type":"PLAIN_TEXT",
    "content": "Joanne Rowling is a British novelist, screenwriter and film producer."
  },
  "encodingType": "UTF8"
}
EOF

curl "https://language.googleapis.com/v1/documents:analyzeSyntax?key=${API_KEY}" -s -X POST -H "Content-Type: application/json" --data-binary @request.json
```

## Multilingual natural language processing
```sh
cat <<EOF >request.json
{
  "document":{
    "type":"PLAIN_TEXT",
    "content":"日本のグーグルのオフィスは、東京の六本木ヒルズにあります"
  }
}

curl "https://language.googleapis.com/v1/documents:analyzeEntities?key=${API_KEY}" -s -X POST -H "Content-Type: application/json" --data-binary @request.json
```
