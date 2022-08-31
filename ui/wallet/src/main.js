import { createApp } from 'vue'
import { createPinia } from 'pinia'

import ToastPlugin from 'vue-toast-notification'
import 'vue-toast-notification/dist/theme-default.css'

import App from '@/containers/App.vue'
import router from '@/includes/router'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)
app.use(router)
app.use(ToastPlugin)

app.mount('#app')
