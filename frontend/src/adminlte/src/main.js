// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import './lib/css'
import './lib/script'
import './lib/global'

import Vue from 'vue'
import App from './App'
import router from './router'
import EventBus from './lib/eventBus.js'
import axios from 'axios'

Vue.prototype.$bus = EventBus
Vue.prototype.$http = axios

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  created: function() {
    const that = this;
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
  components: {
    App
  }
})
