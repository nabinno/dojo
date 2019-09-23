---
title: "Extract, Analyze, and Translate Text from Images with the Cloud ML APIs"
tags: google-cloud-platform, google-cloud-vision, google-cloud-natural-language, web-api, optical-character-recognition, analytics
url: https://www.qwiklabs.com/focuses/1836
---

# Goal
- Creating a Vision API request and calling the API with curl
- Using the text detection (OCR) method of the Vision API
- Using the Translation API to translate text from your image
- Using the Natural Language API to analyze the text

# Task
- [x] Setup and Requirements
- [x] Create an API Key
- [x] Upload an image to a cloud storage bucket
- [x] Create your Vision API request
- [x] Call the Vision API's text detection method
- [x] Sending text from the image to the Translation API
- [x] Analyzing the image's text with the Natural Language API

# Supplement
## Create an API Key
```sh
export API_KEY=AIzaSyD4LscjchT1CiF8cMf-x2eYf0ZZtCG_nuA
```

## Create your Vision API request
```sh
nano ocr-request.json
nano> {
        "requests": [
            {
              "image": {
                "source": {
                    "gcsImageUri": "gs://my-image-bucket-2019-09-24/sign.jpg"
                }
              },
              "features": [
                {
                  "type": "TEXT_DETECTION",
                  "maxResults": 10
                }
              ]
            }
        ]
      }

curl -s -X POST -H "Content-Type: application/json" --data-binary @ocr-request.json  https://vision.googleapis.com/v1/images:annotate?key=${API_KEY}
curl -s -X POST -H "Content-Type: application/json" --data-binary @ocr-request.json  https://vision.googleapis.com/v1/images:annotate?key=${API_KEY} -o ocr-response.json

## Sending text from the image to the Translation API
```sh
nano translation-request.json
nano> {
        "q": "your_text_here",
        "target": "en"
      }

STR=$(jq .responses[0].textAnnotations[0].description ocr-response.json) && STR="${STR//\"}" && sed -i "s|your_text_here|$STR|g" translation-request.json
curl -s -X POST -H "Content-Type: application/json" --data-binary @translation-request.json https://translation.googleapis.com/language/translate/v2?key=${API_KEY} -o translation-response.json
cat translation-response.json

```

## Analyzing the image's text with the Natural Language API
```sh
nano nl-request.json
nano> {
        "document":{
          "type":"PLAIN_TEXT",
          "content":"your_text_here"
        },
        "encodingType":"UTF8"
      }

STR=$(jq .data.translations[0].translatedText  translation-response.json) && STR="${STR//\"}" && sed -i "s|your_text_here|$STR|g" nl-request.json
curl "https://language.googleapis.com/v1/documents:analyzeEntities?key=${API_KEY}" -s -X POST -H "Content-Type: application/json" --data-binary @nl-request.json
```
