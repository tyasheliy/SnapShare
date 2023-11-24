import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './assets/tailwind.css'
import axios from 'axios'

const host = "176.214.93.42"
const debugHost = "localhost"

axios.defaults.baseURL = `http://${host}/api/`

createApp(App).use(router).mount('#app')