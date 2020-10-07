import gql from 'graphql-tag';

export const REGISTER = gql`
    mutation Register($email: String!, $password: String!, $username: String!) {
        register(email: $email, password: $password, username: $username) {
            message
            user {
                id
                username
                firstName
                lastName
                email
                avatar
                createdAt
                updatedAt
            }
        }
    }
`;