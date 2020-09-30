import Head from 'next/head'
import styles from '../styles/Home.module.css'
import Link from 'next/link'
import AppBar from '../components/AppBar'

export default function Home() {
  return (
    <div className={styles.container}>
      <Head>
        <title>Whoops</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <AppBar />
      <main className={styles.main}>
        <h1 className={styles.title}>
          Whoops
        </h1>
      </main>

      <footer className={styles.footer}>

      </footer>
    </div>
  )
}
