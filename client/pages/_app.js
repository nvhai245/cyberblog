import '../styles/globals.css'
import '../styles/NotFound.css'
import { ApolloProvider } from '@apollo/client'
import { useApollo } from '../libs/apolloClient'
import Layout from "../components/Layout";

function MyApp({ Component, pageProps }) {
  const apolloClient = useApollo(pageProps.initialApolloState)
  return (
    <ApolloProvider client={apolloClient}>
        <Layout />
        <Component {...pageProps} />
    </ApolloProvider>
  )
}

export default MyApp
