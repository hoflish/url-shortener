import React, { Component } from 'react';
import { Switch, Route } from 'react-router-dom';
import Header from '../Header/Header';
import ShortenSection from '../ShortenSection/ShortenSection';
import NotFound from '../NotFound/NotFound';

class App extends Component {
  render() {
    return (
      <div>
        <Header />
        <div className="main">
          <Switch>
            <Route exact path="/" component={ShortenSection} />
            <Route component={({ location }) => <NotFound urlPath={location.pathname} />} />
          </Switch>
        </div>
      </div>
    );
  }
}

export default App;
