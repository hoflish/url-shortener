/* TODO: 
    [✔] # Validate user input - Original URL
    [✔] # Use ui toast for client errors
    [x] # Add onSend method to send payload to server
          + set timeout/deadline for client
          + handle response (errors or data)
    [x] # Disable shorten button when process
    [x] # 
*/
import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import { ToastContainer, toast } from 'react-toastify';
import axios from 'axios';
import { isWebURL } from '../../is_web_url';
import InputField from '../InputField/InputField';
import 'react-toastify/dist/ReactToastify.css';
import './ShortenSection.css';

const styles = theme => ({
  button: {
    margin: theme.spacing.unit,
  },
  input: {
    display: 'none',
  },
});

class ShortenSection extends React.Component {
  constructor(props) {
    super(props);
    this.state = { longURL: '' };
    this.toastId = null;
    this.notify = () => {
      if (!toast.isActive(this.toastId)) {
        this.toastId = toast.error('Unable to create short URL', {
          autoClose: 4000,
          className: 'custom-toast',
        });
      }
    };
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  onSend(l) {
    const self = this;
    const form = new FormData();
    form.append('longUrl', l);
    axios
      .post('http://192.168.99.100:8080/api/v1/url', form)
      .then(response => {
        console.log(response);
      })
      .catch(error => {
        console.log(error);
      });
  }

  handleChange(event) {
    const currentState = { longURL: event.target.value };
    const { hasError } = this.state;
    if (hasError) {
      currentState.hasError = undefined;
    }
    this.setState(currentState);
  }

  handleSubmit(event) {
    event.preventDefault();
    const { longURL: l } = this.state;
    if (l === '') return;

    if (!isWebURL.test(l)) {
      this.notify();
      return;
    }
    // trigger onSend method
    this.onSend(l);
  }

  render() {
    const { classes } = this.props;
    return (
      <section className="shorten section">
        <div>
          <ToastContainer hideProgressBar className="toast-container" />
        </div>
        <div className="container">
          <h1>Simplify your links</h1>
          <div className="input-container">
            <form onSubmit={this.handleSubmit}>
              {/* <input /> */}
              <InputField action={this.handleChange} placeholder="Your original URL here" />
              {/* <button type="button" /> */}
              <Button type="submit" variant="contained" className={classes.button}>
                Shorten URL
              </Button>
            </form>
          </div>
        </div>
      </section>
    );
  }
}

ShortenSection.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(ShortenSection);
