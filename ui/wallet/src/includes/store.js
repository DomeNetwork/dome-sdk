import { defineStore } from 'pinia'

export const appStore = defineStore('app', {
  actions: {
    clear() {
      this.domain = 'domenetwork.io'
      this.email = ''
      this.isLoggedIn = false
      this.publicKey = ''
      this.username = ''
    },
    login() {
      this.isLoggedIn = true
    },
    logout() {
      this.isLoggedIn = false
    },
    ready() {
      this.isReady = true
    },
    register(user) {
      this.email = user.email
      this.publicKey = user.publicKey
      this.username = user.username
    },
    unready() {
      this.isReady = false
    },
  },
  getters: {
    user() {
      return {
        domain: this.domain,
        email: this.email,
        publicKey: this.publicKey,
        username: this.username,
      }
    },
  },
  state: () => {
    return {
      balance: 0,
      domain: 'domenetwork.io',
      email: '',
      gas: 0,
      isLoggedIn: false,
      isReady: false,
      publicKey: '',
      username: '',
    }
  },
})
