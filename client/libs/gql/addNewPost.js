import gql from 'graphql-tag'

export const ADD_NEW_POST = gql`
    mutation AddNewPost($title: String!, $authorID: Int!, $content: String!) {
        addNewPost(newPost: {
            id: 0,
            title: $title,
            authorID: $authorID,
            parentID: 0,
            content: $content
        }) {
            message
            post {
                id
            }
        }
    }
`;

