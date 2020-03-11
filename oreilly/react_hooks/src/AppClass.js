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
