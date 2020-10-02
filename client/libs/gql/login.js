import gql from 'graphql-tag';

export const LOGIN = gql`
    mutation Login($email: String!, $password: String!) {
    login(email: $email, password: $password) {
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