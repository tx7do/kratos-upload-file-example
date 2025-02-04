<template>
  <a-space direction="vertical">
    <a-upload
      v-model:file-list="fileListPost"
      :customRequest="handleCustomRequestUploadPost"
      list-type="picture-card"
    >
      <div v-if="fileListPost.length < 8">
        <plus-outlined/>
        <div style="margin-top: 8px">POST Upload</div>
      </div>
    </a-upload>
    <a-upload
      v-model:file-list="fileListPut"
      :customRequest="handleCustomRequestUploadPut"
      list-type="picture-card"
    >
      <div v-if="fileListPut.length < 8">
        <plus-outlined/>
        <div style="margin-top: 8px">PUT Upload</div>
      </div>
    </a-upload>
  </a-space>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { UploadProps } from 'ant-design-vue';
import { PlusOutlined } from '@ant-design/icons-vue';
import { API_URL } from '@/util/const';
import { RcFile, UploadRequestOption } from 'ant-design-vue/es/vc-upload/interface';
import { postFile, putFile } from '@/util/file';

const fileListPost = ref<UploadProps['fileList']>([]);
const fileListPut = ref<UploadProps['fileList']>([]);

function handleCustomRequestUploadPost(options: UploadRequestOption) {
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

function handleCustomRequestUploadPut(options: UploadRequestOption) {
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
</script>

<style scoped lang="less">

</style>
