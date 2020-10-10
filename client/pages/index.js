import Head from 'next/head'
import styles from '../styles/Home.module.css'
import Link from 'next/link'
import AppBar from '../components/AppBar'
import {initializeApollo} from '../libs/apolloClient'
import PostPreview from '../components/PostPreview'
import {GET_FEED} from '../libs/gql/getFeed'
import {GET_USER_BY_ID} from '../libs/gql/getUserById'
import parseAuthCookie from "../libs/parseAuthCookie";
import Layout from "../components/Layout";


function Home({initialApolloState, initialAuthState, initialFeedState}) {
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
                {initialFeedState.map(state =>
                    <PostPreview initialAuthState={initialAuthState} queryData={state.post} authorInfo={state.authorInfo} authorInfoError={state.authorInfoError} />
                )}
            </main>

            <footer className={styles.footer}>

            </footer>
        </div>
    )
}

export async function getServerSideProps({req}) {
    const apolloClient = initializeApollo(null, req)
    const authCookie = await parseAuthCookie(req)
    let authInfo = (authCookie && authCookie.authInfo) ? JSON.parse(authCookie.authInfo) : null
    let postInfos = []
    let postRes
    await apolloClient.query({
        query: GET_FEED,
        variables: {offset: 0, limit: 10}
    }).then(data => {
        postRes = data
        return postRes
    }).catch(err => {
        console.log(err.graphQLErrors)
    })
    console.log("postres", postRes)
    let infos = await postRes.data.getFeed.posts.map(async post => {
        let postInfo
        await apolloClient.query({
            query: GET_USER_BY_ID,
            variables: {requesterId: authInfo ? authInfo.user.id : 0, userId: post.authorID}
        }).then(res => {
            postInfo = {
                post: post,
                authorInfo: res.data,
                authorInfoError: null
            }
        }).catch(err => {
            postInfo = {
                post: post,
                authorInfo: null,
                authorInfoError: err.graphQLErrors
            }
        })
        return postInfo
    })
    await Promise.allSettled(infos).then(promises => {
        infos = promises
    })
    await infos.map(info => {
        postInfos.push(info.value)
    })
    console.log("*************************", postInfos)
    return {
        props: {
            initialApolloState: apolloClient.cache.extract(),
            initialAuthState: authInfo,
            initialFeedState: postInfos
        },
    }
}


export default Home;