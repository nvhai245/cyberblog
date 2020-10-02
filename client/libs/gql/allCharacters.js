import gql from 'graphql-tag';

export const GET_CATEGORY_POSTS = gql`
    query {
    getCategoryPosts(categoryId: 3, offset: 0, limit: 10) {
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