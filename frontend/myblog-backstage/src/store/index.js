import Vue from 'vue'
import Vuex from 'vuex'
import getters from './getter'
import common from './modules/common'
import user from './modules/user'
import tags from './modules/tag'
Vue.use(Vuex)

export default new Vuex.Store({
  state: {
  },
  getters,
  mutations: {
  },
  actions: {
  },
  modules: {
    common,
    user,
    tags
  }
})
