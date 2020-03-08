---
title: React Starter Kit—A Quick Start Guide to the React Ecosystem and Creating Your First App
author: Eve Prcello, Allex Banks
tags: react
url: https://learning.oreilly.com/learning-paths/learning-path-react/9781492035879/
---

# 1. Welcome to React
## React Developer Tools
- [react-detector](https://chrome.google.com/webstore/detail/react-detector/jaaklebbenondhkanegppccanebkdjlh?hl=en-US)
    - react-detectorは、Reactを使用しているWebサイトと使用していないWebサイトを確認できるChrome拡張機能です。
- [show-me-the-react](https://github.com/cymen/show-me-the-react)
    - これは、FirefoxとChromeで使用できる別のツールで、インターネットを閲覧するときにReactを検出します。
- [React Developer Tools](https://chrome.google.com/webstore/detail/react-developer-tools/fmkadmapgofadopljbjfkapdkoienihi?hl=en)
    - これは、ブラウザの開発者ツールの機能を拡張できるプラグインです。React要素を表示できる開発者ツールに新しいタブが作成されます。
	
# 2. Setting Up React Projects with Create React App
```sh
npm i -g create-react-app
create-react-app rate-react
cd $_
npm start
```

# 3. Creating a Stateless Functional Component with React
```js:Star.js
import React from 'react'
import './Star.css'

const Star = ({ selected, onClick }) =>
  <div className={(selected) ? "star selected" : "star"} onClick={onClick} />

export default Star
```

```css:Star.css
div.star {
  cursor: pointer;
  height: 25px;
  width: 25px;
  margin: 2px;
  background-color: #AAAAAA;
  clip-path: polygon(50% 0%, 61% 35%, 98% 35%, 68% 57%, 79% 91%, 50% 70%, 21% 91%, 32% 57%, 2% 35%, 39% 35%);
}

div.star.selected {
  background-color: red;
}
```

```js:App.js
import React from 'react';
import logo from './logo.svg';
import './App.css';
import Star from './Star.js'

function App() {
  return (
    <div className="App">
    <header className="App-header">
    <img src={logo} className="App-logo" alt="logo" />

    <p>Edit <code>src/App.js</code>!!!</p>

    <Star selected={false} />
    <Star selected={true} />
    <Star onClick={() => alert('the third star was clicked')} />
    </header>
    </div>
  );
}

export default App;
```

# 4. Handling State in Nested Components
```js:App.js
import React from 'react';
import logo from './logo.svg';
import './App.css';
import StarRating from './StarRating'

function App() {
  return (
    <div className="App">
    <header className="App-header">
    <img src={logo} className="App-logo" alt="logo" />

    <p>Welcome to the jungle!!!</p>

    <StarRating totalStars={7} starsSelected={3} onChange={newRating => alert(`new rating ${newRating}`)} />
    </header>
    </div>
  );
}

export default App;
```

```js:StarRating.js
import React, { Component } from 'react'
import Star from './Star'
import './StarRating.css'

const createArray = length => [...Array(length)]

export default class StarRating extends Component {
  constructor(props) {
    super(props)
    this.state = {
      starsSelected: props.starsSelected || 0
    }
    this.change = this.change.bind(this)
  }

  change(starsSelected) {
    this.setState({ starsSelected })
    // this.setState({ starsSelected: starsSelected })
    this.props.onChange(starsSelected)
  }

  render() {
    const { totalStars } = this.props
    const { starsSelected } = this.state
    return (
      <div className="star-rating">
      {createArray(totalStars).map((n, i) =>
        <Star key={i} selected={i<starsSelected} onClick={() => this.change(i+1)} />
      )}
      <p>{starsSelected} of {this.props.totalStars}</p>
      </div>
    )
  }
}

StarRating.defaultProps = {
  totalStars: 5
}
```

```css:StarRating.css
div.star-rating {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

div.star-rating p {
  flex-basis: 100%;
}
```

# 5. Managing Complex State Data with React
```js:App.js
import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import StarRating from './StarRating'

class App extends Component {
  constructor(props) {
    super(props)
    this.state = {
      features: [
        {
          id: 0,
          title: 'JSX',
          rating: 2
        },
        {
          id: 1,
          title: 'React DOM',
          rating: 2
        },
        {
          id: 2,
          title: 'Stateless Functional Components',
          rating: 1
        },
        {
          id: 3,
          title: 'Class Components',
          rating: 5
        },
        {
          id: 4,
          title: 'JSX 2',
          rating: 2
        },
      ]
    }

    this.rateFeature = this.rateFeature.bind(this)
  }

  rateFeature(id, rating) {
    const { features } = this.state
    this.setState({
      features: features.map(f => (f.id !== id) ? f : ({ ...f, rating }))
    })
  }

  render() {
    return (
      <div className="App">
      <header className="App-header">
      <img src={logo} className="App-logo" alt="logo" />

      <p>Welcome to the jungle!!!</p>

      <div className="features">
      {this.state.features.map(f =>
        <div key={f.id}>
        <h3>{f.title}</h3>
        <StarRating starsSelected={f.rating} />
        </div>
      )}
      </div>
      </header>
      </div>
    );
  }
}

export default App;
```js:StarRating.js
import React, { Component } from 'react'
import Star from './Star'
import './StarRating.css'

const createArray = length => [...Array(length)]

export default class StarRating extends Component {
  constructor(props) {
    super(props)
    this.state = {
      starsSelected: props.starsSelected || 0
    }
    this.change = this.change.bind(this)
  }

  change(starsSelected) {
    this.setState({ starsSelected })
    // this.setState({ starsSelected: starsSelected })
    this.props.onChange(starsSelected)
  }

  render() {
    const { totalStars } = this.props
    const { starsSelected } = this.state
    return (
      <div className="star-rating">
      {createArray(totalStars).map((n, i) =>
        <Star key={i} selected={i<starsSelected} onClick={() => this.change(i+1)} />
      )}
      <p>{starsSelected} of {this.props.totalStars}</p>
      </div>
    )
  }
}

StarRating.defaultProps = {
  totalStars: 5,
  onChange: f => f
}
```
