/* TODO: (hoflish)
    [✔] # Validate user input - Original URL
    [✔] # Use ui toast for client errors
    [✔] # Add onSend method to send payload to server
          + [✔] handle response (errors)
          + [✔] set timeout/deadline for client
          + [✔] handle response success data
         
    [x] # Disable shorten button when process
          + [x] Use bluebird.js for more control of Promise
    [x] # log to a logging service instead of console ??
    [x] # Add remove button to clean input field
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
import AlertDialog from '../AlertDialog/AlertDialog';
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
    this.state = { longURL: '', open: false, shortURL: '' };
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
    this.handleClickOpen = this.handleClickOpen.bind(this);
    this.handleClose = this.handleClose.bind(this);
  }

  onSend(l) {
    API.post('url', qs.stringify({ longUrl: l }))
      .then(response => {
        // Success
        const res = response.data;
        this.setState({ shortURL: res.data.short_url }, () => {
          this.handleClickOpen();
        });
      })
      .catch(error => {
        // Error
        if (error.response) {
          // The request was made and the server responded with a status code
          // that falls out of the range of 2xx
          // Note: validation errors are rarely reached because of the client-side validation
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

  handleClose = () => {
    this.setState({ open: false });
  };

  handleClickOpen = () => {
    this.setState({ open: true });
  };

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

  handleChange(event) {
    this.setState({ longURL: event.target.value });
  }

  render() {
    const { classes } = this.props;
    const { open, shortURL } = this.state;
    return (
      <section className="shorten section">
        <div className="container">
          <h1>Simplify your links</h1>
          <div className="input-container">
            <form onSubmit={this.handleSubmit}>
              <InputField action={this.handleChange} placeholder="Your original URL here" />
              <Button type="submit" variant="contained" className={classes.button}>
                Shorten URL
              </Button>
            </form>
          </div>
        </div>
        <ToastContainer hideProgressBar className="toast-container" />
        <AlertDialog open={open} shorturl={shortURL} onClose={this.handleClose} />
      </section>
    );
  }
}

ShortenSection.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(ShortenSection);
