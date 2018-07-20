import React from 'react';
import PropTypes from 'prop-types';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import { toast } from 'react-toastify';
import Tooltip from '@material-ui/core/Tooltip';
import { ContentCopy } from '@material-ui/icons';
import { UXMessages as UX } from '../../UXMessages';
import './AlertDialog.css';

class AlertDialog extends React.Component {
  constructor(props) {
    super(props);
    this.toastId = null;
    this.copyToClipboard = this.copyToClipboard.bind(this);
  }

  notify(err) {
    if (!toast.isActive(this.toastId)) {
      this.toastId = toast(err, {
        autoClose: 4000,
        className: 'custom-toast',
      });
    }
  }

  copyToClipboard(event) {
    event.preventDefault();
    this.input.select();
    document.execCommand('copy');
    event.target.focus();
    this.notify(UX.action_copy_to_clipboard);
  }

  render() {
    const { open, onClose, shorturl } = this.props;
    return (
      <div>
        <Dialog
          open={open}
          onClose={onClose}
          aria-labelledby="alert-dialog-title"
          aria-describedby="alert-dialog-description"
        >
          <DialogTitle id="alert-dialog-title">{UX.action_shorturl_created}</DialogTitle>
          <DialogContent>
            <DialogContentText id="alert-dialog-description">
              {/* TODO: STYLE SHORT URL and ADD CREATED MESSAGE */}
              <input
                ref={input => (this.input = input)}
                value={shorturl}
                type="text"
                id="shortUrl"
                readOnly
              />
              <Tooltip title="Copy" placement="top">
                <Button onClick={this.copyToClipboard} className="clipboardIcon">
                  <ContentCopy />
                </Button>
              </Tooltip>
            </DialogContentText>
          </DialogContent>
          <DialogActions>
            <Button onClick={onClose} color="primary" autoFocus>
              {UX.done}
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }
}

AlertDialog.propTypes = {
  open: PropTypes.bool.isRequired,
  onClose: PropTypes.func.isRequired,
  shorturl: PropTypes.string.isRequired,
};

export default AlertDialog;
