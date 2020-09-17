import axios from 'axios'

const client = axios.create({
  baseURL: process.env.VUE_APP_API_BASE_URL
})

export default {
  async signIn (username, password) {
    const {
      data
    } = await client.post('auth/sign-in', {
      username,
      password
    })

    client.defaults.headers.common.Authorization = `Bearer ${data.token}`
  },

  async signOut () {
    delete client.defaults.headers.common.Authorization
  },

  async updatePassword (password) {
    await client.post('update-password', {
      password
    })
  }
}
