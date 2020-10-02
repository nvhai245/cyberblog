import React from 'react'
import { GET_CATEGORY_POSTS } from '../libs/gql/allCharacters'
import { gql, useQuery } from '@apollo/client'

export default function PostPreview() {
    const { loading, error, data, fetchMore, networkStatus } = useQuery(GET_CATEGORY_POSTS, {
        notifyOnNetworkStatusChange: true,
    })
    return (
        <div>
            {data.getCategoryPosts.message}
            {data.getCategoryPosts.posts.map(post => 
                <div>
                    <li>{post.title}</li>
                    <li>{post.content}</li>
                </div>
            )}
        </div>
    )
}
