import React, { useState, useEffect } from 'react'
import styles from '../styles/Home.module.css'
import Link from 'next/link'
import { LOGIN } from '../libs/gql/login'
import { gql, useMutation, useQuery } from '@apollo/client'
import Cookie from 'js-cookie'
import Router from 'next/router'
import Loading from '../components/Loading'

export default function SignIn() {
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [authInfo, setAuthInfo] = useState({loggedIn: false, user: null})
    const handleEmailChange = (e) => {
        setEmail(e.target.value)
    }
    const handlePasswordChange = (e) => {
        setPassword(e.target.value)
    }

    const [login, { loading: mutationLoading, error: mutationError }] = useMutation(LOGIN)

    const handleSubmit = (e) => {
        e.preventDefault()
        login({
            variables: {
                email: email,
                password: password
            }
        }).then(({data}) => {
            console.log(data.login)
            if (data.login && data.login.message == "LOGGED IN!") {
                setAuthInfo({loggedIn: true, user: data.login.user})
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
            <h1>Đăng Nhập</h1>
            <form onSubmit={handleSubmit} className={styles.inputBox}>
                <div className="editor-field editor-field__textbox">
                    <div className="editor-field__label-container">
                        <p className="editor-field__label">Email:</p>
                    </div>

                    <div className="editor-field__container">
                        <input onChange={handleEmailChange} type="email" className="editor-field__input" />
                    </div>
                </div>
                <div className="editor-field editor-field__textbox">
                    <div className="editor-field__label-container">
                        <p className="editor-field__label">Mật Khẩu:</p>
                    </div>

                    <div className="editor-field__container">
                        <input onChange={handlePasswordChange} type="password" className="editor-field__input" />
                    </div>
                </div>
                <div style={{textAlign: "right"}}>
                    {mutationLoading && <Loading />}
                    {mutationError && <p style={{color: "#FF9494", fontSize: "1rem", marginTop: "0.5rem"}}>{mutationError.toString()}</p>}
                </div>
                <div onClick={handleSubmit} className={styles.authButtons}>
                    <div style={{ marginTop: "0.5rem" }} className="btn btn--primary">
                        <button type="submit" className="text-button">
                            <a href="#">
                                [Đăng Nhập]
                            </a>
                        </button>
                    </div>
                </div>
                <p>
                    Bạn chưa có tài khoản? <Link href="/register"><a style={{ textDecoration: "underline" }}>Đăng ký ngay</a></Link>
                </p>
            </form>
        </div>
    )
}
