import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import StarRating from './StarRating.js';

class App extends Component {
  constructor(props) {
    super(props);
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
    };

    this.rateFeature = this.rateFeature.bind(this);
  }

  rateFeature(id, rating) {
    const { features } = this.state;
    this.setState({
      features: features.map(f => (f.id !== id) ? f : ({ ...f, rating }))
    });
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />

          <p>Welcome to the jungle!!!</p>

          <div className="features">{
            this.state.features.map(
              f =>
                <div key={f.id}>
                  <h3>{f.title}</h3>
                  <StarRating starsSelected={f.rating} />
                </div>
            )
          }</div>
        </header>
      </div>
    );
  }
}

export default App;
