import React from 'react'
import { GET_FEED } from '../libs/gql/getFeed'
import { gql, useQuery } from '@apollo/client'

export default function PostPreview(initialApolloState) {
    const { loading, error, data, fetchMore, networkStatus } = useQuery(GET_FEED, {
        notifyOnNetworkStatusChange: true,
        variables: {offset: 0, limit: 10}
    })
    return (
        <div>
            {data.getFeed.message}
            {data.getFeed.posts.map(post =>
                <div>
                    <li>{post.title}</li>
                    <li>{post.content}</li>
                </div>
            )}
        </div>
    )
}
