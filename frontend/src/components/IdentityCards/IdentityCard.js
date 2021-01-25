import React, { useState, useEffect } from 'react';
import { connect } from 'react-redux';
import { Redirect } from 'react-router';
import GridList from "@material-ui/core/GridList";

import Grid from '@material-ui/core/Grid';
import Box from '@material-ui/core/Box';
import Container from '@material-ui/core/Container';
import Button from '@material-ui/core/Button';
import { useStyles } from '../Styles';
import * as actions from '../../store/actions/index';
import jwt_decode from 'jwt-decode';

const IdentityCard = (props) => {
  const { cards } = props;
 const  data= Array.from(props.cards)
  if (!cards || cards.length === 0) return <p>No cards, sorry</p>;
  return (
    <ul>
      <h2 className='list-head'>Available Identity cards</h2>
      {data.map((card) => {
        return (
          <li key={card.nif} className='list'>
            <span className='repo-text'>{card.frontdata} </span>
            <span className='repo-description'>{card.backdata}</span>
          </li>
        );
      })}
    </ul>
  );
};
export default IdentityCard; 