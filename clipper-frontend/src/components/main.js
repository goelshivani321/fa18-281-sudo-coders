import React, {Component} from 'react';
import {Switch,Route} from 'react-router-dom';

//Components
import Login from './login/login.js'; 
//import Search from './search/Search.js';
//import Signup from './signup/signup.js';


class Main extends Component {
    render(){
        return(
                        <Switch>
                {/*Render Different Component based on Route*/}
                {/*<Route exact path="/checklist" render={()=>(<Checklist />) } />*/}
                <Route exact path="/" component={Login}/>
                <Route exact path="/login" render={Login} />
                
            </Switch>
        
        )
    }
}
export default Main;