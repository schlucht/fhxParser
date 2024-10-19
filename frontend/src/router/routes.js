import {createRouter, createMemoryHistory} from 'vue-router';
import App from '../App.vue';

const routes = [
    { path: '/', component: App},
    
  ]
  
  const router = createRouter({
    history: createMemoryHistory(),
    routes,
  })

  export { router }

