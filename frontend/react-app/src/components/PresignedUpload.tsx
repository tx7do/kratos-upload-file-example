import * as React from "react";
import {PlusOutlined} from '@ant-design/icons';

import {Space, Upload} from "antd";

import {
    RcFile,
    UploadFile,
} from 'antd/lib/upload/interface'
import type {UploadRequestOption} from 'rc-upload/lib/interface';

import {postFile, putFile, retrievePostUrl, retrievePutUrl} from "../util/file";

export type PresignedUploadProps = {};

interface PresignedUploadState {
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

class PresignedUpload extends React.Component<PresignedUploadProps, PresignedUploadState> {
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

        retrievePostUrl(aFile, (aFile, postUrl, formData) => {
            postFile(postUrl, aFile, formData, false, onProgress, onSuccess, onError);
        });
    }

    async handleCustomRequestUploadPut(options: UploadRequestOption) {
        const {
            file,
            onSuccess,
            onError,
            onProgress
        } = options;

        const aFile = file as RcFile;

        retrievePutUrl(aFile, (aFile, putUrl) => {
            putFile(putUrl, aFile, {}, onProgress, onSuccess, onError);
        });
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

export default PresignedUpload;
