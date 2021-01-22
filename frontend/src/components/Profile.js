import React from 'react';
import {connect} from 'react-redux';
import {Redirect} from 'react-router';

import * as authActions from '../store/actions/index';

import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import Container from '@material-ui/core/Container';
import {useStyles} from './Styles';
import {withStyles} from '@material-ui/core/styles';
import {borders} from '@material-ui/system';



export const isAuth = () => {
        return localStorage.getItem('username');
}

export const isAuthToken = () => {

    if (localStorage.getItem('token')) {
        return localStorage.getItem('token');
    } else {
        return false;
    }
}
const userName = isAuth();
const token = isAuthToken();
const Profile = props => {

    return (
        <Container maxWidth="sm">
            <Grid>
                <Box border={1} borderRadius={"borderRadius"} borderColor="pink" boxShadow={3}
                     style={{padding: '75px'}}>
                    <h1>Your profile</h1>
                    <table>
                    <tr>
                        <td>userName :</td>
                        <td >{userName}</td>
                    </tr>
                    <tr>
                        <td>token :</td>
                        <td>{token}</td>  
                    </tr>
                    </table>
                </Box>
            </Grid>
        </Container>
    );
}


// get state from reducer
const mapStateToProps = (state) => {
    return {
        token: state.auth.token,
        loading: state.loadingError.loading,
        error: state.loadingError.error,
    };
}

// actions to reducer (dispatch)
const mapDispatchToProps = (dispatch) => {
    return {

    };
}

export default connect(mapStateToProps, mapDispatchToProps)(Profile);
