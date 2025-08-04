import { createApp } from 'vue'
import App from './App.vue'
//引入arco-design
import ArcoVue from '@arco-design/web-vue';
import ArcoVueIcon from '@arco-design/web-vue/es/icon';
import '@arco-design/web-vue/dist/arco.css';



const app = createApp(App);

app.use(ArcoVue, {});
app.use(ArcoVueIcon);


app.mount('#app');