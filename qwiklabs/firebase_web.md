---
title: Firebase Web
tags: firebase
url: https://www.qwiklabs.com/focuses/660?parent=catalog
---

# Goal
- Sync data using the Cloud Firestore and Cloud Storage for Firebase.
- Authenticate your users using Firebase Auth.
- Deploy your web app on Firebase Hosting.
- Send notifications with Firebase Cloud Messaging.

# Task
- [x] Overview
- [x] Setup and requirements
- [x] Get the sample code
- [x] View the starter application
- [x] Set up your Firebase project
- [x] Enable Firebase on your project
- [x] Add a Firebase web app
- [x] Install the Firebase command line interface
- [x] Deploy and run the starter app
- [x] Import and Configure Firebase
- [x] Set up user sign in
- [x] Write messages to Cloud Firestore
- [x] Read messages
- [x] Send Images
- [x] Show Notifications
- [x] Cloud Firestore security rules (optional)
- [x] Cloud Storage security rules (optional)
- [x] Deploy your app using Firebase Hosting
- [x] Congratulations!
- [x] Finish Your Quest

# Supplement
## Get the sample code
```sh
git clone https://github.com/firebase/friendlychat-web
```

## Set up your Firebase project
- Firebase Authentication: to easily let your users sign-in your app.
- Cloud Firestore: to save structured data on the cloud and get instant notification when data changes.
- Cloud Storage: for Firebase to save files in the cloud.
- Firebase Hosting: to host and serve your assets.
- Firebase Cloud Messaging: to send push notifications and display browser popup notifications.

## Install the Firebase command line interface
```sh
firebase --version
firebase login --no-localhost

cd ~/friendlychat-web/web-start/
firebase use --add
```

## Deploy and run the starter app
```sh
firebase serve --only hosting
```

## Set up user sign in
```sh
firebase deploy --except functions
```

## Show Notifications
```sh
YOUR_SERVER_KEY=AAAAjExXgh8:APA91bEWlLP9BaGkzrdOfa1asQBWnWSWu9Ddl1L0k25LK8Y-jJXW1uPXalFtjowycXUvLxf06CGS3b6ZMh2FWIAJsa5tVEiDnHWWGhoyHGjsM1ECuCfEvBsw5eFwRlkSF3Hta9CFnDut

curl -H "Content-Type: application/json" \
     -H "Authorization: key=${YOUR_SERVER_KEY}" \
     -d '{
           "notification": {
             "title": "New chat message!",
             "body": "There is a new message in FriendlyChat",
             "icon": "/images/profile_placeholder.png",
             "click_action": "http://localhost:5000"
           },
           "to": "eoR1iw1qbjbyj99givpG93:APA91bHBJI-P4RckknMfTuyH8NK_Qgq-a8GSYB7cjYtRUS4DTxRc3mh7Zldf__lZ7bVB-Djj4_fAc7l1X4R5MqsjrW-mZY-6i8ziy4LWAtERCMlgwHoEql3UK45DkWdpKJGkFo1EXbtX"
         }' \
     https://fcm.googleapis.com/fcm/send
```

## Cloud Firestore security rules (optional)
```
service cloud.firestore {
  match /databases/{database}/documents {
    // Messages:
    //   - Anyone can read.
    //   - Authenticated users can add and edit messages.
    //   - Validation: Check name is same as auth token and text length below 300 char or that imageUrl is a URL.
    //   - Deletes are not allowed.
    match /messages/{messageId} {
      allow read;
      allow create, update: if request.auth != null
                    && request.resource.data.name == request.auth.token.name
                    && (request.resource.data.text is string
                      && request.resource.data.text.size() <= 300
                      || request.resource.data.imageUrl is string
                      && request.resource.data.imageUrl.matches('https?://.*'));
      allow delete: if false;
    }
    // FCM Tokens:
    //   - Anyone can write their token.
    //   - Reading list of tokens is not allowed.
    match /fcmTokens/{token} {
      allow read: if false;
      allow write;
    }
  }
}
```

## Cloud Storage security rules (optional)
```
// Returns true if the uploaded file is an image and its size is below the given number of MB.
function isImageBelowMaxSize(maxSizeMB) {
  return request.resource.size < maxSizeMB * 1024 * 1024
      && request.resource.contentType.matches('image/.*');
}

service firebase.storage {
  match /b/{bucket}/o {
    match /{userId}/{messageId}/{fileName} {
      allow write: if request.auth != null && request.auth.uid == userId && isImageBelowMaxSize(5);
      allow read;
    }
  }
}
```

# Refernces
- https://firebase.google.com/
- https://firebase.google.com/support/faq/?authuser=1#pricing
- https://firebase.google.com/docs/projects/iam/overview?authuser=1
- https://firebase.google.com/docs/firestore/
- https://firebase.google.com/docs/storage/
- https://firebase.google.com/docs/cloud-messaging/
- https://developer.mozilla.org/en/docs/Web/API/Service_Worker_API
- https://firebase.google.com/docs/firestore/security/get-started
- https://firebase.google.com/docs/database/security/
- https://firebase.google.com/docs/storage/security/start
- https://firebase.google.com/docs/hosting/
