import React, {Component} from "react";
import {Link, Redirect} from "react-router-dom";
import axios from "axios"
import Navbar from "../Navbar/navbar";

class AddPayment extends Component {

    constructor() {
        super();
        this.state = {
            cardNumber: "",
            expiryMonth: 0,
            expiryYear: 0,
            cvv: 0,
            amount: 0,
            cardType: ""
        }
    }

    cardNumberChangeHandler = e => {
        this.setState({
            cardNumber: e.target.value
        })
    }
    expiryMonthChangeHandler = e => {
        this.setState({
            expiryMonth: e.target.value
        })
    }
    expiryYearChangeHandler = e => {
        this.setState({
            expiryYear: e.target.value
        })
    }
    cvvChangeHandler = e => {
        this.setState({
            cvv: e.target.value
        })
    }
    amountChangeHandler = e => {
        this.setState({
            amount: e.target.value
        })
    }
    cardTypeChangeHandler = e => {
        this.setState({
            cardType: e.target.value
        })
    }
    submitPayment = e => {
        let data = {
            cardNumber: this.state.cardNumber ,
            expiryMonth: this.state.expiryMonth,
            expiryYear: this.state.expiryYear,
            cvv: this.state.cvv,
            amount: this.state.amount,
            cardType: this.state.cardType
        }
        axios.post(`http://localhost:3000/payments/cardId/2/payment`, data).then(response=>{
            console.log(response);
        }).catch(err=>console.log(err))
    }
    

    render(){
        return(
        <React.Fragment>
            <div className="container">
                <h3>Add Payment Method</h3>
                <div className="form-group input-group-md"><input onChange={this.cardNumberChangeHandler} placeholder="Card Number" type="number" className="form-control" /></div>
                <div className="form-group input-group-md"><input onChange={this.expiryMonthChangeHandler} placeholder="Expiry Month" type="number" className="form-control" /></div>
                <div className="form-group input-group-md"><input onChange={this.expiryYearChangeHandler} placeholder="Expiry Year" type="number" className="form-control" /></div>
                <div className="form-group input-group-md"><input onChange={this.cvvChangeHandler} placeholder="CVV" type="number" className="form-control" /></div>
                <div className="form-group input-group-md"><input onChange={this.amountChangeHandler} placeholder="Amount" type="number" className="form-control" /></div>
                <div className="form-group input-group-md"><input onChange={this.cardTypeChangeHandler} placeholder="Card Type" type="text" className="form-control" /></div>
                <div ><button className="btn btn-primary" onClick={this.submitPayment}>Pay</button></div>
            </div>
        </React.Fragment>
        )
    }
}

export default AddPayment;