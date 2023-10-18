import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './assets/tailwind.css'
import axios from 'axios'

axios.defaults.baseURL = "http://localhost/api/"

createApp(App).use(router).mount('#app')



//custom functions
function checkAuth() {
    if (localStorage.token === undefined) {
        let promise = new Promise((res, rej) => setTimeout(() => {
            res()
        }, 200))
        promise.then(() => redirectWithUnauthError())
    }
}

function redirectWithUnauthError() {
    console.log("Token has expired or unauthenticated")

    let messageContainer = document.querySelector("#messageContainer")
    let messageSpan = document.querySelector("#messageSpan")

    if (messageContainer === undefined || messageSpan === undefined) {
        return
    }

    messageContainer.classList.add("bg-error")
    messageContainer.classList.remove("opacity-0")

    messageSpan.innerText = ""

    this.$router.push({ name: "login" })
}