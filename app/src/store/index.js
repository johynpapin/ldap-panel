import Vue from 'vue'
import Vuex from 'vuex'
import ErrorsPlugin from './errors-plugin'
import alerts from './modules/alerts'
import data from './modules/data'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    alerts,
    data
  },
  plugins: [ErrorsPlugin]
})
