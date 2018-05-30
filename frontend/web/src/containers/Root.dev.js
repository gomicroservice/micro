import React from "react";
import { Provider } from "react-redux";
import { Route } from "react-router-dom";
import App from "./App";
import DevTools from "./DevTools";

const Root = ({ store }) => (
  <Provider store={store}>
    <div>
      <Route path="/" component={App} />
      <DevTools />
    </div>
  </Provider>
);

export default Root;
