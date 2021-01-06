import * as actionTypes from '../actions/actionTypes';

const initialState = {
    users: [],
    userIdentityCards: [],
};

const reducer = (state = initialState, action) => {
    switch (action.type) {
        case (actionTypes.GET_ALL_USERS):
            return {
                ...state,
                users: action.users,
            };
        case (actionTypes.CREATE_USER):
            const newUser = {
                ID: action.id,
                username: action.username,
            }
            return {
                ...state,
                users: state.users.concat(newUser),
            };
           // let userarray = state.users.filter((user) => (user.ID !== action.user.id));
            //let newarray = userarray.concat(updatedUser);
            /*return {
                ...state,
                users: newarray
            };*/
        /*case (actionTypes.DELETE_USER):
            return {
                ...state,
                users: state.users.filter((user) => (user.ID !== action.id)),
            };
        case (actionTypes.GET_USER_PLACES_ADMIN):
            return {
                ...state,
                userPlaces: action.userPlaces,
            };
        case (actionTypes.ADD_PLACE_TO_USER):
            return {
                ...state,
                userPlaces: state.userPlaces.concat(action.place),
            };
        case (actionTypes.REMOVE_PLACE_FROM_USER):
            return {
                ...state,
                userPlaces: state.userPlaces.filter((place) => (place.ID !== action.id)),
            };*/
        default:
            return state;
    }
}

export default reducer;
