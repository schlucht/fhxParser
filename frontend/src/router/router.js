import {createRouter, createWebHistory} from 'vue-router';

import { routes } from './routes/routes.js';

  const router = createRouter({
    history: createWebHistory(),
    routes, 
  });

  export { router }

