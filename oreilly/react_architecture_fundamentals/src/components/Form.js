import React from 'react';

class Form extends React.Component {
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

export default Form;
