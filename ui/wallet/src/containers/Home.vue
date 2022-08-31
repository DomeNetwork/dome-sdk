
<script>
import { appStore } from '@/includes/store'
import Amount from '@/components/Amount.vue'
import Header from '@/components/Header.vue'
import Menu from '@/components/Menu.vue'

const chainName = 'Ethereum'

export default {
  components: {
    Amount,
    Header,
    Menu,
  },
  methods: {
    async refresh() {
      try {
        this.app.balance = await dome.balance(chainName)
        this.app.gas = await dome.gas(chainName)
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

    try {
      // const keyInfo = JSON.parse(await dome.read('dome-keys'))
      // for (let i = 0; i < keyInfo.keys.length; i++) {
      //   await dome.track('Ethereum', keyInfo.keys[i].name)
      // }

      this.refresh()
    } catch (err) {
      console.error(err)
    }
  },
  setup() {
    const app = appStore()

    return {
      app,
    }
  },
}
</script>

<template>
  <div class="flex-row full-size">
    <Menu />
    <div class="flex-col flex-full body">
      <Header section="Home" />
      <div class="flex-row flex-space">
        <div class="card balance">
          <div class="flex-row">
            <div class="sub-title">Balance:</div>
            <div class="flex-full"></div>
            <a @click="refresh()" href="#">Refresh</a>
          </div>
          <span>
            <Amount :v="app.balance" />
          </span>
        </div>
        <div class="card gas">
          <div class="flex-row">
            <div class="sub-title">Gas Estimate:</div>
            <div class="flex-full"></div>
            <a @click="refresh()" href="#">Refresh</a>
          </div>
          <span>
            <Amount :v="app.gas" />
          </span>
        </div>
      </div>
      <div class="flex-row flex-space">
        <div class="card news">
          <h1 class="title">
            ḎOME Network
          </h1>
          <div class="article">
            <p>
              Layer 1 blockchain optimized for the next billion users
            </p>
            <p>
              A scalable, vertically integrated ecosystem that focuses exclusively on making crypto less
              cryptic – introducing Web3 to the World. The first vertically integrated layer 1 ecosystem,
              with user experience as a tentpole priority. Dome offers infrastructure services, unified
              developer tools and an interoperable first-party wallet.
            </p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.body {
  margin-top: 1vh;
  max-height: 100%;
  overflow-x: hidden;
  overflow-y: auto;
}

.refresh {

}

.balance {
  font-size: 2rem;
  width: 43vw;
}

.gas {
  font-size: 2rem;
  width: 43vw;
}

.news {
  margin-top: 2vh;
  width: 88vw;
}
</style>
