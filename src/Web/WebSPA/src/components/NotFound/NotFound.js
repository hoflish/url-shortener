import React from 'react';
import PropTypes from 'prop-types';

export default function NotFound({ urlPath }) {
  const fullUrl = window.location.href || `http://localhost/${urlPath}`;
  return (
    <div className="container">
      <div className="error-box">
        <p>
          404 Not found - the page <b>{fullUrl}</b> does not exist.
        </p>
      </div>
    </div>
  );
}

NotFound.propTypes = {
  urlPath: PropTypes.string.isRequired,
};
