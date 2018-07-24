import React, { Component } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Header from '../Header/Header';
import ShortenSection from '../ShortenSection/ShortenSection';
import NotFound from '../NotFound/NotFound';

class App extends Component {
  render() {
    return (
      <Router>
        <div>
          <Header />
          <div className="main">
            <Switch>
              <Route exact path="/" component={ShortenSection} />
              <Route component={({ location }) => <NotFound url={location} />} />
            </Switch>
          </div>
        </div>
      </Router>
    );
  }
}

export default App;
