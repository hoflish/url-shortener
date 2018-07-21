import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router } from 'react-router-dom';
import App from './components/App/App';
import './main.css';

const MyApp = () => (
  <Router>
    <App />
  </Router>
);

ReactDOM.render(<MyApp />, document.getElementById('root'));
