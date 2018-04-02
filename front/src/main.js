// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'
import api from './api'
import plugins from './plugins'
import './directive'
import './filter'

Vue.config.productionTip = false

// 全局注册
Vue.use(api)
Vue.use(plugins)

/* eslint-disable no-new */
let $VM = new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
 window.VM = $VM
 window.Vue = Vue
// Object.assign(window,{VM: $VM})

