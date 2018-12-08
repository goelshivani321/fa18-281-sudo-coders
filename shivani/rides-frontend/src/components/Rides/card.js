import React,{ Component } from 'react';
import { Route, withRouter } from 'react-router-dom';
import {Link} from 'react-router-dom';
import './card.css';
import 'tachyons';

class Card extends Component {
  constructor(props) {
    super(props);
    this.state = {
        detailsofProperty:'',
    }
  }
  render() {
    return(
          <div className="card-container">
            <div className='bg-light-orange dib br3 pa3 ma2 grow bw2 shadow-5' style={{backgroundColor: this.props.ride.bgColor}}>
                <table className="card-data">
                  <tr>
                    <td> <h4 className="bold-heading-1"> Customer ID :  {this.props.ride.customerID}</h4></td>
                  </tr>
                  <tr>
                    <td> <h4 className="bold-heading-1"> Ride ID :  {this.props.ride.id}</h4></td>
                  </tr>
                  <tr>
                    <td> <h4 className="bold-heading-1"> Start Time:  {this.props.ride.startTime}</h4></td>
                  </tr>
                  <tr>
                    <td> <h4 className="bold-heading-1"> Start Location:  {this.props.ride.startLocation}</h4></td>
                  </tr>
                  <tr>
                    <td> <h4 className="bold-heading-1"> Ride Type: {this.props.ride.rideType}</h4></td>
                  </tr>
                  <tr>
                    <td> <h4 className="bold-heading-1"> Live Status: {this.props.ride.liveStatus}</h4></td>
                  </tr>
               </table>
        </div>
    </div>

      );
  }
}

export default Card;
