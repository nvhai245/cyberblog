import { useMemo } from 'react'
import { ApolloClient, createHttpLink, InMemoryCache } from '@apollo/client'
import { concatPagination } from '@apollo/client/utilities'

let apolloClient

function createApolloClient(req) {
    const link = createHttpLink({
        uri: 'http://localhost:8080/query', // Server URL (must be absolute)
        credentials: 'include', // Additional fetch() options like `credentials` or `headers`
        headers: req ? {
            cookie: req.headers.cookie,
        } : undefined,
    })
    return new ApolloClient({
        ssrMode: typeof window === 'undefined',
        link: link,
        cache: new InMemoryCache({
            typePolicies: {
                Query: {
                    fields: {
                        allPosts: concatPagination(),
                    },
                },
            },
        }),
    })
}

export function initializeApollo(initialState, req) {
    const _apolloClient = createApolloClient(req)

    // If your page has Next.js data fetching methods that use Apollo Client, the initial state
    // gets hydrated here
    if (initialState) {
        // Get existing cache, loaded during client side data fetching
        const existingCache = _apolloClient.extract()
        // Restore the cache using the data passed from getStaticProps/getServerSideProps
        // combined with the existing cached data
        _apolloClient.cache.restore({ ...existingCache, ...initialState })
    }
    // For SSG and SSR always create a new Apollo Client
    if (typeof window === 'undefined') return _apolloClient
    // Create the Apollo Client once in the client
    if (!apolloClient) apolloClient = _apolloClient

    return _apolloClient
}

export function useApollo(initialState) {
    const store = useMemo(() => initializeApollo(initialState), [initialState])
    return store
}
