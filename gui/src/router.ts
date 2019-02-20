import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Configuration from './views/Configuration.vue'
import Dependency from './views/Dependency.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
    },
    {
      path: '/dependency',
      name: 'dependency',
      // route level code-splitting
      // this generates a separate chunk (about.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      // component: () => import(/* webpackChunkName: "about" */ './views/Dependency.vue'),
      component: Dependency,
    },
    {
      path: '/configuration',
      name: 'configuration',
      component: Configuration,
    },
  ],
})
