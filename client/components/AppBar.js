import React from 'react'
import styles from '../styles/Home.module.css'
import Link from 'next/link'

export default function AppBar() {
    return (
        <div className={styles.appBar}>
            <div className={styles.userInfo}>
            <Link href="/signin">
                <a>Đăng Nhập</a>
            </Link>
            </div>
        </div>
    )
}
