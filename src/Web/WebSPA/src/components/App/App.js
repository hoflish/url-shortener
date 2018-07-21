import React, { Component } from 'react';
import { Route, Switch } from 'react-router-dom';
import Header from '../Header/Header';
import ShortenSection from '../ShortenSection/ShortenSection';

const Main = () => (
  <div>
    <Header />
    <ShortenSection />
  </div>
);

class App extends Component {
  render() {
    return (
      <Switch>
        <Route exact path="/" component={Main} />
      </Switch>
    );
  }
}

export default App;
