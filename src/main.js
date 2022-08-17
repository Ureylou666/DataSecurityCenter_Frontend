import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import './plugin/element-ui'
import './assets/css/global.css'

router.beforeEach((to, from, next) => {
  /* 路由发生变化时 修改页面title */
  if (to.meta.title) {
    document.title = to.meta.title
  }
  next()
})

Vue.prototype.$http = axios
axios.defaults.baseURL = 'http://127.0.0.1:8088/api/v1/'
// axios.interceptors.request.use(config => {
//  config.headers.Authorization = 'Bearer ' + window.sessionStorage.getItem('token')
//  return config
// })

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
