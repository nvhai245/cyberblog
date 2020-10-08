import React from 'react';
import {useRouter} from "next/router";

function Profile(props) {
    const router = useRouter()
    console.log(router.query);
    return (
        <div></div>
    );
}

export default Profile;