
<script>
import { appStore } from '@/includes/store'

export default {
  props: {
    section: String,
  },
  async setup() {
    const { address } = await dome.account('Ethereum')
    const app = appStore()

    return {
      address,
      addressUrl: `https://rinkeby.etherscan.io/address/${address}`,
      app,
    }
  },
}
</script>

<template>
  <div class="flex-row flex-space">
    <div class="card flex-row header">
      <div class="welcome">
        Welcome <span>{{ app.username }}@{{ app.domain }}</span>
      </div>
      <div class="address">
        <a :href="addressUrl" target="_blank">{{ address }}</a>
      </div>
      <div class="flex-full"></div>
      <div class="welcome">
        {{ section }}
      </div>
    </div>
  </div>
</template>

<style>
span {
  color: #0000ff;
}

.header {
  font-size: 1.5rem;
  margin-bottom: 2vh;
  margin-top: 2vh;
  width: 88vw;
}

.header .address {
  margin-left: 1rem;
}
</style>
