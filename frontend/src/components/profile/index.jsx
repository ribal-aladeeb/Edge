import React, { Component } from "react";
import Home from "../Home";
import UploadImages from "./uploadImages";
import decode from "jwt-decode";
import "./styles/profile.css";

export default class Profile extends Component {
  constructor(props) {
    super(props);
    this.infos = new Home();
    this.state = {
      userProfile: [],
      firstName: "",
      lastName: "",
      email: "",
      password: "",
      isTeacher: false
    };
  }

  decodeJwtToken() {
    try {
      const profile = this.getProfile();
      this.setState({
        userProfile: profile
      });
    } catch (err) {
      localStorage.removeItem("jwt"); //if an error occurs while decoding jwt token, logout
      this.props.history.replace("/login");
    }
  }

  getToken() {
    // Retrieves the user token jwt from localStorage
    return localStorage.getItem("jwt");
  }

  getProfile() {
    // Using jwt-decode npm package to decode the token
    return decode(this.getToken());
  }

  getTeacher() {
    if (this.state.userProfile.is_teacher === 0)
      //if it is not a teacher
      this.setState({
        isTeacher: false
      });
    else
      this.setState({
        isTeacher: true
      });
  }

  componentDidMount() {
    let jwt = localStorage.getItem("jwt");
    if (jwt === undefined || jwt === null) {
      //if the user not logged in
      this.props.history.replace("/login"); //go login
    } else {
      // if is logged in, get user profile
      this.decodeJwtToken();
    }
  }

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
      firstName: e.target.value,
      lastName: "",
      email: "",
      password: ""
    });
  };

  render() {
    const labelTeacher = this.state.isTeacher ? "Parent" : "Teacher";
    return (
      <React.Fragment>
        <center>
          <UploadImages />
        </center>
        <br />
        <div className="profilecontainer">
          <img
            src={require("./images/banner.jpg")}
            alt="Welcome"
            className="banner"
          />
          <center>
            <img
              src={require("./images/profile.png")}
              alt="profile"
              className="pp"
            />
          </center>
          <h3>
            {this.state.userProfile.firstname} {this.state.userProfile.lastname}
          </h3>
          <button className="editpic">Update Info</button>
          <br />
          <br />
          <br />
          <br />
        </div>
        <h6> This user is a {labelTeacher}</h6>
        <div className="profile">
          <form>
            <p>First Name</p>
            <input
              className=""
              name="firstName"
              placeholder="First Name"
              value={this.state.userProfile.firstname}
              onChange={e => this.change(e)}
            />
            <br />
            <br />
            <p>Last Name</p>
            <input
              className=""
              name="lastName"
              placeholder="Last Name"
              value={this.state.userProfile.lastname}
              onChange={e => this.change(e)}
            />
            <br />
            <br />
            <p>Email</p>
            <input
              className=""
              name="email"
              placeholder="Email"
              value={this.state.userProfile.email}
              onChange={e => this.change(e)}
            />
            <br />
            <br />
            <p>Password</p>
            <input
              className=""
              name="password"
              type="password"
              placeholder="Password"
              value={this.state.userProfile.password}
              onChange={e => this.change(e)}
            />
            <br />

            <div className="App">
              <input
                type="file"
                name=""
                id=""
                onChange={this.handleselectedFile}
              />
              <button onClick={this.handleUpload}>Upload</button>
              <div> {Math.round(this.state.loaded, 2)} %</div>
            </div>
            <br />
            <button onClick={e => this.onEdit(e)}> Edit </button>
            <button onClick={e => this.onSubmit(e)}> Save </button>
          </form>
        </div>
      </React.Fragment>
    );
  }
}