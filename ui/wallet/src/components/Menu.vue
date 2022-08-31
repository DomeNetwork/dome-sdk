
<script>
import { appStore } from '@/includes/store'

export default {
  methods: {
    go(route) {
      this.$router.push(route)
    },
    logout() {
      this.app.clear()
      this.$router.push('/')
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
  <div class="flex-col menu bg-primary">
    <h5>Menu</h5>
    <a @click="go('/home')" href="#">Home</a>
    <a @click="go('/keys')" href="#">Keys</a>
    <a @click="go('/send')" href="#">Send</a>
    <div class="line"></div>
    <a @click="logout()" href="#">Logout</a>
  </div>
</template>

<style>
a {
  margin-top: 1rem;
  text-align: center;
}

h5 {
  margin: 1rem;
}

.menu {
  /* background-color: #ffffff;
  box-shadow: 0 0 1.5rem rgba(0, 0, 0, 0.25); */
  width: 10vw;
}

.line {
  /* background-color: #dddddd; */
  height: 1px;
  margin-top: 1vh;
}
</style>
