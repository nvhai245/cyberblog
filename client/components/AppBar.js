import React from 'react'
import styles from '../styles/Home.module.css'
import Link from 'next/link'

export default function AppBar({authState}) {
    return (
        <div className={styles.appBar}>
            <div className={styles.userInfo}>
                {authState && authState.loggedIn ?
                    <div className={styles.userInfo}>
                        <Link href={`${authState.user.username}/profile`}>
                            <a>{"[" + authState.user.username + " user]"}</a>
                        </Link>
                        <Link href="/logout">
                            <a>[Đăng xuất]</a>
                        </Link>
                    </div> :
                    <Link href="/signin">
                        <a>[Đăng nhập]</a>
                    </Link>
                }
            </div>
        </div>
    )
}
