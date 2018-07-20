import React from 'react';
import PropTypes from 'prop-types';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import { UXMessages as UX } from '../../UXMessages';

class AlertDialog extends React.Component {
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
              {/* TODO: STYLE SHORT URL  */}
              {shorturl}
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
