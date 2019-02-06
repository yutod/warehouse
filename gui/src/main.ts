import Vue from 'vue';
import App from './App.vue';
import router from './router';
// import execa from 'execa'

const execa = require('execa');

(async () => {
  const { stdout } = await execa('echo', ['Hello World']);
  console.log(stdout);
})();

// (async () => {
// 	const {stdout} = await execa('echo', ['unicorns']);
// 	console.log(stdout);
// 	//=> 'unicorns'
// })();

Vue.config.productionTip = false;

new Vue({
  router,
  render: (h) => h(App),
}).$mount('#app');
