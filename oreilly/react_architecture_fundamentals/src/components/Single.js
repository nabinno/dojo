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
