import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import InputField from '../InputField/InputField';
import './ShortenSection.css';

const styles = theme => ({
  button: {
    margin: theme.spacing.unit,
  },
  input: {
    display: 'none',
  },
});

function ShortenSection(props) {
  const { classes } = props;
  return (
    <section className="shorten section">
      <div className="container">
        <h1>Simplify your links</h1>
        <div className="input-container">
          {/* <input /> */}
          <InputField placeholder="Your original URL here" />
          {/* <button type="button" /> */}
          <Button variant="contained" className={classes.button}>
            Shorten URL
          </Button>
        </div>
      </div>
    </section>
  );
}

ShortenSection.propTypes = {
  classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(ShortenSection);
