import React from 'react'
import styles from '../../styles/Home.module.css'
import Link from 'next/link'

export default function yolo() {
    return (
        <div className={styles.card}>
            <h1>
                <Link href="/"><a>Home</a></Link>
            </h1>
        </div>
    )
}
