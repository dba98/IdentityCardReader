import axios from 'axios';
import jwt_decode from "jwt-decode";
import * as actionTypes from './actionTypes';
import * as api from '../actions/api';
import * as loadingErrorActions from './loadingError';


const getAllIdentityCards = (identityCards) => {
    return {
        type: actionTypes.GET_ALL_IDENTITYCARDS,
        identityCards: identityCards
    }
}

export const fetchAllIdentityCards = (token, id) => {
    return (dispatch) => {
        dispatch(loadingErrorActions.startRequest());
        const auth = {
            headers: {
                Authorization: token
            }
        };
        axios.get(api.URL_GET_ALL_IDENTITYCARDS + "/" + id, auth).then(response => {
            dispatch(getAllIdentityCards(response.data.data));
            dispatch(loadingErrorActions.endRequest());
        }).catch(err => {
            console.log(err)
            dispatch(loadingErrorActions.errorRequest(err.toString()));
        });
    }
}

const addIdentityCard = (identityCard, id) => {
    return {
        type: actionTypes.CREATE_IDENTITYCARD,
        identityCard: identityCard,
        userId: id,
    }
}

export const createIdentityCard = (identityCard, token) => {
    return (dispatch) => {
        const auth = {
            headers: {
                Authorization: token
            }
        };
        axios.post(api.URL_IDENTITYCARD_ADD, identityCard, auth).then(res => {
            dispatch(addIdentityCard(identityCard, jwt_decode(token).id));
        }).catch(err => {
            console.log(err);
        });
    }
}

const onDeleteIdentityCard = (id, userId) => {
    return {
        type: actionTypes.DELETE_IDENTITYCARD,
        id: id,
        userId: userId
    }
}

export const deleteIdentityCard = (id, token) => {
    return (dispatch) => {
        const auth = {
            headers: {
                Authorization: token
            }
        };
        axios.delete(api.URL_IDENTITYCARD_DELETE + id, auth).then(res => {
            dispatch(onDeleteIdentityCard(id, jwt_decode(token).id));
        }).catch(err => {
            console.log(err);
        });
    }
}
