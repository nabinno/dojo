import React, { useState, useEffect, useReducer } from 'react';

import API, { graphqlOperation } from '@aws-amplify/api';
import { useParams } from 'react-router';

import { listPostsBySpecificOwner } from '../graphql/queries';
import { onCreatePost } from '../graphql/subscriptions';
import PostList from '../components/PostList';
import Sidebar from './Sidebar';

const SUBSCRIPTION = 'SUBSCRIPTION';
const INITIAL_QUERY = 'INITIAL_QUERY';
const ADDITIONAL_QUERY = 'ADDITIONAL_QUERY';

const reducer = (state, action) => {
  switch (action.type) {
  case INITIAL_QUERY:
    return action.posts;
  case ADDITIONAL_QUERY:
    return [...state, ...action.posts]
  case SUBSCRIPTION:
    return [action.post, ...state]
  default:
    return state;
  }
};

export default function PostsBySpecifiedUser() {
  const { userId } = useParams();

  const [posts, dispatch] = useReducer(reducer, []);
  const [nextToken, setNextToken] = useState(null);
  const [isLoading, setIsLoading] = useState(true);

  const getPosts = async (type, nextToken = null) => {
    const res = await API.graphql(graphqlOperation(listPostsBySpecificOwner, {
      owner: userId,
      sortDirection: 'DESC',
      limit: 20,
      nextToken: nextToken,
    }));
    console.log(res);
    dispatch({ type: type, posts: res.data.listPostsBySpecificOwner.items })
    setNextToken(res.data.listPostsBySpecificOwner.nextToken);
    setIsLoading(false);
  };

  const getAdditionalPosts = () => {
    if (nextToken === null) return; //Reached the last page
    getPosts(ADDITIONAL_QUERY, nextToken);
  };

  useEffect(() => {
    getPosts(INITIAL_QUERY);

    const subscription = API.graphql(graphqlOperation(onCreatePost)).subscribe({
      next: (msg) => {
        const post = msg.value.data.onCreatePost;
        if (post.owner !== userId) return;
        dispatch({ type: SUBSCRIPTION, post: post });

      }

    });
    return () => subscription.unsubscribe();

  }, []);


  return (
    <React.Fragment>
      <Sidebar
        activeListItem='profile'
      />
      <PostList
        isLoading={isLoading}
        posts={posts}
        getAdditionalPosts={getAdditionalPosts}
        listHeaderTitle={userId}
      />
    </React.Fragment>
  );
}
