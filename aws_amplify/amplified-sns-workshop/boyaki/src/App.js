import React from 'react';
import './App.css';

import { withAuthenticator } from 'aws-amplify-react';

import {
  HashRouter,
  Switch,
  Route,
  Redirect,
} from 'react-router-dom';

import { makeStyles, createMuiTheme, ThemeProvider } from '@material-ui/core/styles';
import CssBaseline from '@material-ui/core/CssBaseline';

import AllPosts from './containers/AllPosts';
import PostsBySpecifiedUser from './containers/PostsBySpecifiedUser';

import Amplify from '@aws-amplify/core';
import PubSub from '@aws-amplify/pubsub'
import awsmobile from './aws-exports';

Amplify.configure(awsmobile);
PubSub.configure(awsmobile)

const drawerWidth = 240;

const theme = createMuiTheme({
  palette: {
    type: 'dark',
    primary: {
      main: '#1EA1F2',
      contrastText: "#fff",
    },
    background: {
      default: '#15202B',
      paper: '#15202B',
    },
    divider: '#37444C',
  },
  overrides: {
    MuiButton: {
      color: 'white',
    },
  },
  typography: {
    fontFamily: [
      'Arial',
    ].join(','),
  },
  status: {
    danger: 'orange',
  },
});

const useStyles = makeStyles(theme => ({
  root: {
    display: 'flex',
    height: '100%',
    width: 800,
    marginLeft: 'auto',
    marginRight: 'auto',
  },
  appBar: {
    marginLeft: drawerWidth,
  },
  drawer: {
    width: drawerWidth,
    flexShrink: 0,
  },
  drawerPaper: {
    width: drawerWidth,
  },
  toolbar: theme.mixins.toolbar,
  content: {
    flexGrow: 1,
    backgroundColor: theme.palette.background.default,
    padding: theme.spacing(3),
  },
}));

function App() {
  const classes = useStyles();
  return (
    <div className={classes.root} >
      <ThemeProvider theme={theme}>
        <CssBaseline />
        <HashRouter>
          <Switch>
            <Route exact path='/' component={AllPosts} />
            <Route exact path='/global-timeline' component={AllPosts} />
            <Route exact path='/:userId' component={PostsBySpecifiedUser}/>
            <Redirect path="*" to="/" />
          </Switch>
        </HashRouter>
      </ThemeProvider>
    </div>
  );
}

export default withAuthenticator(App, {
  signUpConfig: {
    hiddenDefaults: ['phone_number']
  }
});
