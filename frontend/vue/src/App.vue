<template>
  <div class="app">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="1" tab="服务直接上传">
        <DirectUpload/>
      </a-tab-pane>
      <a-tab-pane key="2" tab="预签名上传">
        <PresignedUpload/>
      </a-tab-pane>
      <a-tab-pane key="3" tab="UEditor">
        <UEditor/>
      </a-tab-pane>
    </a-tabs>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';

import DirectUpload from '@/views/DirectUpload.vue';
import PresignedUpload from '@/views/PresignedUpload.vue';
import UEditor from '@/views/UEditor.vue';

const activeKey = ref('1');
</script>

<style lang="less">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
