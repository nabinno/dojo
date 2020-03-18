---
title: Hands-On Full-Stack Web Development with GraphQL and React
author: Sebastian Grebe
tags: graphql, react
url: https://learning.oreilly.com/library/view/hands-on-full-stack-web/9781789134520/cover.xhtml
---

# 1. Preparing Your Development Environment
## 1.a. Application architecture
Node.js, Express.js, Apollo, SQL, Sequelize, and React.

### The basic setup
![](https://learning.oreilly.com/library/view/hands-on-full-stack-web/9781789134520/assets/b36a0c95-ddb1-4061-ab98-2760dc14a3be.png)

## 1.b. Setting up React
```sh
mkdir ~/graphbook
cd ~/graphbook
npm init
npm install --save react react-dom
```

### 1.b.1. Preparing and configuring webpack
```sh
mkdir public
touch public/index.html

cat <<EOF >public/index.html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-
  scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Graphbook</title>
  </head>
  <body>
    <div id="root"></div>
  </body>
</html>
EOF

npm install --save-dev @babel/core babel-eslint babel-loader @babel/preset-env @babel/preset-react clean-webpack-plugin css-loader eslint file-loader html-webpack-plugin style-loader url-loader webpack webpack-cli webpack-dev-server @babel/plugin-proposal-decorators @babel/plugin-proposal-function-sent @babel/plugin-proposal-export-namespace-from @babel/plugin-proposal-numeric-separator @babel/plugin-proposal-throw-expressions @babel/plugin-proposal-class-properties

npx install-peerdeps --dev eslint-config-airbnb

cat <<EOF >.eslintrc
{
  "extends": ["airbnb"],
  "env": {
    "browser": true,
    "node": true
  },
  "rules": {
    "react/jsx-filename-extension": "off"
  }
}
EOF

cat <<EOF >webpack.client.config.js
const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const buildDirectory = 'dist';
const outputDirectory = buildDirectory + '/client';
module.exports = {
  mode: 'development',
  entry: './src/client/index.js',
  output: {
    path: path.join(__dirname, outputDirectory),
    filename: 'bundle.js'
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader'
        }
      },
      {
        test: /\.css$/,
        use: ['style-loader', 'css-loader']
      }
    ]
  }, 
  devServer: {
    port: 3000,
    open: true
  },
  plugins: [
    new CleanWebpackPlugin([buildDirectory]),
    new HtmlWebpackPlugin({
      template: './public/index.html'
    })
  ]
};
EOF

mkdir -p src/client
```

### 1.b.2. Render your first React component
```js:src/client/index.js
import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';

ReactDOM.render(<App />, document.getElementById('root'));
```

```js:src/client/App.js
import React, { Component } from 'react';

export default class App extends Component {
  render() {
    return (
      <div>Hello World!</div>
    );
  }
}
```

```.babelrc
{
  "plugins": [
    ["@babel/plugin-proposal-decorators", { "legacy": true }],
    "@babel/plugin-proposal-function-sent",
    "@babel/plugin-proposal-export-namespace-from",
    "@babel/plugin-proposal-numeric-separator",
    "@babel/plugin-proposal-throw-expressions",
    ["@babel/plugin-proposal-class-properties", { "loose": false }]
  ],
  "presets": ["@babel/env","@babel/react"]
}
```

### 1.b.3. Rendering arrays from React state
```js:src/client/App.js
import React, { Component } from 'react';

const posts = [
  {
    id: 2,
    text: 'Lorem ipsum',
    user: {
      avatar: '/uploads/avatar1.png',
      username: 'Test User'
    }
  },
  {
    id: 1,
    text: 'Lorem ipsum',
    user: {
      avatar: '/uploads/avatar2.png',
      username: 'Test User 2'
    }
  }
];

export default class App extends Component {
  state = {
    posts: posts
  }

  // constructor(props) {
  //   super(props);
  //
  //   this.state = {
  //     posts：posts
  //   };
  // }

  render() {
    const { posts } = this.state;

    return (
      <div className="container">
        <div className="feed">
          {posts.map(
            (post, i) =>
              <div key={post.id} className="post">
                <div className="header">
                  <img src={post.user.avatar} />
                  <h2>{post.user.username}</h2>
                </div>
                <p className="content">
                  {post.text}
                </p>
              </div>
          )}
        </div>
      </div>
    );
  }
}
```

### 1.b.7. Production build with webpack
## 1.c. Useful development tools
### 1.c.1. Analyzing bundle size
## 1.d. Summary

# 2. Setting up GraphQL with Express.js
## Node.js and Express.js
### Setting up Express.js
### Running Express.js in development
## Routing in Express.js
### Serving our production build
## Using Express.js middleware
### Installing important middleware
### Express Helmet
### Compression with Express.js
### CORS in Express.js
## Combining Express.js with Apollo
### Writing your first GraphQL schema
### Implementing GraphQL resolvers
### Sending GraphQL queries
### Using multiples types in GraphQL schemas
### Writing your first GraphQL mutation
## Back end debugging and logging
### Logging in Node.js
### Debugging with Postman
## Summary

# 3. Connecting to The Database
## Using databases in GraphQL
### Installing MySQL for development
### Creating a database in MySQL
## Integrating Sequelize into our stack
### Connecting to a database with Sequelize
### Using a configuration file with Sequelize
## Writing database models
### Your first database model
### Your first database migration
### Importing models with Sequelize
## Seeding data with Sequelize
## Using Sequelize with Apollo
### Global database instance
### Running the first database query
## One-to-one relationships in Sequelize
### Updating the table structure with migrations
### Model associations in Sequelize
### Seeding foreign key data
## Mutating data with Sequelize
## Many-to-many relationships
### Model and migrations
#### Chat model
#### Message model
### Chats and messages in GraphQL
### Seeding many-to-many data
### Creating a new chat
### Creating a new message
## Summary

# 4. Integrating React into the Back end with Apollo
## Setting up Apollo Client
### Installing Apollo Client
### Testing the Apollo Client
### Binding the Apollo Client to React
## Using the Apollo Client in React
### Querying in React with the Apollo Client
#### Apollo HoC query
#### The Apollo Query component
## Mutations with the Apollo Client
### The Apollo Mutation HoC
### The Apollo Mutation component
### Updating the UI with the Apollo Client
#### Refetching queries
#### Updating the Apollo cache
#### Optimistic UI
#### Polling with the Query component
## Implementing chats and messages
### Fetching and displaying chats
### Fetching and displaying messages
### Sending messages through Mutations
## Pagination in React and GraphQL
## Debugging with the Apollo Client Developer Tools
## Summary

# 5. Reusable React Components
## Introducing React patterns
### Controlled components
### Stateless functions
### Conditional rendering
### Rendering child components
## Structuring our React application
### The React file structure
### Efficient Apollo React components
#### The Apollo Query component
#### The Apollo Mutation component
## Extending Graphbook
### The React context menu
#### FontAwesome in React
#### React helper components
#### The GraphQL updatePost mutation
#### The Apollo deletePost mutation
### The React application bar
### The React Context API versus Apollo Consumer
#### The React Context API
#### Apollo Consumer
## Documenting React applications
### Setting up&amp;#xA0;React Styleguidist
### React PropTypes
## Summary

# 6. Authentication with Apollo and React
## JSON Web Tokens
## localStorage versus cookie
## Authentication with GraphQL
### Apollo login mutation
### The React login form
### Apollo sign up mutation
### React sign up form
### Authenticating GraphQL requests
### Accessing the user context from resolver functions
#### Chats and messages
#### CurrentUser GraphQL query
### Logging out using React
## Summary

# 7. Handling Image Uploads
## Setting up Amazon Web Services
### Creating an AWS S3 bucket
### Generating AWS access keys
## Uploading images to Amazon S3
### GraphQL image upload mutation
### React image cropping and uploading
## Summary

# 8. Routing in React
## Setting up React Router
### Installing React Router
### Implementing your first route
### Secured routes
### Catch-all routes in React Router
## Advanced routing with React Router
### Parameters in routes
### Querying the user profile
### Programmatic navigation in React Router
### Remembering the redirect location
## Summary

# 9. Implementing Server-Side Rendering
## Introduction to server-side rendering
## SSR in Express.js
## Authentication with SSR
## Running Apollo queries with SSR
## Summary

# 10. Real-Time Subscriptions
## GraphQL and WebSockets
## Apollo Subscriptions
### Subscriptions on the Apollo Server
### Subscriptions on the Apollo Client
## Authentication with Apollo Subscriptions
## Notifications with Apollo Subscriptions
## Summary

# 11. Writing Tests
## Testing with Mocha
### Our first Mocha test
### Starting the back end with Mocha
### Verifying the correct routing
## Testing GraphQL with Mocha
### Testing the authentication
### Testing authenticated requests
## Testing React with Enzyme
## Summary

# 12. Optimizing GraphQL with Apollo Engine
## Setting up Apollo Engine
## Analyzing schemas with Apollo Engine
## Performance metrics with Apollo Engine
## Error tracking with Apollo Engine
## Caching with Apollo Server and the Client
## Summary

# 13. Continuous Deployment with CircleCI and Heroku
## Preparing the final production build
### Code-splitting with React Loadable and webpack
### Code-splitting with SSR
## Setting up Docker
### What is Docker?
### Installing Docker
### Dockerizing your application
#### Writing your first Dockerfile
### Building and running Docker containers
### Multi-stage Docker production builds
## Amazon Relational Database Service
## Configuring Continuous Integration
## Deploying applications to Heroku
## Summary

# 14. Other Books You May Enjoy
## Leave a review - let other readers know what you think
