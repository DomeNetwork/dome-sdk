
<script>
import { appStore } from '@/includes/store'
import Header from '@/components/Header.vue'
import Menu from '@/components/Menu.vue'

export default {
  components: {
    Header,
    Menu,
  },
  methods: {
    addressUrl(name) {
      return `https://rinkeby.etherscan.io/address/${name}`
    },
    async generate() {
      try {
        this.index++
        const acct = await dome.generate('Ethereum')
        await dome.key(JSON.stringify({
          name: acct.address,
          publicKey: acct.publicKey,
          used: false,
        }))

        // await dome.track('Ethereum', addr)

        await this.refresh()
      } catch (err) {
        console.error(err)
      }
    },
    async refresh() {
      try {
        const keys = await dome.keys()

        this.index = keys.length
        this.keys = keys
        this.$forceUpdate()

        await dome.write('dome-keys', JSON.stringify({
          index: this.index,
          keys: this.keys,
        }))
      } catch (err) {
        console.error(err)
      }
    },
  },
  async mounted() {
    if (!this.app.isLoggedIn) {
      await this.$router.replace('/')
      return
    }

    await this.refresh()

    // for (let i = 0; i < this.keys.length; i++) {
    //   await dome.track('Ethereum', this.keys[i].name)
    // }
  },
  async setup() {
    const app = appStore()

    let index = 0
    let keys = []

    try {
      const keyInfo = JSON.parse(await dome.read('dome-keys'))
      index = keyInfo.index
      keys = keyInfo.keys
    } catch (err) {
      // Do nothing, this will error if the value is not set yet.
    }

    return {
      app,
      index,
      keys,
    }
  },
}
</script>

<template>
  <div class="flex-row full-size">
    <Menu />
    <div class="flex-col flex-full body">
      <Header section="Keys" />
      <div class="flex-row flex-space">
        <div class="card keys">
          <button class="btn btn-success" @click="generate()">Generate Key</button>
          <h3>Keys:</h3>
          <div class="list">
            <table class="table table-hover">
              <thead>
                <tr>
                  <th scope="col">Account</th>
                  <th scope="col">Used</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="key in keys" :key="key.publicKey">
                  <th scope="row">
                    <a v-if="key.used" :href="addressUrl(key.name)" target="_blank">{{ key.name }}</a>
                    <span v-else>{{ key.name }}</span>
                  </th>
                  <td>{{ key.used }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.keys {
  margin-top: 1vh;
  width: 88vw;
}

.keys button {
  max-width: 15vw;
}

.keys h3 {
  margin-bottom: 1vh;
  margin-top: 2vh;
}
</style>
