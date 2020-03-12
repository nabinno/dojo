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
## a. Introducing the useState Hook
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

## b. Use Previous State with useState
```js:AppClass.js
import React, { Component } from 'react';

class App extends Component {
  state = {
    count: 0
  };

  incrementCount = () => {
    this.setState(
      prevState => ({
        count: prevState.count + 1
      })
    );
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
    setCount(prevCount => prevCount + 1);
  };

  return (
    <button onClick={incrementCount}>
      I was clicked {count} times
    </button>
  );
};

export default App;
```

## c. Toggle State with useState
```js:AppClass.js
import React, { Component } from 'react';

class App extends Component {
  state = {
    count: 0,
    isOn: false
  };

  incrementCount = () => {
    this.setState(
      prevState => ({
        count: prevState.count + 1
      })
    );
  };

  toggleLight = () => {
    this.setState(prevState => ({
      isOn: !prevState.isOn
    }));
  }

  render() {
    return (
      <>
        <div className="App">
          <h2>Counter</h2>
          <button
            onClick={this.incrementCount}
          >
            I was cliced {this.state.count} times
          </button>

          <h2>Toggle Light</h2>
          <div
            style={{
              height: '50px',
              width: '50px',
              background: this.state.isOn ? 'yellow' : 'grey'
            }}
            onClick={this.toggleLight}
          >
          </div>
        </div>
      </>
    );
  }
}

export default App;
```

```js:AppFunction.js
import React, { useState } from 'react';

const App = () => {
  const [count, setCount] = useState(0);

  const [isOn, setIsOn] = useState(false);

  const incrementCount = () => {
    setCount(prevCount => prevCount + 1);
  };

  const toggleLight = () => {
    setIsOn(prevIsOn => !prevIsOn);
  };

  return (
    <>
      <h2>Counter</h2>
      <button onClick={incrementCount}>
        I was clicked {count} times
      </button>

      <h2>Toggle Light</h2>
      <img
        src={
          isOn
            ? 'https://icon.now.sh/highlight/fd0'
            : 'https://icon.now.sh/highlight/aaa'
        }
        style={{
          height: '50px',
          width: '50px',
        }}
        alt="Flashlight"
        onClick={toggleLight}
      />
    </>
  );
};

export default App;
```

## d. Introducing the useEffect Hook
```js:AppClass.js
import React, { Component } from 'react';

class App extends Component {
  state = {
    count: 0,
    isOn: false
  };

  componentDidMount() {
    document.title = `You have been clicked ${this.state.count} times`;
  }

  componentDidUpdate() {
    document.title = `You have been clicked ${this.state.count} times`;
  }
```

```js:AppFunction.js
import React, { useState, useEffect } from 'react';

const App = () => {
  const [count, setCount] = useState(0);
  const [isOn, setIsOn] = useState(false);

  useEffect(() => {
    document.title = `You have clicked ${count} times`;
  });
```

## e. Cleaning up Side Effects in useEffect
```js:AppClass.js
import React, { Component } from 'react';

class App extends Component {
  state = {
    count: 0,
    isOn: false,
    x: null,
    y: null
  };

  componentDidMount() {
    document.title = `You have been clicked ${this.state.count} times`;
    window.addEventListener('mousemove', this.handleMouseMove);
  }

  componentDidUpdate() {
    document.title = `You have been clicked ${this.state.count} times`;
  }

  componentWillUnmound() {
    window.removeEventListener('mousemove', this.handleMouseMove);
  }

  handleMouseMove = event => {
    this.setState({
      x: event.pageX,
      y: event.pageY
    });
  }

  incrementCount = () => {
    this.setState(
      prevState => ({
        count: prevState.count + 1
      })
    );
  };

  toggleLight = () => {
    this.setState(prevState => ({
      isOn: !prevState.isOn
    }));
  }

  render() {
    return (
      <>
        <div className="App">
          <h2>Mouse Position</h2>
          <p>X position: {this.state.x}</p>
          <p>Y position: {this.state.y}</p>


        </div>
      </>
    );
  }
}

export default App;
```

```js:AppFunction.js
import React, { useState, useEffect } from 'react';

const App = () => {
  const [count, setCount] = useState(0);
  const [isOn, setIsOn] = useState(false);
  const [mousePosition, setMousePosition] = useState({ x: null, y: null });

  useEffect(() => {
    document.title = `You have clicked ${count} times`;
    window.addEventListener('mousemove', handleMouseMove);

    return () => {
      window.removeEventListener('mousemove', handleMouseMove);
    };
  }, [count]);

  const handleMouseMove = event => {
    setMousePosition({
      x: event.pageX,
      y: event.pageY
    });
  };

  const incrementCount = () => {
    setCount(prevCount => prevCount + 1);
  };

  const toggleLight = () => {
    setIsOn(prevIsOn => !prevIsOn);
  };

  return (
    <>
      <h2>Mouse Position</h2>
      {JSON.stringify(mousePosition, null, 2)}
    </>
  );
};

export default App;
```

## f. Using / Cleaning up Multiple Listeners in useEffect
## g. Cleaning up Listeners without a Supportive API
## h. Comparing Function Component and Class Components

# 3. Building Stateful Components with Functions
## a. Implementing a Login Form with Multiple State Values
## b. Implementing a Register Form with a Single State Value
## c. Comparing Ways of Managing State

# 4. Data Fetching with Hooks / Replacing Class Lifecycle Methods (useEffect, useRef)
## a. Intro to Data Fetching Project
## b. Fetching Data on component Mount with useEffect
## c. Using Async / Await for Fetching Data in useEffect
## d. Fetching Search Results on Component Update with useEffect
## e. Fetching Data upon Submitting Form
## f. Using the useRef Hook on our Search Input
## g. Displaying Loading State with useState
## h. Error Handling and Displaying Errors with useState
## i. Styling our Project with TailwindCSS (Optional)

# 5. Building a Complete CRUD App with React Hooks / Replacing Redux
## a. Project Setup for our CRUD App
## b. Avoiding Props Drilling with React Context and the useContext Hook
## c. Replacing Redux with the useReducer Hook
## d. Combining useContext and useReducer to Make Initial App State
## e. Styling our TodoList Component with TailwindCSS
## f. Toggling Todos / 'TOGGLE_TODO' case
## g. Removing Todos / 'REMOVE_TODO' case
## h. Adding Todos and TodoForm Component / 'ADD_TODO' case
## i. Updating Todos / 'UPDATE_TODO' case
## j. Improving our App

# 6. Connecting our App to an API
## a. Creating / Deploying our API to Persist App Data
## b. Creating a Custom Hook to Fetch Initial App Data
## c. Delete Request to Remove Todos
## d. Performing Post Request to Add Todos
## e. Performing Patch Request to Toggle Todos
## f. Finishing our App
