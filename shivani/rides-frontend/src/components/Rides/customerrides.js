import React, {Component} from "react";
import {Link, Redirect} from "react-router-dom";
import axios from "axios"
import Navbar from "../Navbar/navbar";
import './index.css';
import Card from './card.js'
const queryString = require('query-string');

class CustomerRideDetails extends Component {
  constructor(props) {
    super(props);
    this.state={
      rideDetails:[]
    }
  }
  

    render(){
        return(
        <React.Fragment>
          <div className="carddetails">
            {
                this.state.rideDetails != undefined ?
                  this.state.rideDetails.map((ride) => {
                    if (ride.liveStatus == "live"){
                      ride.bgColor = "lightgreen"
                    } else {
                      ride.bgColor = "antiquewhite"
                    }
                   return(<Card ride={ride}/>);
                  })
                  :  'No Rides for this User!'
            }
            <Link to='/rides' ><button type="button" className="btn-lg btn-primary returnmain">Return back to rides page</button></Link>
            </div>
        </React.Fragment>
        )
    }

    componentDidMount(){
      let parsed = queryString.parse(this.props.location.search);
      console.log(parsed.customerId);

      let host ="http://ec2-52-9-131-191.us-west-1.compute.amazonaws.com";
      let port = 3000;
      // let userid = sessionStorage.getItem("userid")
       let customerId = parsed.customerId
      axios.get(`${host}:${port}/rides/${customerId}`).then(response=> {
        console.log(response);
        this.setState({rideDetails: response.data})
    }).catch(err=>console.log(err))
    }
}

export default CustomerRideDetails;
