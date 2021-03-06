import Profile from "./components/Profile";
import Home from "./components/Home";
import Login from "./components/Login";
import SignUp from "./components/SignUp";
import Logout from "./components/Logout";
import IdentityCard from "./components/IdentityCards/IdentityCard"
import React from 'react';
import {BrowserRouter, Redirect, Route, Switch} from 'react-router-dom';
import './App.css';
import Layout from './components/Layout';
import {connect} from 'react-redux';
import AppIdentityCard from "./components/IdentityCards/AppIdentityCard";




const App = props => {


    const routes = (
        <Switch>
            <Route path="/home" component={Home}/>
            <Route path="/signup" component={SignUp}/>
            <Route path="/logout" component={Logout}/>
            <Route path="/profile" component={Profile}/>
            <Route path="/login" component={Login}/>
            <Route path="/identitycards" component={AppIdentityCard}/>
            <Route render={() => <h1>Not found!</h1>}/>
        </Switch>
    );

    return (
        <React.Fragment>
            <BrowserRouter>
                <div className="App">
                    <Layout>
                        {routes}
                    </Layout>
                </div>
            </BrowserRouter>
        </React.Fragment>
    );

}

// get state from reducer
const mapStateToProps = (state) => {
    return {
        token: state.auth.token,
    };
}



export default connect(mapStateToProps, null)(App);
