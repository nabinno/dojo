import React from 'react';
import { connect } from 'react-redux';
import logo from './logo.svg';
import Form from './components/Form.js';
import Grid from './components/Grid.js';
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
