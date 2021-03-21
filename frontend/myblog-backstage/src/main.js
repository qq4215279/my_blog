import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './permission';
import store from './store'
import config from './config/website'

import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';

console.log(process.env.NODE_ENV)
if (process.env.NODE_ENV === 'development' && config.dev.mock) {
  require('./mock/index')
}

Vue.config.productionTip = false
Vue.use(ElementUI);
new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
