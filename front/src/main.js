// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'

import router from './router'
import store from './store'
import api from './api'
import './plugins'
import './directive'
import './filter'
import './venders'
import './assets/styles'
import App from './App'


Vue.config.productionTip = false
window.Vue = Vue

// 全局注册
Vue.use(api)

/* eslint-disable no-new */
let $VM = new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})

// vender window
Object.assign(window, {VM: $VM})



