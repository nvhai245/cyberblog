import React, {useState, useEffect} from 'react';
import dynamic from 'next/dynamic';
import Router, { useRouter } from 'next/router'
import 'react-draft-wysiwyg/dist/react-draft-wysiwyg.css';
import styles from "../../../styles/Home.module.css";
import {convertToRaw} from 'draft-js';
import AppBar from "../../../components/AppBar";
import { ADD_NEW_POST } from '../../../libs/gql/addNewPost'
import parseAuthCookie from "../../../libs/parseAuthCookie";
import {useMutation} from "@apollo/client";
import {REGISTER} from "../../../libs/gql/register";
const Editor = dynamic(() =>
    import('react-draft-wysiwyg').then((mod) => mod.Editor), { ssr: false }
)
const EditorState = dynamic(() =>
    import('draft-js').then((mod) => mod.EditorState), { ssr: false }
)

function NewPost({initialAuthState}) {
    const router = useRouter()
    const [addNewPost, {loading: mutationLoading, error: mutationError}] = useMutation(ADD_NEW_POST)
    const [edtState, setEdtState] = useState(EditorState.createEmpty)
    const [rawState, setRawState] = useState("")
    const [title, setTitle] = useState("")
    const handleEditorStateChange = (eState) => {
        setEdtState(eState)
    }
    const uploadImage = () => {
        return new Promise((resolve) => {
            resolve({ data: { link: "https://i.redd.it/fouvjs5eyuwz.jpg"}});
        });
    }
    const handleTitleChange = (e) => {
        setTitle(e.target.value)
    }
    const handleSubmit = () => {
        let rawState = JSON.stringify(convertToRaw(edtState.getCurrentContent()))
        if (initialAuthState && initialAuthState.loggedIn) {
            addNewPost({
                variables: {
                    title: title,
                    authorID: initialAuthState.user.id,
                    content: rawState
                }
            }).then(({data}) => {
                console.log(data.addNewPost)
                if (data.addNewPost && data.addNewPost.message == "add new post successfully!") {
                    Router.push(`/user/${initialAuthState.user.username}/post/${data.addNewPost.post.id}`)
                }
            }).catch(error => {
                console.log(error.toString())
            })
        } else {
            Router.push("/signin")
        }
    }
    return (
        <div>
            <AppBar authState={initialAuthState} />
            <div className="container">
                <div className="newPostBox">
                    <div className="titleBox">
                        <div className="editor-field__label-container">
                            <p style={{fontSize: "1.4em"}} className="editor-field__label">[Tiêu đề]:</p>
                        </div>
                        <div className="editor-field__container" style={{width: "50%", marginLeft: "1rem"}}>
                            <input onChange={handleTitleChange} type="text" className="editor-field__input" style={{color: "white"}} />
                        </div>
                    </div>
                    <Editor
                        editorState={edtState}
                        toolbarClassName="editorToolBar"
                        wrapperClassName="editorWrapper"
                        editorClassName="editorTextArea"
                        onEditorStateChange={handleEditorStateChange}
                        toolbar={{
                            image: {
                                uploadCallback: uploadImage,
                                previewImage: true,
                            }
                        }}
                    />
                    <div onClick={handleSubmit} className={styles.authButtons}>
                        <div style={{ marginTop: "0.5rem" }} className="btn btn--primary">
                            <button type="submit" className="text-button">
                                <a href="#">
                                    [Lưu bài viết]
                                </a>
                            </button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default NewPost;

export async function getServerSideProps({req}) {
    const authCookie = await parseAuthCookie(req)
    return {
        props: {
            initialAuthState: (authCookie && authCookie.authInfo) ? JSON.parse(authCookie.authInfo) : null,
        },
    }
}