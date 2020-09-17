const ALERT_DURATION_MS = 5000

const state = () => ({
  alert: null,
  timeoutId: null
})

const mutations = {
  setAlert (state, alert) {
    state.alert = alert
  },

  setTimeoutId (state, timeoutId) {
    state.timeoutId = timeoutId
  }
}

const actions = {
  async createAlert ({ commit, state }, alert) {
    if (state.timeoutId !== null) {
      clearTimeout(state.timeoutId)
      commit('setTimeoutId', null)
    }

    commit('setAlert', alert)

    commit('setTimeoutId', setTimeout(() => {
      commit('setAlert', null)
      commit('setTimeoutId', null)
    }, ALERT_DURATION_MS))
  },

  async createAlertFromError ({ dispatch }, error) {
    dispatch('createAlert', {
      type: 'error',
      message: 'Une erreur est survenue, veuillez réessayer.'
    })
  }
}

export default {
  namespaced: true,
  state,
  actions,
  mutations
}
