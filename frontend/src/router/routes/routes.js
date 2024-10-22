
const routes =[
    {
        path: '/',
        name: 'home',
        component: () => import('../../pages/OtsHome.vue'),
    },
    {
        path: '/operation',
        name: 'operation',
        component: () => import('../../pages/OtsHome.vue'),
    },
    {
        path: '/unit',
        name: 'unit',
        component: () => import('../../pages/OtsHome.vue'),
    },
    {
        path: '/recipes',
        name: 'recipes',
        component: () => import('../../pages/OtsHome.vue'),
    },
    {
        path: '/fhx',
        name: 'fhx',
        component: () => import('../../pages/OtsFhx.vue'),
    },
    {
        path: '/users',
        name: 'users',
        component: () => import('../../pages/OtsHome.vue'),
    },
    {
        path: '/plant',
        name: 'plant',
        component: () => import('../../pages/OtsPlants.vue'),
    },
    {
        path: '/login',
        name: 'login',
        component: () => import('../../pages/OtsHome.vue'),
    },
    {
        path: '/logout',
        name: 'logout',
        component: () => import('../../pages/OtsHome.vue'),
    },
    {
        path: '/register',
        name: 'register',
        component: () => import('../../pages/OtsHome.vue'),
    },
];
export { routes }