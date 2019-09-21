---
title: "Create serverless logic with Azure Functions"
tags: azure, azure-functions
url: https://docs.microsoft.com/learn/modules/create-serverless-logic-with-azure-functions/index
---

# Goal
- Decide if serverless computing is right for your business need
- Create an Azure Function app in the Azure portal
- Execute a function using triggers
- Monitor and test your Azure Function from the Azure portal

# Task
- [x] Decide if serverless computing is right for your business needs
- [x] Exercise - Create a function app in the Azure portal
- [x] Run your code on-demand with Azure Functions
- [ ] Exercise - Add logic to the function app
- [ ] Summary

# Supplement
## Decide if serverless computing is right for your business needs
**Function as a Service (FaaS)**
- Azure Logic Apps (Azure App Service)
- Azure Functions (C#, F#, JavaScript, etc.)

## Run your code on-demand with Azure Functions
**Triggers**
- BLOB storage
- Azure Cosmos DB
- Event Grid
- HTTP
- Microsoft Graph event
- Queue storage
- Service Bus
- Timer

## Exercise - Add logic to the function app
```js
module.exports = async function (context, req) {
    context.log('JavaScript HTTP trigger function processed a request.');

    if (req.query.name || (req.body && req.body.name)) {
        context.res = {
            // status: 200, /* Defaults to 200 */
            body: "Hello " + (req.query.name || req.body.name)
        };
    }
    else {
        context.res = {
            status: 400,
            body: "Please pass a name on the query string or in the request body"
        };
    }
};
```

## Summary
