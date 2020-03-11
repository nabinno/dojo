---
title: React Hooks
author: reed barger
tags: react, react-hooks, javascript
title: https://learning.oreilly.com/videos/react-hooks/9781839210075/9781839210075-video1_1
---

# 1. Getting Started
## React Hook Intro
## What You Will Need for this Course
## Install React 16.7 to Use React Hooks
```sh
mkdir react_hooks
cd $_
create-react-app .

npm start
```

## What are React Hooks?

# 2. Moving from Classes to Function Components (useState, useEffect)
## Introducing the useState Hook
```js:AppClass.js
import React, { Component } from 'react';

class App extends Component {
  state = {
    count: 0
  };

  incrementCount = () => {
    this.setState({
      count: this.state.count + 1
    });
  };

  render() {
    return (
      <div className="App">
        <button
          onClick={this.incrementCount}
        >
          I was cliced {this.state.count} times
        </button>
      </div>
    );
  }
}

export default App;
```

```js:AppFunction.js
import React, { useState } from 'react';

const App = () => {
  const [count, setCount] = useState(0);

  const incrementCount = () => {
    setCount(count + 1);
  };

  return (
    <button onClick={incrementCount}>
      I was clicked {count} times
    </button>
  );
};

export default App;
```

## Use Previous State with useState


## Toggle State with useState
## Introducing the useEffect Hook
## Cleaning up Side Effects in useEffect
## Using / Cleaning up Multiple Listeners in useEffect
## Cleaning up Listeners without a Supportive API
## Comparing Function Component and Class Components

# 3. Building Stateful Components with Functions
## Implementing a Login Form with Multiple State Values
## Implementing a Register Form with a Single State Value
## Comparing Ways of Managing State

# 4. Data Fetching with Hooks / Replacing Class Lifecycle Methods (useEffect, useRef)
## Intro to Data Fetching Project
## Fetching Data on component Mount with useEffect
## Using Async / Await for Fetching Data in useEffect
## Fetching Search Results on Component Update with useEffect
## Fetching Data upon Submitting Form
## Using the useRef Hook on our Search Input
## Displaying Loading State with useState
## Error Handling and Displaying Errors with useState
## Styling our Project with TailwindCSS (Optional)

# 5. Building a Complete CRUD App with React Hooks / Replacing Redux
## Project Setup for our CRUD App
## Avoiding Props Drilling with React Context and the useContext Hook
## Replacing Redux with the useReducer Hook
## Combining useContext and useReducer to Make Initial App State
## Styling our TodoList Component with TailwindCSS
## Toggling Todos / 'TOGGLE_TODO' case
## Removing Todos / 'REMOVE_TODO' case
## Adding Todos and TodoForm Component / 'ADD_TODO' case
## Updating Todos / 'UPDATE_TODO' case
## Improving our App

# 6. Connecting our App to an API
## Creating / Deploying our API to Persist App Data
## Creating a Custom Hook to Fetch Initial App Data
## Delete Request to Remove Todos
## Performing Post Request to Add Todos
## Performing Patch Request to Toggle Todos
## Finishing our App
