import * as React from "react";
import {PlusOutlined} from '@ant-design/icons';

import {Space, Upload} from "antd";

import {
    RcFile,
    UploadFile,
} from 'antd/lib/upload/interface'
import type {UploadRequestOption} from 'rc-upload/lib/interface';


import {API_URL} from "../util/const";
import {postFile, putFile} from "../util/file";

export type DirectUploadProps = {};

interface DirectUploadState {
    fileListPost: UploadFile[];
    fileListPut: UploadFile[];
}

const uploadButtonPost = (
    <button style={{border: 0, background: 'none'}} type="button">
        <PlusOutlined/>
        <div style={{marginTop: 8}}>POST Upload</div>
    </button>
);

const uploadButtonPut = (
    <button style={{border: 0, background: 'none'}} type="button">
        <PlusOutlined/>
        <div style={{marginTop: 8}}>PUT Upload</div>
    </button>
);

class DirectUpload extends React.Component<DirectUploadProps, DirectUploadState> {
    constructor(props: {}) {
        super(props);
        this.state = {
            fileListPost: [],
            fileListPut: [],
        };
        this.handleCustomRequestUploadPost = this.handleCustomRequestUploadPost.bind(this);
        this.handleCustomRequestUploadPut = this.handleCustomRequestUploadPut.bind(this);
    }

    async handleCustomRequestUploadPost(options: UploadRequestOption) {
        const {
            file,
            onSuccess,
            onError,
            onProgress
        } = options;

        const aFile = file as RcFile;

        const url = `${API_URL}/file:upload`;

        postFile(url, aFile, {}, true, onProgress, onSuccess, onError);
    }

    async handleCustomRequestUploadPut(options: UploadRequestOption) {
        const {
            file,
            onSuccess,
            onError,
            onProgress
        } = options;

        const aFile = file as RcFile;

        const url = `${API_URL}/file:upload`;

        putFile(url, aFile, {}, onProgress, onSuccess, onError);
    }

    render() {
        return (
            <Space direction={"vertical"}>
                <Upload
                    listType="picture-card"
                    customRequest={this.handleCustomRequestUploadPost}
                >
                    {this.state.fileListPost.length >= 8 ? null : uploadButtonPost}
                </Upload>
                <Upload
                    listType="picture-card"
                    customRequest={this.handleCustomRequestUploadPut}
                >
                    {this.state.fileListPut.length >= 8 ? null : uploadButtonPut}
                </Upload>
            </Space>
        );
    }
}

export default DirectUpload;
