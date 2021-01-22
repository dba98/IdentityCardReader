import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card'; 
import Typography from '@material-ui/core/Typography';

const useStyles = makeStyles({
    root: {
      minWidth: 275,
      maxWidth: 300,
    },
    bullet: {
      display: 'inline-block',
      margin: '0 2px',
      transform: 'scale(0.8)',
    },
    title: {
      fontSize: 14,
    },
    pos: {
      marginBottom: 12,
    },
  });



const IdentityCard = () =>{
    const classes = useStyles();
    const cardInfo = [
        {title: "joao maneta", text: "funciona pls", text2 :"bruto"},
        {title: "anacleto mongo", text: "era fixe se funcionasses", text2 :"grande e grosso"},
        {title: "rudoflo estroncio", text: "deixa me passar á cadeira", text2 :"manooobro"},
    ];

    const renderCard = (card,index) => {
        return(
            <Card className={classes.root} variant="outlined" key={index}>
            <CardContent>
              <Typography className={classes.title} color="textSecondary" gutterBottom>
                {card.title}
              </Typography>
              <Typography className={classes.text} color="textSecondary">
                  {card.text}
              </Typography>
              <Typography variant="body2" component="p">
              {card.text2}
              </Typography>
            </CardContent>
          </Card>
        )
    }
    return <div className="grid">
        {cardInfo.map(renderCard)}
        </div>
};


export default IdentityCard;


