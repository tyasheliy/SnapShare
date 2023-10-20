import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './assets/tailwind.css'
import axios from 'axios'

const host = "192.168.0.100"
const debugHost = "localhost"

axios.defaults.baseURL = `http://${debugHost}/api/`

createApp(App).use(router).mount('#app')