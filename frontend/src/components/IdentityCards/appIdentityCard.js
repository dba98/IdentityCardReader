import React, {useEffect, useState} from 'react';
import IdentityCard from './IdentityCard'
import WithCardLoading from './WithCardLoading';


function AppIdentityCard() {
  const CardLoading = WithCardLoading(IdentityCard);
  const [appState, setAppState] = useState({
    loading: false,
    cards: [],
  });


  const params={
    nif:localStorage.getItem('nif'),
    image1:'',
    image2:''
  }
  useEffect(() => {
    setAppState({ loading: true });
    const apiUrl = `http://localhost:8082/api/users/getIdentityCardInfo`;
    fetch(apiUrl, {
      method:'POST',
      headers:{
        'Authorization': 'Bearer' + localStorage.getItem('token')
      },
      body: JSON.stringify(params)
       
    })
      .then((res) => res.json())
      .then((cards) => {
        setAppState({ loading: false, cards: cards });
      });
  }, [setAppState]);
  
  return (
    <div className='App'>
      <div className='container'>
        <h1>My Identity Cards</h1>
      </div>
      <div className='repo-container'>
        <CardLoading isLoading={appState.loading} cards={appState.cards} />
      </div>
      <footer>
        <div className='footer'>
        
        </div>
      </footer>
    </div>
  );
}

export default AppIdentityCard;