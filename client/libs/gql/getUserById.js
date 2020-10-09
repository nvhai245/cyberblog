import gql from 'graphql-tag';

export const GET_USER_BY_ID = gql`
    query GetUserById($requesterId: Int!, $userId: Int!, ) {
        getUserById(requestorId: $requesterId, userId: $userId) {
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