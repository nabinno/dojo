import React, { Component } from 'react';
import Star from './Star.js';
import './StarRating.css';

const createArray = length => [...Array(length)];

export default class StarRating extends Component {
  constructor(props) {
    super(props);
    this.state = {
      starsSelected: props.starsSelected || 0
    };
    this.change = this.change.bind(this);
  }

  change(starsSelected) {
    this.setState({ starsSelected });
    // this.setState({ starsSelected: starsSelected })
    this.props.onChange(starsSelected);
  }

  render() {
    const { totalStars } = this.props;
    const { starsSelected } = this.state;
    return (
      <div className="star-rating">
        {
          createArray(totalStars)
            .map(
              (n, i) =>
                <Star key={i}
                      selected={i<starsSelected}
                      onClick={() => this.change(i+1)} />
            )
        }
        <p>{starsSelected} of {this.props.totalStars}</p>
      </div>
    );
  }
}

StarRating.defaultProps = {
  totalStars: 5,
  onChange: f => f
};
