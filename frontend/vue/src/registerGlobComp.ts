import type { App } from 'vue';

import {
  Button,
  Card,
  Col,
  DatePicker,
  Divider,
  Dropdown,
  Image,
  Input,
  Layout,
  Popconfirm,
  Row,
  Space,
  Switch,
  Tabs,
  Tag,
  TimePicker,
  Tree,
  Upload,
} from 'ant-design-vue';

import VueUeditorWrap from 'vue-ueditor-wrap';

/**
 * 注册全局组件
 * @param app
 */
export function registerGlobComp(app: App) {
  app
    .use(Input)
    .use(Button)
    .use(Layout)
    .use(Space)
    .use(Card)
    .use(Switch)
    .use(Popconfirm)
    .use(Dropdown)
    .use(TimePicker)
    .use(DatePicker)
    .use(Tag)
    .use(Tabs)
    .use(Image)
    .use(Divider)
    .use(Row)
    .use(VueUeditorWrap)
    .use(Col)
    .use(Upload)
    .use(Tree);
}
