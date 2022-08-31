
<script>
import { appStore } from '@/includes/store'
import { password, secret, username } from '@/includes/regex'
import { ref } from 'vue'

export default {
  computed: {
    isValid() {
      if (this.step === 1) {
        return this.verify === this.words;
      } else if (this.step === 2) {
        return username.test(this.username) && password.test(this.password)
      } else if (this.step === 3) {
        return secret.test(this.secret)
      }
      return true
    },
  },
  data: () => ({
    password: '',
    secret: '',
    username: '',
    verify: '',
  }),
  methods: {
    login() {
      this.$router.push('/')
    },
    async next() {
      if (this.step < 4) {
        if (this.step === 1) {
          // Do nothing
        } else if (this.step === 2) {
          await dome.load(this.words)

          const user = {
            domain: 'domenetwork.io',
            email: `${this.username}@${this.app.domain}`,
            publicKey: await dome.publicKey(),
            username: this.username,
          }
          this.app.register(user)

          user.password = this.password
          await dome.register(JSON.stringify(user))
        } else if (this.step === 3) {
          await dome.save(this.secret)

          const user = this.app.user
          user.password = await dome.encrypt(this.secret, this.password)
          await dome.write('dome-user', JSON.stringify(user))
        }

        this.step++
      }
    },
    prev() {
      if (this.step > 0) {
        this.step--
      }
    },
  },
  async setup() {
    const app = appStore()
    const step = ref(0)

    const words = await dome.mnemonic()

    return {
      app,
      step,
      words,
    }
  },
}
</script>

<template>
  <div class="flex-col flex-center full-size">
    <div class="card card-full flex-col flex-center">
      <div v-if="step === 0" class="flex-col">
        <h2>Mnemonic</h2>
        <div class="error">
          Write these words below down!<br />
          This will allow you to recover your wallet in case of corruption or lockout.
        </div>
        <div class="words">{{ words }}</div>
      </div>
      <div v-if="step === 1" class="flex-col">
        <h2>Verify Mnemonic</h2>
        <div>Provide the words from the previous screen in the exact same order.</div>
        <div class="form-group">
          <label>Mnemonic:</label>
          <textarea :class="{ 'form-control': true, 'is-valid': isValid }" rows="5" v-model="verify"></textarea>
        </div>
      </div>
      <div v-if="step === 2" class="flex-col">
        <h2>Register User</h2>
        <div>To continue the registration process, please provide a username and password.</div>
        <label>Username:</label>
        <input :class="{ 'form-control': true, 'is-valid': isValid }" placeholder="Username" type="text" v-model="username" />
        <label>Password:</label>
        <input :class="{ 'form-control': true, 'is-valid': isValid }" placeholder="Password" type="password" v-model="password" />
      </div>
      <div v-if="step === 3" class="flex-col">
        <h2>Secure Data</h2>
        <div>Provide a secure PIN, 4-6 digits, to encrypt your local wallet.  Even if the wallet is stolen your data will remain locked and requires a secret to unlock it.</div>
        <label>Secret:</label>
        <input :class="{ 'form-control': true, 'is-valid': isValid }" placeholder="Secret" type="password" v-model="secret" />
      </div>
      <div v-if="step === 4" class="flex-col">
        <h2>Congratulations!</h2>
        <div>You have finished the setup of your wallet.  Now login to continue to the home page.</div>
      </div>
      <div class="flex-full"></div>
      <div class="flex-row flex-space">
        <button v-if="step === 0" class="btn btn-secondary btn-lg" @click="$router.push('/')">Login to Wallet</button>
        <button v-if="step > 0 && step < 4" class="btn btn-secondary btn-lg" @click="prev()">Back</button>
        <div v-if="step < 4" class="flex-full"></div>
        <button v-if="step < 4" class="btn btn-primary btn-lg" @click="next()" :disabled="!isValid">Continue</button>
        <button v-if="step === 4" class="btn btn-success btn-lg" @click="login()" :disabled="!isValid">Continue to Login</button>
      </div>
    </div>
  </div>
</template>

<style>
h2 {
  text-align: center;
  width: 680px;
}

.error {
  color: #ff0000;
  font-size: 1.5rem;
  text-align: center;
}

.words {
  color: #000000;
  font-size: 2rem;
  font-weight: bold;
  margin-top: 2rem;
  text-align: center;
}
</style>
