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
