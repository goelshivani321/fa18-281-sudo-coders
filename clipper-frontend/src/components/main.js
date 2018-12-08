import React, {Component} from 'react';
import {Switch,Route} from 'react-router-dom';

//Components
import Login from './login/login.js'; 
import Navbar from "./Navbar/navbar"
import AddPayment from './AddPayment/index.jsx';
//import Search from './search/Search.js';
//import Signup from './signup/signup.js';


class Main extends Component {
    render(){
        return(
          <div>
                {/*Render Different Component based on Route*/}
                {/*<Route exact path="/checklist" render={()=>(<Checklist />) } />*/}
                <Route  path="/" component={Navbar}/>
                <Route exact path="/login" component={Login} />
                <Route exact path="/addpayment" component={AddPayment} />
          </div>
        
        )
    }
}
export default Main;