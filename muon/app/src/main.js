import Vue from 'vue'
import VueRouter from 'vue-router'
import Vuex from 'vuex'
import VueI18n from 'vue-i18n'
import App from './App.vue'
import routes from '@/routes'
import mixin from '@/mixin'
import Store from '@/store'
import Configuration from '@/config'
import '@/assets/css/style.css'


async function loadI18n () {
  const i18n = await Promise.all(Configuration.ui.languages.map(async lang => ({
    code: lang,
    messages: await fetch(`${Configuration.url.static}/i18n/${lang}.json`).then(res => res.json())
  })))
  return i18n.reduce((total, current) => ({
    ...total,
    [current.code]: current.messages
  }), {})
}

async function initApplication () {
  Vue.config.productionTip = false
  Vue.use(VueRouter)
  Vue.use(Vuex)
  Vue.use(VueI18n)
  Vue.mixin(mixin)
  
  const router = new VueRouter({
    routes
  })
  
  const store = new Vuex.Store(Store)
  
  const i18n = new VueI18n({
    locale: Configuration.ui.languages[0],
    messages: await loadI18n()
  })
  
  window.app = new Vue({
    router,
    store,
    i18n,
    render: h => h(App),
  })
  window.app.$mount('#app')
}

window.onload = initApplication
