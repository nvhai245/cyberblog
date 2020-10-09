import gql from 'graphql-tag';

export const GET_POST_BY_ID = gql`
    query GetPostById($requesterId: Int!, $postId: Int!, ) {
        getPostById(requesterId: $requesterId, postId: $postId) {
            message
            post {
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