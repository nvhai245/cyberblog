import React, {useEffect, useState} from 'react';
import styles from '../styles/Home.module.css'
import Loading from "../components/Loading";
import Link from "next/link";
import {REGISTER} from '../libs/gql/register'
import Cookie from "js-cookie";
import Router from "next/router";
import {useMutation} from "@apollo/client";

function Register(props) {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [username, setUsername] = useState("")
    const [passwordConfirm, setPasswordConfirm] = useState("")
    const [authInfo, setAuthInfo] = useState({loggedIn: false, user: null})

    const [register, {loading: mutationLoading, error: mutationError}] = useMutation(REGISTER)

    const handleEmailChange = (e) => {
        setEmail(e.target.value)
    }
    const handlePasswordChange = (e) => {
        setPassword(e.target.value)
    }
    const handleUsernameChange = (e) => {
        setUsername(e.target.value)
    }

    const handlePasswordConfirmChange = (e) => {
        setPasswordConfirm(e.target.value)
    }
    const handleSubmit = (e) => {
        e.preventDefault()
        register({
            variables: {
                email: email,
                password: password,
                username: username,
            }
        }).then(({data}) => {
            console.log(data.register)
            if (data.register && data.register.message == "REGISTERED SUCCESSFUL!") {
                setAuthInfo({loggedIn: true, user: data.register.user})
            }
        }).catch(error => {
            console.log(error.toString())
        })
    }
    useEffect(() => {
        const changeLoginState = async () => {
            if (authInfo.loggedIn) {
                await Cookie.set("authInfo", authInfo)
                return Router.push("/")
            }
        }
        changeLoginState()
    }, [authInfo])
    return (
        <div className={styles.main}>
            <h1>Đăng Ký</h1>
            <form onSubmit={handleSubmit} className={styles.inputBox}>
                <div className="editor-field editor-field__textbox">
                    <div className="editor-field__label-container">
                        <p className="editor-field__label">Email:</p>
                    </div>

                    <div className="editor-field__container">
                        <input onChange={handleEmailChange} type="email" className="editor-field__input" required />
                    </div>
                </div>
                <div className="editor-field editor-field__textbox">
                    <div className="editor-field__label-container">
                        <p className="editor-field__label">Username:</p>
                    </div>

                    <div className="editor-field__container">
                        <input onChange={handleUsernameChange} type="text" className="editor-field__input" required />
                    </div>
                </div>
                <div className="editor-field editor-field__textbox">
                    <div className="editor-field__label-container">
                        <p className="editor-field__label">Mật Khẩu:</p>
                    </div>

                    <div className="editor-field__container">
                        <input onChange={handlePasswordChange} type="password" className="editor-field__input" required />
                    </div>
                </div>
                <div className="editor-field editor-field__textbox">
                    <div className="editor-field__label-container">
                        <p className="editor-field__label">Xác nhận mật khẩu:</p>
                    </div>

                    <div className="editor-field__container">
                        <input onChange={handlePasswordConfirmChange} type="password" className="editor-field__input" required />
                    </div>
                </div>
                <div style={{textAlign: "right"}}>
                    {mutationLoading && <Loading/>}
                    {mutationError &&
                    <p style={{color: "#FF9494", fontSize: "1rem", marginTop: "0.5rem"}}>{mutationError.toString()}</p>}
                </div>
                <div onClick={handleSubmit} className={styles.authButtons}>
                    <div style={{marginTop: "0.5rem"}} className="btn btn--primary">
                        <button type="submit" className="text-button">
                            <a href="#">
                                [Đăng Ký]
                            </a>
                        </button>
                    </div>
                </div>
            </form>
        </div>
    );
}

export default Register;