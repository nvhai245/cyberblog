import React, {useState} from 'react';
import parseAuthCookie from "../../../../libs/parseAuthCookie";
import Router, {useRouter} from "next/router";
import {GET_POST_BY_ID} from "../../../../libs/gql/getPostById";
import {initializeApollo} from "../../../../libs/apolloClient";
import {EditorState, convertFromRaw} from "draft-js";
import AppBar from "../../../../components/AppBar";
import {GET_USER_BY_ID} from "../../../../libs/gql/getUserById";
import {format, formatDistance, formatRelative, subDays} from 'date-fns';
import {vi} from 'date-fns/locale';
import Link from 'next/link';
import dynamic from "next/dynamic";
import styles from "../../../../styles/Home.module.css";
import {useMutation} from "@apollo/client";
import {ADD_NEW_POST} from "../../../../libs/gql/addNewPost";
import {PUBLISH_POST} from "../../../../libs/gql/posts";
import Editor from 'draft-js-plugins-editor';
import createImagePlugin from 'draft-js-image-plugin';
const imagePlugin = createImagePlugin();
import 'react-draft-wysiwyg/dist/react-draft-wysiwyg.css';

import createAlignmentPlugin from 'draft-js-alignment-plugin';

import createFocusPlugin from 'draft-js-focus-plugin';

import createResizeablePlugin from 'draft-js-resizeable-plugin';

import createBlockDndPlugin from 'draft-js-drag-n-drop-plugin';

const focusPlugin = createFocusPlugin();
const resizeablePlugin = createResizeablePlugin();
const blockDndPlugin = createBlockDndPlugin();
const alignmentPlugin = createAlignmentPlugin();

// const Editor = dynamic(() =>
//     import('react-draft-wysiwyg').then((mod) => mod.Editor), {ssr: false}
// )

function PostViewById({initialAuthState, queryData, queryError, authorInfo, authorInfoError}) {
    let queryResult
    const router = useRouter()
    if (queryError || !queryData) {
        console.log(queryError)
        return (
            <div>
                <AppBar authState={initialAuthState}/>
                <div style={{color: "white", marginTop: "1rem", marginLeft: "1rem"}}>
                    {queryError && queryError[0].message.replace('rpc error: code = Unknown desc = ', '')}
                </div>
            </div>
        )
    } else {
        queryResult = queryData.data.getPostById
    }
    const contentState = convertFromRaw(JSON.parse(queryResult.post.content));
    const editorState = EditorState.createWithContent(contentState);
    const [editState, setEditState] = useState(true)
    const handleEdit = () => {

    }
    const [publishPost, {loading: mutationLoading, error: mutationError}] = useMutation(PUBLISH_POST)
    const handlePublish = (e) => {
        publishPost({
            variables: {
                requesterId: initialAuthState ? initialAuthState.user.id : authorInfo.getUserById.user.id,
                postId: queryResult && queryResult.post.id
            }
        }).then(res => {
            if (res.data.publishPost && res.data.publishPost.message == "publish post successfully!") {
                console.log(res.data.publishPost)
                Router.push("/")
            }
        }).catch(error => {
            console.log(error)
        })
    }
    return (
        <div>
            <AppBar authState={initialAuthState}/>
            <div className="PostViewContainer">
                <div className="PostView">
                    <div className="post_title">
                        {queryResult.post.title}
                    </div>
                    <div className="post_info">
                        <i>
                            Đăng bởi [<Link
                            href={`/user/${authorInfo && authorInfo.getUserById.user.username}/profile`}><a>{authorInfo && authorInfo.getUserById.user.username}</a></Link>]
                            vào {format(new Date(queryData.data.getPostById.post.createdAt * 1000), 'PPPPpp', {locale: vi})}
                        </i>
                    </div>
                    <Editor
                        editorState={editorState} readOnly={editState}
                        toolbarHidden={editState}
                        plugins={[imagePlugin, focusPlugin, resizeablePlugin, blockDndPlugin, alignmentPlugin]}
                    />
                </div>
                {initialAuthState && authorInfo && initialAuthState.loggedIn && initialAuthState.user.id === authorInfo.getUserById.user.id &&
                <div className="post_options">
                    <div>
                        <div style={{marginTop: "0.5rem"}} className="btn btn--primary">
                            <button onClick={handleEdit} type="submit" className="text-button">
                                <a href="#">
                                    [Chỉnh sửa]
                                </a>
                            </button>
                        </div>
                    </div>
                    <div>
                        <div style={{marginTop: "0.5rem"}} className="btn btn--primary">
                            <button onClick={handlePublish} type="submit" className="text-button">
                                <a href="#">
                                    [Xuất bản]
                                </a>
                            </button>
                        </div>
                    </div>
                </div>
                }
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
        await apolloClient.query({
            query: GET_POST_BY_ID,
            variables: {requesterId: authInfo ? authInfo.user.id : 0, postId: query.postId}
        }).then(async (data) => {
            queryData = data
            await apolloClient.query({
                query: GET_USER_BY_ID,
                variables: {requesterId: authInfo ? authInfo.user.id : 0, userId: data.data.getPostById.post.authorID}
            }).then(res => {
                authorInfo = res.data
            }).catch(err => {
                authorInfoError = err
            })
        }).catch(err => {
            queryError = err
        })
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