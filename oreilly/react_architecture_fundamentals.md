---
title: React Architecture Fundamentals
author: emmanuel henri, eve porcello, lin clark, alex banks
tags: react, javascript
url: https://learning.oreilly.com/learning-paths/learning-path-react/9781492039440/continue
---

# 1. A Cartoon Guide to the Wilds of Data Handling in React
# 2. Redux

# 3. Flux/Redux Flow of Data
```
## Flux
Actions -> { Dispatcher -> Store -> View -> Actions }...

## Redux
Reducers -> { Store -View -> Actions ->  Reducers }...
```

# 4. Initial Setup for Redux
```sh
mkdir react_architecture_fundamentals
cd $_
create-react-app .
yarn add react-redux redux redux-immutable-state-invariant redux-thunk
```

# 5. Code Our Redux Actions
```js:store/actions.js
export const getInitialNotes = () => {
  return {
    type: 'GET_NOTES'
  };
};

export const addNewNote = (note) => {
  return {
    type: 'ADD_NOTE',
    note
  };
};

export const removeNote = (note) => {
  return {
    type: 'REMOVE_NOTE',
    note
  };
};
```

# 6. Code Our Redux Reducers
```js:store/reducers.js
const initialState = {
  note: [
    {
      id: '27812b',
      title: 'An example note',
      details: 'This is an example of a note'
    },
    {
      id: '278922',
      title: 'Another example note',
      details: 'This is an example of a note.... again'
    }
  ],
  name: 'Nab'
};

export default (state = initialState, action) => {
  switch (action.type) {
  case 'ADD_NOTE':
    return {
      ...state,
      notes: [...state.notes, action.note]
    };
  case 'REMOVE_NOTE':
    return {
      ...state,
      notes: state.notes.filter(note => note !== action.note)
    };
  case 'GET_NOTES':
    return {
      ...state
    };
  default:
    return state;
  }
};
```

# 7. Redux Project Base Store
```js:store/store.js
import { createStore, applyMiddleware, compose } from 'redux';
import rooReducer from './reducers.js';
import reduxImmutableStateInvariant from 'redux-immutable-state-invariant';
import thunk from 'redux-thunk';

export default function configureStore(initialState) {
  return createStore(
    rootReducer,
    initialState,
    compose(
      applyMiddleware(thunk, reduxImmutableStateInvariant()),
      window.devToolsExtension ? window.devToolsExtension() : f => f
    )
  );
}
```

```js:index.js
import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';
import { Provider } from 'react-redux';
import configureStore from './store/store.js';
import { getInitialNotes } from './store/actions';

const store = configureStore();
store.dispatch(getInitialNotes());

ReactDOM.render(
  <Provider store={store}>
    <App />
  </Provider>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
```

# 8. Update Our App View
```js:App.js
import React from 'react';
import { connect } from 'react-redux';
import logo from './logo.svg';
import './App.css';
import { getInitialNotes, addNewNote, removeNote } from './store/actions.js';

const styles = {
  textAlign: 'center',
  margin: 0,
  padding: 0,
  fontFamily: 'sans-serif'
};

class App extends React.Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.js</code> and save to reload.
          </p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
        </header>

        <Header name={this.props.name} />
        <Form addNewNote={this.props.addNewNote} />
        <Grid notes={this.props.notes} removeNote={this.props.removeNote} />
      </div>
    );
  }
}

const mapDispatchProps = (dispatch, ownProps) => {
  return {
    getInitialNotes: () => {
      dispatch(getInitialNotes());
    },
    AddNewNote: (note) => {
      dispatch(addNewNote(note));
    },
    removeNote: (id) => {
      dispatch(removeNote(id));
    },
  };
};

const mapStateTopProps = (state, ownProps) => {
  return {
    notes: state.notes,
    name: state.name
  };
};

export default connect(mapDispatchProps, mapStateTopProps)(App);
```

# 9. Updating Our Form View
```js:components/Form.js
import React from 'react';

class From extends React.Component {
  render() {
    return (
      <form
        onSubmit={event => {
          event.preventDefault();
          const note = {
            id: require('crypto').randomBytes(5).toString('hex'),
            title: this.refs.title.value,
            details: this.refs.details.value
          };
          this.refs.title.value = '';
          this.refs.details.value = '';
          this.props.addNewNote(note);
        }}
        className="col s12"
      >
        <div className="row">
          <div className="input-field col s3">
            <input
              id="title"
              name="currentTitle"
              type="text"
              ref="title"
              classNmae="validate"
            />
            <label htmlFor="title">Title</label>
          </div>
          <div className="input-field col s7">
            <input
              id="currentDetails"
              name="currentDetails"
              type="text"
              ref="details"
            />
            <label htmlFor="details">Details</label>
          </div>
          <div className="input-field col s2">
            <button
              className="btn-large waves-effect waves-light"
              name="action"
              type="submit"
            >Add note</button>
          </div>
        </div>
      </form>
    );
  }
}
```

# 10. Updating Our Grid and Single View
```js:components/Grid.js
import React from 'react';
import Single from './Single.js';

class Grid extends React.Component {
  removeNote(id) {
    this.props.removeNote(id);
  }

  renderItems() {
    return this.props.notes.map(
      item =>
        <Single
          key={item.id}
          note={item}
          removeNote={this.removeNote.bind(this)}
        />
    );
  }

  render() {
    return (
      <div classNmae="row">
        <ul>
          {this.renderItems()}
        </ul>
      </div>
    );
  }
}

export default Grid;
```

```js:components/Single.js
import React from 'react';

const Single = (props) => {
  return (
    <li className="col s4">
      <div className="card teal darken-1">
        <div className="card-content white-text">
          <span className="card-title">{props.note.title}</span>
          <p>{props.note.details}</p>
        </div>
        <div className="card-action">
          <a onClick={() => props.removerNote(props.note)}>Delete</a>
        </div>
      </div>
    </li>
  );
};

export default Single;
```

