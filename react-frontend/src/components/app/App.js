import React, { Component } from 'react';
import logo from '../../logo.svg';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import injectTapEventPlugin from 'react-tap-event-plugin';

import './App.css';
import {Querybox} from '../querybox/Querybox';

injectTapEventPlugin();

class App extends Component {
  render() {
    return (
      <MuiThemeProvider>
        <div className="App">
          <div className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
            <h2>Welcome to Coveo GraphQL Search</h2>
          </div>
          <p className="App-intro">
            Searching for 'Example Search Term'
          </p>
          <div className="search-section">
            <Querybox/>
          </div>
        </div>
      </MuiThemeProvider>
    );
  }
}

export default App;
