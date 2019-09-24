---
title: "Implementing an AI Chatbot with Dialogflow"
tags: google-cloud-platform, chatbot, dialogflow
url: https://www.qwiklabs.com/focuses/634
---

# Goal
- Basics concepts and constructs of Dialogflow, including intent, entity and context
- Chatbot workflow
- Life of a conversation

# Task
- [x] Dialogflow Concepts and Constructs
- [x] Deploy a simple Dialogflow application to submit helpdesk tickets
- [x] Create Intents
- [x] Allow Fulfillment to Store Help Ticket Data
- [x] Verify that Tickets are Logged in Datastore
- [x] Testing your Chatbot
- [x] Test your Understanding

# Supplement
![](implementing_an_ai_chatbot_with_dialogflow.png)

```uml
skinparam monochrome true
skinparam backgroundColor #EEEEFF

actor Speaker as S
actor Home as H
box Dialogflow
  actor Intent as D
  actor API as A
  actor Fulfilment as F
  participant Datastore as DS
end box

activate S
S -> H: Ok Google, \ntalk to Helpdesk Support
activate H
H -> D
activate D
D <-> A: |Request|\nSTT, NLP, Knowledge Graph, \nML Ranking, User Profile
activate A
deactivate A
H <- D: |Invoke action|\nHelpdesk support
deactivate D
H -> S: Hi! Welcome to Helpdesk
deactivate H

alt Reset my password
  S -> H: I want to \nreset my password
  activate H
  H -> D
  activate D
  D <-> A: |Request|\nSTT
  activate A
  deactivate A
  D -> F
  activate F
  F <-> DS
  activate DS
  deactivate DS
  F -> D
  deactivate F
  D <-> A: |Request|\nTTS
  activate A
  deactivate A
  D -> H
  deactivate D
  H -> S: No problem! \nPlease provide me \nyour username
  deactivate H
end
```

## Allow Fulfillment to Store Help Ticket Data
**index.js**
```js
'use strict';
const http = require('http');
// Imports the Google Cloud client library
const Datastore = require('@google-cloud/datastore');
// Your Google Cloud Platform project ID
const projectId = 'REPLACE_WITH_YOUR_PROJECT_ID';
// Instantiates a client
const datastore = Datastore({
  projectId: projectId
});
// The kind for the new entity
const kind = 'ticket';
exports.dialogflowFirebaseFulfillment = (req, res) => {
  console.log('Dialogflow Request body: ' + JSON.stringify(req.body));
  // Get the city and date from the request
  let ticketDescription = req.body.queryResult['queryText']; // incidence is a required param
  //let name = req.body.result.contexts[0].parameters['given-name.original'];
  let username = req.body.queryResult.outputContexts[1].parameters['given-name.original'];
  let phone_number = req.body.queryResult.outputContexts[1].parameters['phone-number.original'];
  console.log('description is ' +ticketDescription);
  console.log('name is '+ username);
  console.log('phone number is '+ phone_number);
  function randomIntInc (low, high) {
    return Math.floor(Math.random() * (high - low + 1) + low);
  }
  let ticketnum = randomIntInc(11111,99999);
  // The Cloud Datastore key for the new entity
  const taskKey = datastore.key(kind);
  // Prepares the new entity
  const task = {
    key: taskKey,
    data: {
      description: ticketDescription,
      username: username,
      phoneNumber: phone_number,
      ticketNumber: ticketnum
    }
  };
  console.log("incidence is  " , task);
  // Saves the entity
  datastore.save(task)
  .then(() => {
    console.log(`Saved ${task.key}: ${task.data.description}`);
    res.setHeader('Content-Type', 'application/json');
    //Response to send to Dialogflow
    res.send(JSON.stringify({ 'fulfillmentText': "I have successfully logged your ticket, the ticket number is " + ticketnum + ". Someone from the helpdesk will reach out to you within 24 hours."}));
    //res.send(JSON.stringify({ 'fulfillmentText': "I have successfully logged your ticket, the ticket number is " + ticketnum + ". Someone from the helpdesk will reach out to you within 24 hours.", 'fulfillmentMessages': "I have successfully logged your ticket, the ticket number is " + ticketnum +  ". Someone from the helpdesk will reach out to you within 24 hours."}));
  })
  .catch((err) => {
    console.error('ERROR:', err);
    res.setHeader('Content-Type', 'application/json');
    res.send(JSON.stringify({ 'speech': "Error occurred while saving, try again later", 'displayText': "Error occurred while saving, try again later" }));
  });
}
```

**package.json**
```json
{
  "name": "dialogflowFirebaseFulfillment",
  "description": "This is the default fulfillment for a Dialogflow agents using Cloud Functions for Firebase",
  "version": "0.0.1",
  "private": true,
  "license": "Apache Version 2.0",
  "author": "Google Inc.",
  "engines": {
    "node": "8"
  },
  "scripts": {
    "start": "firebase serve --only functions:dialogflowFirebaseFulfillment",
    "deploy": "firebase deploy --only functions:dialogflowFirebaseFulfillment"
  },
  "dependencies": {
    "actions-on-google": "^2.2.0",
    "firebase-admin": "^5.13.1",
    "firebase-functions": "^2.0.2",
    "dialogflow": "^0.6.0",
    "dialogflow-fulfillment": "^0.5.0",
    "@google-cloud/datastore": "^1.1.0"
  }
}
```

