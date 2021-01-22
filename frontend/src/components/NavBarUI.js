import React from 'react';

import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';
import {NavLink, useHistory} from 'react-router-dom';
import AppBar from "@material-ui/core/AppBar";
import { makeStyles } from '@material-ui/core/styles';
import Toolbar from "@material-ui/core/Toolbar";
import IconButton from "@material-ui/core/IconButton";
import Typography from "@material-ui/core/Typography";
import MenuIcon from '@material-ui/icons/Menu';
import Button from "@material-ui/core/Button";
import Profile from "./Profile";

const useStyles = makeStyles((theme) => ({
    root: {
        flexGrow: 1,
    },
    menuButton: {
        marginRight: theme.spacing(2),
    },
    title: {
        flexGrow: 1,
        alignSelf: 'center',
        color: 'white'
    },
}));

const NavbarUI = props => {
    const classes = useStyles();
    const history = useHistory();
    function onProfileButtonClick() {
        history.push("/profile");
    }

    function onIdentityCardsButtonClick() {
        history.push("/identitycards");

    }

    return (
        /*<Navbar collapseOnSelect expand="lg" bg="dark" variant="dark">
            <Navbar.Toggle aria-controls="responsive-navbar-nav"/>
            <Navbar.Collapse id="responsive-navbar-nav">
                <Nav className="mr-auto">
                    <Nav.Link as={NavLink} to="/home">Home</Nav.Link>
                    <Nav.Link as={NavLink} to="/identityCards">Identity Cards</Nav.Link>
                    <Nav.Link as={NavLink} to="/profile">Profile</Nav.Link>
                </Nav>
                <Nav>
                    <Nav.Link as={NavLink} to="/logout">Logout</Nav.Link>
                </Nav>
            </Navbar.Collapse>
        </Navbar>
        */
        <div className={classes.root}>
        <AppBar position="static">
            <Toolbar>
                <IconButton edge="start" className={classes.menuButton} color="inherit" aria-label="menu">
                    <MenuIcon />
                </IconButton>
                <Button color="black" className={classes.title} onClick={onIdentityCardsButtonClick}>IdentityCards</Button>
               <Button color="inherit" className={classes.menuButton} onClick={onProfileButtonClick}>Profile</Button>
            </Toolbar>
        </AppBar>
        </div>
    );
}

export default NavbarUI;
