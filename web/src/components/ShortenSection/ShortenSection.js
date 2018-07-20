/* TODO: 
    [✔] # Validate user input - Original URL
    [✔] # Use ui toast for client errors
    [x] # Add onSend method to send payload to server
          + [✔] handle response (errors)
          + [✔] set timeout/deadline for client
          + [x] handle response success data
         
    [x] # Disable shorten button when process
    [x] # TODO:(hoflish): log to a logging service instead of console ??

*/
import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import { ToastContainer, toast } from 'react-toastify';
import qs from 'qs';
import { isWebURL } from '../../isWebURL';
import API from '../../api';
import { UXMessages as UX } from '../../UXMessages';
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
    this.notifyErrors = err => {
      if (!toast.isActive(this.toastId)) {
        this.toastId = toast.error(err, {
          autoClose: 4000,
          className: 'custom-toast',
        });
      }
    };
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  async onSend(l) {
    const self = this;
    API.post('url', qs.stringify({ longUrl: l }))
      .then(response => {
        // TODO: show response short URL in a modal with copy feature
        // Success
        console.log(response);
      })
      .catch(error => {
        // Error
        if (error.response) {
          // The request was made and the server responded with a status code
          // that falls out of the range of 2xx
          const { status } = error.response;
          const { message } = error.response.data;
          switch (status) {
            case 400:
              this.notifyErrors(UX.original_url_invalid);
              break;
            case 422:
              this.notifyErrors(message);
              break;
            case 500:
              // log this error
              this.notifyErrors(UX.error_internal_server);
              break;
            default:
              this.notifyErrors(message);
          }
        } else if (error.request) {
          // The request was made but no response was received
          const req = error.request;
          if (req.readyState === 4 && req.status === 0) {
            this.notifyErrors(UX.error_unexpected);
          }
          console.log(req.message);
        } else {
          // Something happened in setting up the request that triggered an Error
          if (error.code === 'ECONNABORTED') {
            this.notifyErrors(UX.error_request_timeout);
          }
          console.log('Error', error.message);
        }
        console.log(error.config);
      });
  }

  handleChange(event) {
    this.setState({ longURL: event.target.value });
  }

  handleSubmit(event) {
    event.preventDefault();
    const { longURL: l } = this.state;
    if (l === '') return;

    if (!isWebURL.test(l)) {
      this.notifyErrors(UX.original_url_invalid);
      return;
    }
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
