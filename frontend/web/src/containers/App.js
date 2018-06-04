import React, { Component } from "react";
import CustomizedForm from "../components/CustomizedForm";
import "./App.css";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <div className="App">
        <CustomizedForm />
      </div>
    );
  }
}

export default App;
