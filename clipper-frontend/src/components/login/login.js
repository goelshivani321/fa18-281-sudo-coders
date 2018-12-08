import React, { Component } from 'react';
import {Link} from 'react-router-dom';
import serialize from 'form-serialize';
import axios from 'axios';

class Login extends Component {
    
    constructor(props){
        super(props)

    }


    onFormSubmit = (e)=>{
        e.preventDefault();
        var form = serialize(e.target, { hash: true });
        const userdata = {
            username:form.username,
            password:form.password 
        }
        
        axios.post("http://localhost:3000/ping",userdata).then((response)=>{
            console.log(response.status);
            console.log(response.data);
        }).catch((error)=>{
            console.log(error);
        })
    }

    render() { 
        return (
            <div>
                
            <div className="row justify-content-center w-100" style={{backgroundColor:'#F4F4F4',marginTop:'5%',marginBottom:'10%'}}>
                <div className="col-md-4 add-border-signup text-center" style={{backgroundColor:'white'}}>
                    <h1 className="homeaway-h1 justify-content-centre mb-2" style={{fontSize:'38px'}} >Login into VTA Clipper</h1>
                    
                    
                    <form onSubmit={this.onFormSubmit}>
                    <div className="mt-4" style={{border:'0px 2px 0px 2px'}}>
                        <input type="text" name="username"  className="width-100" placeholder="Username"/>
                    </div>
        
                    <div className="mt-4" style={{border:'0px 2px 0px 2px'}}>
                        <input type="password" name="password"  placeholder="Password" className="width-100" />
                    </div>
        
                    <div className="mt-4">
                        <button type="submit" className="btn btn-primary btn-block" style={{color:'white',height:'40px',fontSize:'22px'}}>Login</button> 
                    </div>

                    <div className="mt-4">
                        <Link type="submit" to="/signup" className="btn btn-primary btn-block" style={{color:'white',height:'40px',fontSize:'22px'}}>Sign Up</Link> 
                    </div>
                    </form>
                </div>
            </div>
        </div>
        );

    }
}
 
export default Login;