import React, { Component } from 'react';
import { Link } from 'react-router-dom';

class HomeAwayPlainNavBar extends Component {
    
    constructor(props){
        super(props);
    }

    render() { 

        return (
            <div className="row w-100 add-border-signup pb-3" style={{ backgroundColor: '#FFFFFF', height: '80px', borderTopStyle: 'none', borderLeftStyle: 'none', borderRightStyle: 'none' }}>
                <div className="col-md-1"></div>
                <div className="col-md-2">
                    <Link to="/"> <p>VTA Clipper System</p> </Link>
                </div>
                <div className="col-md-6"></div>
                <div className="col-md-1">
                    
                </div>
                <div className="col-md-1"></div>
            </div>
          );
    }
}
 
export default HomeAwayPlainNavBar;