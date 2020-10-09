import React from 'react';
import parseAuthCookie from "../../../../libs/parseAuthCookie";
import Router, { useRouter } from "next/router";
import {GET_POST_BY_ID} from "../../../../libs/gql/getPostById";
import {initializeApollo} from "../../../../libs/apolloClient";
import { EditorState, convertFromRaw } from "draft-js";
import AppBar from "../../../../components/AppBar";
import {GET_USER_BY_ID} from "../../../../libs/gql/getUserById";
import { format, formatDistance, formatRelative, subDays } from 'date-fns';
import { vi } from 'date-fns/locale';
import Link from 'next/link';
import dynamic from "next/dynamic";
import 'react-draft-wysiwyg/dist/react-draft-wysiwyg.css';

const Editor = dynamic(() =>
    import('react-draft-wysiwyg').then((mod) => mod.Editor), { ssr: false }
)
function PostViewById({initialAuthState, queryData, queryError, authorInfo, authorInfoError}) {
    let queryResult
    const router = useRouter()
    if (queryError) {
        return (
            <div style={{color: "white", marginTop: "1rem", marginLeft: "1rem"}}>
                {queryError[0].message.replace('rpc error: code = Unknown desc = ', '')}
            </div>
        )
    } else {
        queryResult = queryData.data.getPostById
    }
    const contentState = convertFromRaw(JSON.parse(queryResult.post.content));
    const editorState = EditorState.createWithContent(contentState);
    return (
        <div>
            <AppBar authState={initialAuthState} />
            <div className="PostView">
                <div className="post_title">
                    {queryResult.post.title}
                </div>
                <div className="post_info">
                    <i>
                        Đăng bởi [<Link href={`/user/${authorInfo && authorInfo.getUserById.user.username}/profile`}><a>{authorInfo && authorInfo.getUserById.user.username}</a></Link>] vào {format(new Date(queryData.data.getPostById.post.createdAt * 1000), 'PPPPpp', {locale: vi})}
                    </i>
                </div>
                <Editor
                    editorState={editorState} readOnly={true}
                    toolbarHidden
                />
            </div>
        </div>
    );
}

export default PostViewById;

export async function getServerSideProps({req, query}) {
    const authCookie = await parseAuthCookie(req)
    let authInfo = (authCookie && authCookie.authInfo) ? JSON.parse(authCookie.authInfo) : null
    const apolloClient = initializeApollo(null, req)
    let queryData, queryError, authorInfo, authorInfoError
    if (authInfo) {
        await apolloClient.query({
            query: GET_POST_BY_ID,
            variables: {requesterId: authInfo.user.id, postId: query.postId}
        }).then(async (data) => {
            queryData = data
            await apolloClient.query({
                query: GET_USER_BY_ID,
                variables: {requesterId: authInfo.user.id, userId: data.data.getPostById.post.authorID}
            }).then(res => {
                authorInfo = res.data
            }).catch(err => {
                authorInfoError = err
            })
        }).catch(err => {
            queryError = err
        })
    }
    return {
        props: {
            initialAuthState: authInfo,
            queryData: queryData ? queryData : null,
            queryError: queryError ? queryError.graphQLErrors : null,
            authorInfo: authorInfo ? authorInfo : null,
            authorInfoError: authorInfoError ? authorInfoError.graphQLErrors : null
        },
    }
}