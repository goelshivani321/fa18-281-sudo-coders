import React, { Component } from "react";
import { Link } from "react-router-dom";

//create the Navbar Component
class Navbar extends Component {
  constructor() {
    super();
    this.userLogout = this.userLogout.bind(this);
  }
  userLogout() {
    sessionStorage.removeItem("email");
    this.forceUpdate();
  }
  render() {
    let email = sessionStorage.getItem("email");
    return (
      <div>
        <nav id="homepage-nav" className="navbar navbar-default">
          <div className="container-fluid">
            <div className="navbar-header">
              <Link to="/" alt="homepage">
                <h2>VTA Clipper Card</h2>
              </Link>
            </div>
            
              <ul className="nav navbar-right">
              <li className="nav-item">
                <Link to="/users">
                  <button
                    id="login-btn"
                    className="btn btn-primary"
                    style={{ color: "white" }}
                >
                    Users
                  </button>
                </Link>
             </li>
              <li className="nav-item">
                <Link to="/cards">
                  <button
                    id="login-btn"
                    className="btn btn-primary"
                    style={{ color: "white" }}
                    >
                    Cards
                  </button>
                </Link>
                    </li>
              <li className="nav-item">
                <Link to="/rides">
                  <button
                    id="login-btn"
                    className="btn btn-primary"
                    style={{ color: "white" }}
                    >
                    Rides
                  </button>
                </Link>
                    </li>
              <li className="nav-item">
                <Link to="/addpayment">
                  <button
                    id="login-btn"
                    className="btn btn-primary"
                    style={{ color: "white" }}
                    >
                    Payment
                  </button>
                </Link>
                    </li>
              <li className="nav-item">
                <Link to="/login">
                  <button
                    id="login-btn"
                    className="btn btn-primary"
                    style={{ color: "white" }}
                    >
                    Login
                  </button>
                </Link>
                    </li>
              </ul>
          </div>
        </nav>
      </div>
    );
  }
}

export default Navbar;
