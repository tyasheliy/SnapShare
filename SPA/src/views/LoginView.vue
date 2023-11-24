<template>
    <div>
        <Navbar />
        <div class="frame bg-secondary h-screen flex flex-col justify-center items-center">
            <div class="cont lg:w-1/2 w-full h-2/3 flex flex-col items-center">
                <div class="h-1/3 flex items-center">
                    <span class="text-foreground text-4xl font-bold">Log in</span>
                </div>
                <div class="h-2/3 flex flex-col items-center space-y-2">
                    <div id="messageContainer" ref="messageContainer"
                        class="text-foreground text-center rounded-sm opacity-0 transition-all duration-300 w-full">
                        <span id="messageSpan" ref="messageSpan" class="inline-block w-2/3 p-1">
                            error message
                        </span>
                    </div>
                    <form @submit.prevent="submit()" class="flex flex-col space-y-2 w-48">
                        <div class="h-1/3 flex flex-col space-y-3">
                            <input ref="usernameInput" placeholder="Username"
                                class="text-center p-2 border-2 text-foreground focus:text-primary border-foreground rounded-lg hover:border-accent bg-primary focus:bg-foreground transition-all duration-150"
                                v-model="username" type="text" />
                            <input ref="passwordInput" placeholder="Password"
                                class="text-center p-2 border-2 text-foreground focus:text-primary border-foreground rounded-lg hover:border-accent bg-primary focus:bg-foreground transition-all duration-150"
                                v-model="password" type="password" />
                        </div>
                        <div class="h-1/3  text-foreground text-center flex-col flex">
                            <button ref="submitButton" id="submitButton"
                                class="m-10 rounded-xl p-3 border-2 border-foreground hover:bg-accent focus:bg-accept transition-all duration-150"
                                type="submit">Submit</button>
                            <div class="flex flex-col items-center space-y-2">
                                <router-link to="/register" class="text-accent hover:underline w-16">Register</router-link>
                                <router-link to="/forgot" class="text-accent hover:underline">Forgot password</router-link>
                            </div>
                        </div>

                    </form>
                </div>
                <Footer />
            </div>
        </div>
    </div>
</template>

<script>
import Footer from '@/components/Footer.vue';
import Navbar from '@/components/Navbar.vue';
import axios from 'axios';

export default {
    components: {
        Navbar,
        Footer
    },
    methods: {
        showErrorMessage(message) {
            this.$refs.messageSpan.innerText = message

            this.$refs.messageContainer.classList.add("bg-error")
            this.$refs.messageContainer.classList.remove("opacity-0")

            this.$refs.submitButton.classList.add("mt-4")
        },
        showAcceptedMessage(message) {
            this.$refs.messageSpan.innerText = message
            this.$refs.messageContainer.classList.add("bg-accept")
            this.$refs.messageContainer.classList.remove("opacity-0")
        },
        clearForm() {
            let cont = this.$refs.messageContainer

            cont.classList.add("opacity-0")
            cont.classList.forEach((c) => {
                if (c.startsWith("bg-")) {
                    cont.classList.remove(c)
                }
            })

            this.$refs.usernameInput.classList.remove("border-error")
            this.$refs.passwordInput.classList.remove("border-error")
        },
        submit() {
            this.clearForm()

            let empty = false

            if (this.username === undefined) {
                this.$refs.usernameInput.classList.remove("border-foreground")
                this.$refs.usernameInput.classList.add("border-error")
                empty = true
            }

            if (this.password === undefined) {
                this.$refs.passwordInput.classList.remove("border-foreground")
                this.$refs.passwordInput.classList.add("border-error")
                empty = true
            }

            if (empty) {
                this.showErrorMessage("Fill in the fields!")
                return
            }

            let data = {
                "username": this.username,
                "password": this.password
            }

            axios.post("identity/auth", data)
                .then((response) => {
                    localStorage.token = response.data.token
                    this.$router.push({ name: "home" })
                })
                .catch((error) => {
                    let message = error.response.data.message.toLowerCase()

                    if (message.includes("password") || message.includes("credentials")) {
                        this.$refs.passwordInput.classList.remove("border-foreground")
                        this.$refs.passwordInput.classList.add("border-error")
                    }

                    if (message.includes("username") || message.includes("credentials")) {
                        this.$refs.usernameInput.classList.remove("border-foreground")
                        this.$refs.usernameInput.classList.add("border-error")
                    }

                    this.showErrorMessage(error.response.data.message)
                })

        }
    }
}
</script>