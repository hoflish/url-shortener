import React from 'react';
import PropTypes from 'prop-types';
import './InputField.css';

const InputField = props => {
  const { action, ...otherProps } = props;

  function onAction(e) {
    e.preventDefault();
    if (e.isTrusted && !e.repeat) {
      if (action) {
        action(e);
      }
    }
  }

  return <input onChange={onAction} className="input" {...otherProps} />;
};

InputField.propTypes = {
  placeholder: PropTypes.string.isRequired,
  action: PropTypes.func.isRequired,
};

export default InputField;
