import React from 'react';

export default function({ url }) {
  return (
    <div className="container">
      <p>
        404: Page not found - the page <b>{window.location.href || url.pathname}</b> does not exist.
      </p>
    </div>
  );
}
