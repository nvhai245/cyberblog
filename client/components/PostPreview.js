import React, {useState} from 'react';
import parseAuthCookie from "../libs/parseAuthCookie";
import Router, {useRouter} from "next/router";
import {GET_POST_BY_ID} from "../libs/gql/getPostById";
import {initializeApollo} from "../libs/apolloClient";
import {EditorState, convertFromRaw} from "draft-js";
import AppBar from "../components/AppBar";
import {GET_USER_BY_ID} from "../libs/gql/getUserById";
import {format, formatDistance, formatRelative, subDays} from 'date-fns';
import {vi} from 'date-fns/locale';
import Link from 'next/link';
import dynamic from "next/dynamic";
import styles from "../styles/Home.module.css";
import {useMutation} from "@apollo/client";
import {ADD_NEW_POST} from "../libs/gql/addNewPost";
import {PUBLISH_POST} from "../libs/gql/posts";
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

function PostPreview({initialAuthState, queryData, queryError, authorInfo, authorInfoError}) {
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
        queryResult = queryData
    }
    const contentState = convertFromRaw(JSON.parse(queryResult.content));
    const editorState = EditorState.createWithContent(contentState);
    const [editState, setEditState] = useState(true)
    const handleEdit = () => {

    }
    const [publishPost, {loading: mutationLoading, error: mutationError}] = useMutation(PUBLISH_POST)
    const handlePublish = (e) => {
        publishPost({
            variables: {
                requesterId: initialAuthState ? initialAuthState.user.id : authorInfo.getUserById.user.id,
                postId: queryResult && queryResult.id
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
            <div className="PostViewContainer">
                <div className="PostView">
                    <div className="post_title">
                        {queryResult.title}
                    </div>
                    <div className="post_info">
                        <i>
                            Đăng bởi [<Link
                            href={`/user/${authorInfo && authorInfo.getUserById.user.username}/profile`}><a>{authorInfo && authorInfo.getUserById.user.username}</a></Link>]
                            vào {format(new Date(queryResult.createdAt * 1000), 'PPPPpp', {locale: vi})}
                        </i>
                    </div>
                    <Editor
                        editorState={editorState} readOnly={editState}
                        toolbarHidden={editState}
                        plugins={[imagePlugin, focusPlugin, resizeablePlugin, blockDndPlugin, alignmentPlugin]}
                    />
                </div>
            </div>
        </div>
    );
}

export default PostPreview;