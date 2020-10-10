import gql from 'graphql-tag';

export const PUBLISH_POST = gql`
    mutation PublishPost($requesterId: Int!, $postId: Int!, ) {
        publishPost(requesterId: $requesterId, postId: $postId) {
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