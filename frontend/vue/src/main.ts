import { createApp } from 'vue';
import App from './App.vue';
import { registerGlobComp } from '@/registerGlobComp';

async function bootstrap() {
  const app = createApp(App);

  // 注册全局组件
  registerGlobComp(app);

  app.mount('#app');
}

await bootstrap();
