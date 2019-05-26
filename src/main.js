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
global.api = 'https://yapi.airstone.me/mock/11' + '/api/'
// global.api = 'http://localhost:8080/'
global.md5 = md5
global.QRCode = QRCode

console.log(md5 ('hlynbnb'))

axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'
axios.defaults.withCredentials=true

// axios.defaults.headers.put['Content-Type'] = 'application/x-www-form-urlencoded'
// axios.defaults.headers.delete['Content-Type'] = 'application/x-www-form-urlencoded'

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
