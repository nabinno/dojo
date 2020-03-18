import React, { Component } from 'react';
import '../../assets/css/style.css';

const posts = [
  {
    id: 2,
    text: 'Lorem ipsum',
    user: {
      avatar: '/uploads/avatar1.png',
      username: 'Test User'
    }
  },
  {
    id: 1,
    text: 'Lorem ipsum',
    user: {
      avatar: '/uploads/avatar2.png',
      username: 'Test User 2'
    }
  }
];

export default class App extends Component {
  state = {
    posts: posts
  }

  // constructor(props) {
  //   super(props);
  //
  //   this.state = {
  //     posts：posts
  //   };
  // }

  render() {
    const { posts } = this.state;

    return (
      <div className="container">
        <div className="feed">
          {posts.map(
            (post, i) =>
              <div key={post.id} className="post">
                <div className="header">
                  <img src={post.user.avatar} />
                  <h2>{post.user.username}</h2>
                </div>
                <p className="content">
                  {post.text}
                </p>
              </div>
          )}
        </div>
      </div>
    );
  }
}
