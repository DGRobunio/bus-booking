// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import jquery from 'jquery'
import axios from 'axios'
import md5 from 'js-md5'
import QRCode from 'qrcodejs2'
import 'bootstrap'
import 'popper.js'
import 'bootstrap/dist/css/bootstrap.css'

global.jquery = jquery
global.$ = axios
global.address = 'https://yapi.airstone.me/mock/11'
global.api = global.address + '/api/'
global.md5 = md5
global.QRCode = QRCode

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')