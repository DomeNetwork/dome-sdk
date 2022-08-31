
<script>
import { appStore } from '@/includes/store'
import config from '@/includes/config'

export default {
  methods: {
    async load() {
      if (!!window.dome) {
        // this.$router.push('/register')
        await dome.config(JSON.stringify(config))

        this.app.ready()
      } else {
        setTimeout(this.load, 3000)
      }
    },
  },
  mounted() {
    setTimeout(this.load, 5000)
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
  <div id="overlay" class="flex-col flex-center full-size">
    <div class="card card-full">
      <video muted="" autoplay="" playsinline="" id="hero-video">
        <source src="/1848292884454272583d-square_1_700px.mp4" type="video/mp4" media="all and (max-width: 700px)">
      </video>
    </div>
  </div>
</template>

<style>
#overlay {
  background-color: #ffffff;
}

#hero-video {
  background: "/6286975b1361243853266525_d-png-bg3.jpeg";
  left: 0;
  position: absolute;
  top: 0;
  width: 720px;
}
</style>
