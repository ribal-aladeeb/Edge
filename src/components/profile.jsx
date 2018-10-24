import React, { Component } from "react";

export default class Profile extends Component {
  state = {
    firstName: "",
    lastName: "",
    email: "",
    password: ""
  };
  change = e => {
    this.setState({
      [e.target.name]: e.target.value
    });
  };
  onSubmit = e => {
    e.preventDefault();
    this.props.onSubmit(this.state);
    console.log(this.state);
  };

  onEdit = e => {
    this.setState({
      firstName: "",
      lastName: "",
      email: "",
      password: ""
    });
    //e.preventDefault();
    //this.props.onEdit(this.state);
    //console.log(this.state);
  };

  render() {
    return (
      <form>
        <input
          name="firstName"
          placeholder="First Name"
          value={this.state.firstName}
          onChange={e => this.change(e)}
        />
        <br />
        <input
          name="lastName"
          placeholder="Last Name"
          value={this.state.lastName}
          onChange={e => this.change(e)}
        />
        <br />
        <input
          name="email"
          placeholder="email"
          value={this.state.email}
          onChange={e => this.change(e)}
        />
        <br />
        <input
          name="password"
          type="password"
          placeholder="password"
          value={this.state.password}
          onChange={e => this.change(e)}
        />
        <br />
        <button onClick={e => this.onEdit(e)}> Edit </button>
        <button onClick={e => this.onSubmit(e)}> Save </button>
      </form>
    );
  }
}
