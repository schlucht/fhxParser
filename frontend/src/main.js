import { createApp } from 'vue';
import './assets/css/main.css';
import App from './App.vue';
import { router } from './router/routes';

const app = createApp(App);
app.use(router);
app.mount('#app');

