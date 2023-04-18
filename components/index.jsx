// @flow
// react stuff
import * as React from 'react';
import { render } from 'react-dom';

// =====================
import AppRouting from './routes';

// Stylesheets
// import s from './index.less';
import './defaults.less';

const initialState = {};

window.renderApp = () => {
  render(
    <React.StrictMode>
      <AppRouting />
    </React.StrictMode>,
    document.getElementById('app')
  );
};