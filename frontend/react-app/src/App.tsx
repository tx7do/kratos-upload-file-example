import React from 'react';
import './App.css';

import {Tabs} from 'antd';
import type {TabsProps} from 'antd';

import DirectUpload from './components/DirectUpload';
import PresignedUpload from './components/PresignedUpload';
import UEditor from './components/UEditor';

const items: TabsProps['items'] = [
    {
        key: '1',
        label: '服务直接上传',
        children: <DirectUpload/>,
    },
    {
        key: '2',
        label: '预签名上传',
        children: <PresignedUpload/>,
    },
    {
        key: '3',
        label: 'UEditor',
        children: <UEditor/>,
    },
];

const App: React.FC = () => <Tabs defaultActiveKey="1" items={items}/>;

export default App;
