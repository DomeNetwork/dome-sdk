
<script>
import { appStore } from '@/includes/store'
import { RouterView } from 'vue-router'
import Splash from '@/containers/Splash.vue'

export default {
  components: {
    RouterView,
    Splash,
  },
  computed: {
    isReady() {
      return this.app.isReady
    },
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
  <Suspense>
    <div class="full-size">
      <Splash v-if="!isReady" />
      <RouterView v-if="isReady" />
    </div>
  </Suspense>
</template>

<style>
@import '@/assets/css/quartz-bootstrap.min.css';
@import '@/assets/css/main.css';
</style>
