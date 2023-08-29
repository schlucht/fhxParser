import './assets/css/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
// import { usePlantStore } from './stores/plant_store.js';

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)
// app.use(usePlantStore)
app.mount('#app')
