import React from 'react';

const Profile = props => {

    return (
        <tr>
            <td>{props.id}</td>
            <td>{props.username}</td>
        </tr>
    );
}

export default Profile;
