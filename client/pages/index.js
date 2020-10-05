import Head from 'next/head'
import styles from '../styles/Home.module.css'
import Link from 'next/link'
import AppBar from '../components/AppBar'
import { initializeApollo } from '../libs/apolloClient'
import PostPreview from '../components/PostPreview'
import { GET_CATEGORY_POSTS } from '../libs/gql/allCharacters'


function Home() {
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
          <PostPreview />
        </main>

        <footer className={styles.footer}>

        </footer>
      </div>
    )
  }

  export async function getServerSideProps({req}) {
    const apolloClient = initializeApollo(null, req)
    let res = await apolloClient.query({
      query: GET_CATEGORY_POSTS,
    })
    return {
      props: {
        initialApolloState: apolloClient.cache.extract(),
      },
    }
  }


  export default Home;