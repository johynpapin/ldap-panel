import router from '@/router'
import DataService from '@/services/data-service'

const state = {
  isAuthenticated: false
}

const mutations = {
  setIsAuthenticated (state, isAuthenticated) {
    state.isAuthenticated = isAuthenticated
  }
}

const actions = {
  async signIn ({ commit, dispatch }, { username, password }) {
    await DataService.signIn(username, password)
    commit('setIsAuthenticated', true)

    router.push({
      name: 'Home'
    })

    dispatch('alerts/createAlert', {
      type: 'success',
      message: 'Vous êtes bien connecté.'
    }, {
      root: true
    })
  },

  async signOut ({ commit, dispatch }) {
    await DataService.signOut()
    commit('setIsAuthenticated', false)

    router.push({
      name: 'Home'
    })

    dispatch('alerts/createAlert', {
      type: 'success',
      message: 'Vous êtes bien déconnecté.'
    }, {
      root: true
    })
  },

  async updatePassword ({ dispatch }, { password }) {
    await DataService.updatePassword(password)

    dispatch('alerts/createAlert', {
      type: 'success',
      message: 'Votre mot de passe a bien été mis à jour.'
    }, {
      root: true
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
