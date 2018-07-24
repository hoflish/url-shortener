import React, { Component } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import API from '../../api';
import Header from '../Header/Header';
import ShortenSection from '../ShortenSection/ShortenSection';
import NotFound from '../NotFound/NotFound';

const sidRegex = /^([a-zA-Z0-9_-]){9,12}$/; // Characters that make short id
const baseShorURL = 'http://localhost:3000/';

const goToOriginalURL = ({ match }) => {
  const { sid } = match.params;
  if (!sidRegex.test(sid)) {
    return <Route component={({ location }) => <NotFound url={location} />} />;
  }

  API.get(`url?shortUrl=${baseShorURL}${sid}`)
    .then(response => {
      // Success
      const { data } = response.data;
      window.location.href = data.long_url;
      return null;
    })
    .catch(error => {
      // Error
      if (error.response) {
        const { status, message } = error.response.data;
        if (status === 404) {
          return <Route component={({ location }) => <NotFound url={location} />} />;
        }
        console.log(message);
      } else if (error.request) {
        // TODO: Render server errors
        console.log(error.request, '503 Service Unavailable Error');
      } else {
        console.log('Oops! something went wrong');
      }
      return null;
    });
  return null;
};

class App extends Component {
  render() {
    return (
      <Router>
        <div>
          <Header />
          <div className="main">
            <Switch>
              <Route exact path="/" component={ShortenSection} />
              <Route path="/:sid" component={goToOriginalURL} />
              <Route component={({ location }) => <NotFound url={location} />} />
            </Switch>
          </div>
        </div>
      </Router>
    );
  }
}

export default App;
