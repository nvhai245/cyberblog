import '../styles/globals.css'
import '../styles/signin.scss'
import '../styles/NotFound.css'
import { ApolloProvider } from '@apollo/client'
import { useApollo } from '../libs/apolloClient'

function MyApp({ Component, pageProps }) {
  const apolloClient = useApollo(pageProps.initialApolloState)
  return (
    <ApolloProvider client={apolloClient}>
      <Component {...pageProps} />
    </ApolloProvider>
  )
}

export default MyApp
