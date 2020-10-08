import React, {useState, useRef} from 'react';
import dynamic from 'next/dynamic';
import { useRouter } from 'next/router'
import 'react-draft-wysiwyg/dist/react-draft-wysiwyg.css';
import styles from "../../../styles/Home.module.css";
const Editor = dynamic(() =>
    import('react-draft-wysiwyg').then((mod) => mod.Editor), { ssr: false }
)
const EditorState = dynamic(() =>
    import('draft-js').then((mod) => mod.EditorState), { ssr: false }
)

function NewPost(props) {
    const router = useRouter()
    const [edtState, setEdtState] = useState(EditorState.createEmpty)
    const handleEditorStateChange = (eState) => {
        setEdtState(eState)
    }
    const uploadImage = () => {

    }
    const handleTitleChange = () => {

    }
    const handleSubmit = () => {

    }
    return (
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
    );
}

export default NewPost;