
<script>
import { appStore } from '@/includes/store'
import { secret } from '@/includes/regex'

export default {
  data: () => ({
    secret: '',
  }),
  computed: {
    isValid() {
      return secret.test(this.secret)
    },
  },
  methods: {
    async login() {
      try {
        await dome.open(this.secret)

        const user = JSON.parse(await dome.read('dome-user'))
        user.password = await dome.decrypt(this.secret, user.password)

        await dome.login(JSON.stringify({
          domain: user.domain,
          password: user.password,
          username: user.username,
        }))

        this.app.register(user)
        this.app.login()

        this.$router.push('/home')
      } catch (err) {
        console.error(err)
      }
    },
    register() {
      this.$router.push('/register')
    },
  },
  mounted() {
    if (!this.found) {
      this.$router.push('/register')
    }
  },
  async setup() {
    const app = appStore()
    const found = await dome.check()

    return {
      app,
      found,
    }
  },
}
</script>

<template>
  <div class="flex-col flex-center full-size">
    <div class="card card-full flex-col flex-center">
      <div class="flex-col">
        <h2>Mnemonic</h2>
        <div>Provide your secure PIN, 4-6 digits, to decrypt your local wallet.</div>
        <label>Secret:</label>
        <input :class="{ 'form-control': true, 'is-valid': isValid }" placeholder="Secret" type="password" v-model="secret" />
      </div>
      <div class="flex-full"></div>
      <div class="flex-col flex-space">
        <button class="btn btn-success btn-lg" @click="login()" :disabled="!isValid">Login to Wallet</button>
        <button class="btn btn-secondary btn-sm" @click="register()">Register</button>
      </div>
    </div>
  </div>
</template>

<style>
h2 {
  text-align: center;
  width: 680px;
}
</style>
