import Head from 'next/head'
import styles from '../styles/Home.module.css'
import Link from 'next/link'
import AppBar from '../components/AppBar'
import {initializeApollo} from '../libs/apolloClient'
import PostPreview from '../components/PostPreview'
import {GET_FEED} from '../libs/gql/getFeed'
import parseAuthCookie from "../libs/parseAuthCookie";


function Home({initialApolloState, initialAuthState}) {
    return (
        <div className={styles.container}>
            <Head>
                <title>Whoops</title>
                <link rel="icon" href="/favicon.ico"/>
            </Head>
            <AppBar authState={initialAuthState} />
            <main className={styles.main}>
                <h1 className={styles.title}>
                    Whoops
                </h1>
                <PostPreview initialApolloState={initialApolloState} />
            </main>

            <footer className={styles.footer}>

            </footer>
        </div>
    )
}

export async function getServerSideProps({req}) {
    const apolloClient = initializeApollo(null, req)
    const authCookie = await parseAuthCookie(req)
    let res = await apolloClient.query({
        query: GET_FEED,
        variables: {offset: 0, limit: 10}
    })
    return {
        props: {
            initialApolloState: apolloClient.cache.extract(),
            initialAuthState: (authCookie && authCookie.authInfo) ? JSON.parse(authCookie.authInfo) : null,
        },
    }
}


export default Home;