import Vue from 'vue'

import Cookies from 'js-cookie'

import 'normalize.css/normalize.css' // a modern alternative to CSS resets

import Element from 'element-ui'
import './styles/element-variables.scss'

import '@/styles/index.scss' // global css

import App from './App'
import store from './store'
import router from './router'

import './icons' // icon
// import './permission' // permission control
import './utils/error-log' // error log

import * as filters from './filters' // global filters

/**
 * If you don't want to use mock-server
 * you want to use MockJs for mock api
 * you can execute: mockXHR()
 *
 * Currently MockJs will be used in the production environment,
 * please remove it before going online ! ! !
 */
if (process.env.NODE_ENV === 'production') {
  const { mockXHR } = require('../mock')
  mockXHR()
}

Vue.use(Element, {
  size: Cookies.get('size') || 'medium' // set element-ui default size
})

// register global utility filters
Object.keys(filters).forEach(key => {
  Vue.filter(key, filters[key])
})

Vue.config.productionTip = false

new Vue({
  el: '#app',
  router,
  store,
  created: function() {
    const that = this
    /* eslint-disable */
    let jump = function(ele, e) {
      let href = ele.attr('href');
      const now_url = document.location.href;
      const url_prefix = '/admin/vue';
      href = href.replace(window.location.protocol + "//" + window.location.host, "");
      if (href.indexOf(url_prefix) !== -1 && now_url.indexOf(url_prefix) !== -1) {
        e.preventDefault();
        that.$router.push({ path: href.replace(url_prefix, '') })
      }
    };
    $('.sidebar a').on('click', function(e) {
      jump($(this), e)
    });
    $('.nav.nav-tabs a').on('click', function(e) {
      jump($(this), e)
    })
  },
  render: h => h(App)
})
