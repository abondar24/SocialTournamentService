import Vue from 'vue'
import App from './App.vue'
import VueResource from 'vue-resource'
import BootstrapVue from 'bootstrap-vue'

Vue.config.productionTip = false;
Vue.use(VueResource);
Vue.use(BootstrapVue);

Vue.prototype.$hostname =  "http://"+window.location.hostname+":"+window.location.port;

new Vue({
  render: h => h(App)
}).$mount('#app');
