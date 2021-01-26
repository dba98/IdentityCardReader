import React, {useEffect, useState} from 'react';
import IdentityCard from './IdentityCard';
import axios from 'axios';
import List from '@material-ui/core/List'
import ListItem from '@material-ui/core/ListItem'
import { withStyles, makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

const classes = makeStyles(theme => ({
  identityCards: {
    width: '100%',
    maxWidth: '80%',
    margin: 'auto',
    backgroundColor: theme.palette.background.paper,
},
backdrop: {
    zIndex: theme.zIndex.drawer + 1,
    color: '#fff',
},
authTextFileds: {
    '& .MuiTextField-root': {
        margin: theme.spacing(1),
        width: '25ch',
    },
},
}))

const StyledTableCell = withStyles((theme) => ({
  head: {
    backgroundColor: theme.palette.common.black,
    color: theme.palette.common.white,
  },
  body: {
    fontSize: 14,
  },
}))(TableCell);

const StyledTableRow = withStyles((theme) => ({
  root: {
    '&:nth-of-type(odd)': {
      backgroundColor: theme.palette.action.hover,
    },
  },
}))(TableRow);


const useStyles = makeStyles({
  table: {
    minWidth: 700,
  },
});


 export default class AppIdentityCard extends React.Component{
  state = {
    cards: []

  }

   componentDidMount =() => {
    axios({
      method:'GET',
      url: 'http://localhost:8082/api/users/getAllIdentityCardInfo',
      headers:{
        'Authorization': 'Bearer' + localStorage.getItem('token')
      }
    })
    .then((res)=> {
      console.log(res.data)
      const cards = res.data.data
      this.setState({
        cards
      });
    })
    .catch(error => console.error('Error: ${error}'));
    
  }


    handleRemove(id) {
      const params={
        nif:id,
        image1: '',
        image2: ''
      }
    axios({
      method:'DELETE',
      url: 'http://localhost:8082/api/users/deleteIdentityCardInfo',
      headers:{
        'Authorization': 'Bearer' + localStorage.getItem('token')
      },
      data: JSON.stringify(params)
    })
    const cards = this.state.cards
    var posicao = -1;
    cards.forEach((element,index) => {
      if(element.nif==id){
        posicao= index;
        return;
      }
    });
    cards.splice(posicao,1);
    console.log(cards);
    this.setState({
      cards
    })
  }

  
  render(){
    return(
      <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="customized table">
        <TableHead>
          <TableRow>
            <StyledTableCell>nif</StyledTableCell>
            <StyledTableCell align="right">FrontData</StyledTableCell>
            <StyledTableCell align="right">BackData</StyledTableCell>
            <StyledTableCell align="right"></StyledTableCell>
          </TableRow>
        </TableHead>
        <TableBody>
        {this.state.cards.map((card) => (
            <StyledTableRow key={card.nif}>
              <StyledTableCell component="th" scope="row">
                {card.nif}
              </StyledTableCell>
              <StyledTableCell align="right">{card.image1}</StyledTableCell>
              <StyledTableCell align="right">{card.image2}</StyledTableCell>
              <StyledTableCell align="right">
                <button onClick={()=> this.handleRemove(card.nif)}>
                  Delete
                </button>
              </StyledTableCell>
            </StyledTableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
  }
   


