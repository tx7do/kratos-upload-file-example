import * as React from "react";
import {Space} from "antd";

export type DirectUploadProps = {};

class DirectUpload extends React.Component<DirectUploadProps> {
    constructor(props: DirectUploadProps) {
        super(props);
        this.handleUploadFile = this.handleUploadFile.bind(this);
    }

    async handleUploadFile(file?: File) {
    }

    render() {
        return (
            <Space>
            </Space>
        );
    }
}

export default DirectUpload;
