---
title: "Detect Labels, Faces, and Landmarks in Images with the Cloud Vision API"
tags: google-cloud-platform, google-cloud-vision
url: https://www.qwiklabs.com/focuses/1841
---

# Goal
- Creating a Vision API request and calling the API with curl
- Using the label, face, and landmark detection methods of the vision API

# Task
- [ ] Create an API Key
- [ ] Upload an Image to a Cloud Storage bucket
- [ ] Create your Vision API request
- [ ] Label Detection
- [ ] Web Detection
- [ ] Face and Landmark Detection
- [ ] Calling the Vision API and parsing the response
- [ ] Explore other Vision API methods

# Supplement
## Create an API Key
```sh
export API_KEY=<YOUR_API_KEY>
```

## Create your Vision API request
```sh
nano request.json
nano> {
  "requests": [
      {
        "image": {
          "source": {
              "gcsImageUri": "gs://my-bucket-name/donuts.png"
          }
        },
        "features": [
          {
            "type": "LABEL_DETECTION",
            "maxResults": 10
          }
        ]
      }
  ]
}
```

## Label Detection
```sh
curl -s -X POST -H "Content-Type: application/json" --data-binary @request.json  https://vision.googleapis.com/v1/images:annotate?key=${API_KEY}
```

## Web Detection
```sh
nano request.json
nano> {
  "requests": [
      {
        "image": {
          "source": {
              "gcsImageUri": "gs://my-bucket-name/donuts.png"
          }
        },
        "features": [
          {
            "type": "WEB_DETECTION",
            "maxResults": 10
          }
        ]
      }
  ]
}

curl -s -X POST -H "Content-Type: application/json" --data-binary @request.json  https://vision.googleapis.com/v1/images:annotate?key=${API_KEY}
```

## Face and Landmark Detection
```sh
nano request.jso
nano> {
  "requests": [
      {
        "image": {
          "source": {
              "gcsImageUri": "gs://my-bucket-name/selfie.png"
          }
        },
        "features": [
          {
            "type": "FACE_DETECTION"
          },
          {
            "type": "LANDMARK_DETECTION"
          }
        ]
      }
  ]
}
```

## Calling the Vision API and parsing the response
```sh
curl -s -X POST -H "Content-Type: application/json" --data-binary @request.json  https://vision.googleapis.com/v1/images:annotate?key=${API_KEY}
```
