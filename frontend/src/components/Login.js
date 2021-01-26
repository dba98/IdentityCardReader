import React, { useState, useEffect } from 'react';
import { connect } from 'react-redux';
import { Redirect } from 'react-router';

import * as authActions from '../store/actions/index';

import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import Container from '@material-ui/core/Container';
import { useStyles } from './Styles';

const Login = props => {

    // styles
    const classes = {
        root: {
            background: "black"
        },
        input: {
            color: "white"
        }
    };

    // state
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    // dependency
    const { onTryAutoLogin } = props;

    useEffect(() => {
        onTryAutoLogin();
    }, [onTryAutoLogin])

    const onSubmitHandler = (event) => {
        event.preventDefault();
        props.onAuth(username, password);
    }

    let error = props.error ? <label style={{ color: 'red' }}>Login failed!</label> : null;

    let isAuth = props.token ? <Redirect to='/home' /> : null;

    return (
        <Container maxWidth="sm" >
            <Grid>
                <Box border={1} borderRadius={"borderRadius"} borderColor="pink" boxShadow={3} style={{ padding: '75px' }}>
                    {isAuth}
                    <form onSubmit={onSubmitHandler} className={classes.authTextFileds} noValidate autoComplete="off">
                        <h1>Login</h1>
                        <div>
                            <TextField required id="standard-basic" label="Username" variant="outlined"  size="small" color="secondary"
                                       InputProps={{
                                           classes: {
                                               className: classes.input
                                           }
                                       }} onChange={event => {
                                setUsername(event.target.value);
                            }} />
                        </div>
                        <div>
                            <TextField required id="standard-password-input" label="Password" variant="outlined" size ="small" color= "secondary"type="password" onChange={event => {
                                setPassword(event.target.value);
                            }} />
                        </div>
                        {error}
                        <div>
                            <Button type="submit" variant="contained" color="primary">Login</Button>
                        </div>
                    </form>
                </Box>
            </Grid>
        </Container >
    );
}

// get state from reducer
const mapStateToProps = (state) => {
    return {
        token: state.auth.token,
        isAdmin: state.auth.isAdmin,
        loading: state.loadingError.loading,
        error: state.loadingError.error,
    };
}

// actions to reducer (dispatch)
const mapDispatchToProps = (dispatch) => {
    return {
        onAuth: (username, password) => dispatch(authActions.auth(username, password)),
        onTryAutoLogin: () => dispatch(authActions.authCheckState()),
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(Login);
