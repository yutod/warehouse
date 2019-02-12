import Vue from 'vue';
import App from './App.vue';
import router from './router';
import axios from 'axios'
import VueAxios from 'vue-axios'
// import ApolloClient from 'apollo-boost';
// import VueApollo from 'vue-apollo';

Vue.config.productionTip = false;

// Vue.use(VueApollo);
Vue.use(VueAxios, axios);

// const apolloClient = new ApolloClient({
  // uri: 'http://localhost:3001/graphql',
  // fetchOptions: { useGETForQueries: true },
// });

// const apolloProvider = new VueApollo({
  // defaultClient: apolloClient,
// });

new Vue({
  router,
  // apolloProvider,
  render: (h) => h(App),
}).$mount('#app');
