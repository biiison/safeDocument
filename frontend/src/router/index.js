// Instalar as dependências necessárias antes de rodar o projeto:
// npm install vue-router@4 vuetify@3 axios

import { createApp } from 'vue';
import { createRouter, createWebHistory } from 'vue-router';
import App from '../App.vue';
import Vuetify from 'vuetify';
import 'vuetify/styles';
import Home from '../views/Home.vue';
import AddCPF from '../views/AddCPF.vue';

const routes = [
    { path: '/', component: Home },
    { path: '/add', component: AddCPF },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

const app = createApp(App);
app.use(router);
app.use(Vuetify);
app.mount('#app');
