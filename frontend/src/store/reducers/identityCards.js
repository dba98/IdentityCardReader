import * as actionTypes from '../actions/actionTypes';

const initialState = {
    identityCards: [],
};

const reducer = (state = initialState, action) => {

    switch (action.type) {

        /*case (actionTypes.ADD_SOCKET_LIST):
            console.log(action)
            let newArrayPlacesList = state.places.filter((place) => (place.ID !== action.id));
            let placeList = state.places.find(place => place.ID === action.id);

            placeList = { ...placeList, users: action.users };
            newArrayPlacesList = newArrayPlacesList.concat(placeList);
            
            return {
                ...state,
                places: newArrayPlacesList,
            };
        case (actionTypes.ADD_SOCKET_PEOPLE):
            console.log(action)
            let newArrayPlaces = state.places.filter((place) => (place.ID !== action.id));
            let place = state.places.find(place => place.ID === action.id);
            place = { ...place, people: action.numUsers };
            newArrayPlaces = newArrayPlaces.concat(place);
            return {
                ...state,
                places: newArrayPlaces,
            };
*/
            ////////////

        case (actionTypes.GET_ALL_IDENTITYCARDS):
            return {
                ...state,
                identityCards: action.identityCards,
            };
        case (actionTypes.CREATE_IDENTITYCARD):
            const newIdentityCard = {
                ...action.identityCard,
                ID: action.id,
                image: ''
            }
            return {
                ...state,
                identityCards: state.identityCards.concat(newIdentityCard),
            };
        case (actionTypes.DELETE_IDENTITYCARD):
            return {
                ...state,
                identityCards: state.identityCards.filter((identityCard) => (identityCard.ID !== action.id)),
            };
        default:
            return state;
    }
}

export default reducer;
