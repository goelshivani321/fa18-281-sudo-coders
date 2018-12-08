import React, {Component} from "react";
import {Link, Redirect} from "react-router-dom";
import axios from "axios"
import Navbar from "../Navbar/navbar";
import './index.css';
class RideDetails extends Component {
    constructor(props) {
        super(props);
        this.state={
            customerId:"1001"
        }
    }

    sourceChangeHandler = e => {
        this.setState({
            customerId: e.target.value
        })
    }

    showRides = () => {
        window.location.assign('/customerrides?customerId=' + this.state.customerId)
    }

    createRide = () => {
        window.location.assign('/createride?customerId=' + this.state.customerId)
    }

    render(){
        return(
        <React.Fragment>
            <div className="container">
                <br></br>
                <h3>Rides</h3>
                <br></br>
                <div className="ride-card1">
                <div className='bg-dark-orange dib br10 pa3 ma2 grow bw10 shadow-2'>
                  <div className="form-group input-group-md ride-card"><label className="label-class"> Enter Customer ID: </label></div>
                  <div className="form-group input-group-md ride-card"><input placeholder="shivani.mangal@sjsu.edu" onChange={this.sourceChangeHandler} type="text" className="form-control" /></div>
                  <br></br>
                  <div >
                    <button type="button" className="btn-lg btn-primary showride" onClick={this.showRides}>Show Rides</button>
                    <button type="button" className="btn-lg btn-primary showride" onClick={this.createRide}>Purchase Ride</button>
                  </div>
                </div>
                </div>
            </div>
        </React.Fragment>
        )
    }
}

export default RideDetails;
