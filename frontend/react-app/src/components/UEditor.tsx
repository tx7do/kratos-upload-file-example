import * as React from "react";
import {Space} from "antd";

// @ts-ignore
import RcUEditor from 'react-ueditor-wrap';

import {API_URL, OSS_URL} from "../util/const";

export type UEditorProps = {};

const ueditorConfigUrl = `${API_URL}/ueditor`;
const getImageDownloadUrl = OSS_URL;

const editorConfig = {
    // UEditorPlus 的根目录，绝对路径。
    UEDITOR_HOME_URL: '/static/UEditorPlus/',
    // UEditorPlus 的跨域根目录，绝对路径。
    UEDITOR_CORS_URL: '/static/UEditorPlus/',

    // 后端服务地址，后端处理参考
    // https://open-doc.modstart.com/ueditor-plus/backend.html
    serverUrl: ueditorConfigUrl,

    catchRemoteImageEnable: false /* 抓取远程图片是否开启,默认true */,

    autoHeightEnabled: false /* 固定高度且带滚动条 */,
    autoSyncData: false,
    autoFloatEnabled: true,

    theme: 'default',
    themePath: '/static/UEditorPlus/themes/',

    /* 上传图片配置项 */
    imageActionName: 'uploadImage' /* 执行上传图片的action名称 */,
    imageFieldName: 'file' /* 提交的图片表单名称 */,
    imageMaxSize: 2_048_000 /* 上传大小限制，单位B */,
    imageAllowFiles: [
        '.png',
        '.jpg',
        '.jpeg',
        '.gif',
        '.bmp',
    ] /* 上传图片格式显示 */,
    imageCompressEnable: true /* 是否压缩图片,默认是true */,
    imageCompressBorder: 1600 /* 图片压缩最长边限制 */,
    imageInsertAlign: 'none' /* 插入的图片浮动方式 */,
    imageUrlPrefix: getImageDownloadUrl /* 图片访问路径前缀 */,

    /* 涂鸦图片上传配置项 */
    scrawlActionName: 'uploadScrawl' /* 执行上传涂鸦的action名称 */,
    scrawlFieldName: 'file' /* 提交的图片表单名称 */,
    scrawlMaxSize: 2_048_000 /* 上传大小限制，单位B */,
    scrawlUrlPrefix: getImageDownloadUrl /* 图片访问路径前缀 */,
    scrawlInsertAlign: 'none',

    /* 截图工具上传 */
    snapscreenActionName: 'uploadImage' /* 执行上传截图的action名称 */,
    snapscreenUrlPrefix: getImageDownloadUrl /* 图片访问路径前缀 */,
    snapscreenInsertAlign: 'none' /* 插入的图片浮动方式 */,

    /* 抓取远程图片配置 */
    catcherLocalDomain: ['127.0.0.1', 'localhost', 'img.baidu.com'],
    catcherActionName: 'catchImage' /* 执行抓取远程图片的action名称 */,
    catcherFieldName: 'file' /* 提交的图片列表表单名称 */,
    catcherUrlPrefix: getImageDownloadUrl /* 图片访问路径前缀 */,
    catcherMaxSize: 2_048_000 /* 上传大小限制，单位B */,
    catcherAllowFiles: [
        '.png',
        '.jpg',
        '.jpeg',
        '.gif',
        '.bmp',
    ] /* 抓取图片格式显示 */,

    /* 上传视频配置 */
    videoActionName: 'uploadVideo' /* 执行上传视频的action名称 */,
    videoFieldName: 'file' /* 提交的视频表单名称 */,
    videoUrlPrefix: getImageDownloadUrl /* 视频访问路径前缀 */,
    videoMaxSize: 102_400_000 /* 上传大小限制，单位B，默认100MB */,
    videoAllowFiles: [
        '.flv',
        '.swf',
        '.mkv',
        '.avi',
        '.rm',
        '.rmvb',
        '.mpeg',
        '.mpg',
        '.ogg',
        '.ogv',
        '.mov',
        '.wmv',
        '.mp4',
        '.webm',
        '.mp3',
        '.wav',
        '.mid',
    ] /* 上传视频格式显示 */,

    /* 上传文件配置 */
    fileActionName: 'uploadFile' /* controller里,执行上传视频的action名称 */,
    fileFieldName: 'file' /* 提交的文件表单名称 */,
    fileUrlPrefix: getImageDownloadUrl /* 文件访问路径前缀 */,
    fileMaxSize: 51_200_000 /* 上传大小限制，单位B，默认50MB */,
    fileAllowFiles: [
        '.png',
        '.jpg',
        '.jpeg',
        '.gif',
        '.bmp',
        '.flv',
        '.swf',
        '.mkv',
        '.avi',
        '.rm',
        '.rmvb',
        '.mpeg',
        '.mpg',
        '.ogg',
        '.ogv',
        '.mov',
        '.wmv',
        '.mp4',
        '.webm',
        '.mp3',
        '.wav',
        '.mid',
        '.rar',
        '.zip',
        '.tar',
        '.gz',
        '.7z',
        '.bz2',
        '.cab',
        '.iso',
        '.doc',
        '.docx',
        '.xls',
        '.xlsx',
        '.ppt',
        '.pptx',
        '.pdf',
        '.txt',
        '.md',
        '.xml',
    ] /* 上传文件格式显示 */,

    /* 列出指定目录下的图片 */
    imageManagerActionName: 'listImage' /* 执行图片管理的action名称 */,
    imageManagerListSize: 20 /* 每次列出文件数量 */,
    imageManagerUrlPrefix: getImageDownloadUrl /* 图片访问路径前缀 */,
    imageManagerInsertAlign: 'none' /* 插入的图片浮动方式 */,
    imageManagerAllowFiles: [
        '.png',
        '.jpg',
        '.jpeg',
        '.gif',
        '.bmp',
    ] /* 列出的文件类型 */,

    /* 列出指定目录下的文件 */
    fileManagerActionName: 'listFile' /* 执行文件管理的action名称 */,
    fileManagerUrlPrefix: getImageDownloadUrl /* 文件访问路径前缀 */,
    fileManagerListSize: 20 /* 每次列出文件数量 */,
    fileManagerAllowFiles: [
        '.png',
        '.jpg',
        '.jpeg',
        '.gif',
        '.bmp',
        '.flv',
        '.swf',
        '.mkv',
        '.avi',
        '.rm',
        '.rmvb',
        '.mpeg',
        '.mpg',
        '.ogg',
        '.ogv',
        '.mov',
        '.wmv',
        '.mp4',
        '.webm',
        '.mp3',
        '.wav',
        '.mid',
        '.rar',
        '.zip',
        '.tar',
        '.gz',
        '.7z',
        '.bz2',
        '.cab',
        '.iso',
        '.doc',
        '.docx',
        '.xls',
        '.xlsx',
        '.ppt',
        '.pptx',
        '.pdf',
        '.txt',
        '.md',
        '.xml',
    ],
};

class UEditor extends React.Component<UEditorProps> {
    constructor(props: {}) {
        super(props);
    }

    render() {
        return (
            <div className="center-content">
                <Space>
                    <RcUEditor
                        ueditorUrl={'/static/UEditorPlus/ueditor.all.js'}
                        ueditorConfigUrl={'/static/UEditorPlus/ueditor.config.js'}
                        editorConfig={editorConfig}
                    />
                </Space>
            </div>
        );
    }
}

export default UEditor;
