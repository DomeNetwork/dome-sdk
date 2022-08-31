
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
    async search() {
      try {
        const results = await dome.search(JSON.stringify({ term: this.term }))
        this.results = Array.isArray(results) ? results : [results]
        this.tx = ''
        this.$forceUpdate()
      } catch (err) {
        console.error(err)
      }
    },
    async select(username) {
      try {
        const addr = await dome.lookup(JSON.stringify({
          domain: this.app.domain,
          username,
        }))
        this.to = addr
        this.$forceUpdate()
      } catch (err) {
        console.error(err)
      }
    },
    async send() {
      try {
        const txHash = await dome.send('Ethereum', this.to, this.amount, this.data)
        this.tx = `https://rinkeby.etherscan.io/tx/${txHash}`

        this.amount = 0
        this.data = ''
        this.to = ''

        this.$forceUpdate()
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
  },
  async setup() {
    const app = appStore()

    return {
      amount: '0',
      app,
      data: '',
      results: [],
      term: '',
      to: '',
      tx: '',
    }
  },
}
</script>

<template>
  <div class="flex-row full-size">
    <Menu />
    <div class="flex-col flex-full body">
      <Header section="Send" />
      <div class="flex-col flex-center">
        <div class="card search">
          <div class="flex-row flex-space">
            <label>Username:</label>
            <div class="spacer-w"></div>
            <input class="form-control" placeholder="Username" type="text" v-model="term" />
            <div class="spacer-w"></div>
            <button class="btn btn-success btn-md" @click="search()">Search</button>
          </div>
          <h3 v-if="results.length > 0">Search results:</h3>
          <div v-if="results.length > 0" class="list">
            <table class="table table-hover">
              <thead>
                <tr>
                  <th scope="col">User Handle</th>
                  <th></th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="handle in results" :key="handle">
                  <th scope="row">{{ handle }}</th>
                  <th>
                    <a @click="select(handle)" href="#">Select</a>
                  </th>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
        <div class="flex-row flex-space to">
          <div class="card flex-row flex-space">
            <label>To:</label>
            <input class="form-control" placeholder="To Address" type="text" v-model="to" />
          </div>
          <div class="spacer-w"></div>
          <div class="card flex-row flex-space">
            <label>Data:</label>
            <input class="form-control" placeholder="Data" type="text" v-model="data" />
          </div>
        </div>
        <div class="card flex-row flex-space send">
          <div class="flex-row flex-space">
            <label>Amount (WEI):</label>
            <input class="form-control" placeholder="Amount" type="text" v-model="amount" />
          </div>
        </div>
        <div v-if="!!tx" class="card tx">
          View TX:<a :href="tx" target="_blank">{{ tx }}</a>
        </div>
        <button class="btn btn-success btn-lg" @click="send()">Send</button>
      </div>
    </div>
  </div>
</template>

<style>
label {
  margin-right: 1rem;
}

.spacer-h {
  height: 1vh;
}

.spacer-w {
  width: 1vw;
}

.search {
  margin-top: 1vh;
  width: 88vw;
}

.to {
  margin-top: 1vh;
  width: 88vw;
}

.send {
  margin-top: 1vh;
  width: 88vw;
}

.tx {
  margin-top: 1vh;
  width: 88vw;
}
</style>
