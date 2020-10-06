import gql from 'graphql-tag';

export const GET_FEED = gql`
    query GetFeed($offset: Int!, $limit: Int!) {
        getFeed(offset: $offset, limit: $limit) {
            message
            posts {
                id
                title
                authorID
                parentID
                content
                published
                upVote
                createdAt
                updatedAt
                publishedAt
            }
        }
    }
`;