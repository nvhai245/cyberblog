import React, {useEffect} from 'react'
import Router from 'next/router'
import Cookie from "js-cookie";

const Logout = () => {
    useEffect(() => {
        const signOut = async () => {
            await Cookie.remove("authInfo")
            Router.push("/")
        }
        signOut()
    }, [])
    return null
}

export default Logout