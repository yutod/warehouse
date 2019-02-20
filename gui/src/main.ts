import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import VueAxios from 'vue-axios'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import '@fortawesome/fontawesome-free/css/all.css'

Vue.config.productionTip = false

Vue.use(VueAxios, axios)
Vue.use(Vuetify, {
  iconfont: 'fa',
})

new Vue({
  router,

  render: (h) => h(App),
}).$mount('#app')
