import React from 'react';
import styles from '../styles/Loading.module.css'

function Loading(props) {
    return (
        <div>
            <span className={styles.spinner + " " + styles.three} />
        </div>
    );
}

export default Loading;