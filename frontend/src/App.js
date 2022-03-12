import {
  BrowserRouter as Router, Route
} from 'react-router-dom';
import MainNavigate from './ui/MainNavigate';
import React, {Fragment} from 'react';

function App() {
  return (
      <Router>
        <Fragment>
        <MainNavigate/>
        </Fragment>
      </Router>
  );
}

export default App;