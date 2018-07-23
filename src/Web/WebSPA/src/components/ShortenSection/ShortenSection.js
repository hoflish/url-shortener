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
    API.post('url', qs.stringify({ longUrl: l }), {
      headers: {
        'content-type': 'application/x-www-form-urlencoded',
      },
    })
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
          const { status } = error.response;
          switch (status) {
            case 400:
            case 422:
              this.notifyErrors(UX.original_url_invalid);
              break;
            default:
              this.notifyErrors(UX.error_internal_server);
          }
        } else if (error.request) {
          this.notifyErrors(UX.error_internal_server);
        } else {
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
