import React, {Component} from "react";
import {Link, Redirect} from "react-router-dom";
import axios from "axios"
import Navbar from "../Navbar/navbar";
import './index.css';

const queryString = require('query-string');

class CreateRide extends Component {
  constructor(props) {
    super(props);
    this.state={
      source:"",
        customerId:"1001",
        status:""
    }
  }

    sourceChangeHandler = e => {
        this.setState({
            source: e.target.value
        })
    }

    makePayment = () => {
        let parsed = queryString.parse(this.props.location.search);
        console.log(parsed.customerId);

        let host = "http://ec2-52-9-131-191.us-west-1.compute.amazonaws.com"
        let port = 3000
        let data = {
            "customerID" : parsed.customerId,
            "startLocation" : this.state.source,
            "rideType" : "Light Rail"
        }
        axios.post(`${host}:${port}/rides`, data).then(response=>{
            console.log(response)
            this.setState({
                status: "Success!"
            })
        }).catch(err=>console.log(err))
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
                  <div className="form-group input-group-md ride-card"><label className="label-class"> Enter Source: </label></div>
                  <div className="form-group input-group-md ride-card"><input onChange={this.sourceChangeHandler} placeholder="Tasman Drive" type="text" className="form-control" /></div>
                  <br></br>
                  <div >

                    <button onClick={this.makePayment} className="btn-lg btn-primary showride" >Make Payment</button>
                      <Link to='/rides' ><button type="button" className="btn-lg btn-primary returnmain">Return back to rides page</button></Link>
                  </div>
                    <div><label id="status">{this.state.status}</label></div>
                </div>
                </div>
            </div>
        </React.Fragment>
        )
    }
}

export default CreateRide;
