<template>
    <div class="frame h-screen bg-secondary flex items-center justify-center">
        <Navbar />
        <div class="cont lg:w-1/2 w-full h-5/6 flex flex-col">
            <div class="h-1/5 flex items-center justify-center text-foreground">
                <span class="text-3xl font-bold">Register</span>
            </div>
            <form @submit.prevent="submit()" class="h-4/5 flex flex-col items-center text-foreground">
                <span ref="messageContainer" class="m-4 opacity-0 p-4 rounded-sm transition-all duration-150">Create a free
                    unlimited account</span>
                <div class="flex flex-col items-center space-y-3">
                    <input ref="usernameInput" v-model="username" type="text" placeholder="Username"
                        class="text-center p-2 border-2 text-foreground focus:text-primary border-foreground rounded-lg hover:border-accent bg-primary focus:bg-foreground transition-all duration-150" />
                    <input ref="emailInput" v-model="email" type="text" placeholder="Email"
                        class="text-center p-2 border-2 text-foreground focus:text-primary border-foreground rounded-lg hover:border-accent bg-primary focus:bg-foreground transition-all duration-150" />
                    <input ref="passwordInput" v-model="password" type="password" placeholder="Password"
                        class="text-center p-2 border-2 text-foreground focus:text-primary border-foreground rounded-lg hover:border-accent bg-primary focus:bg-foreground transition-all duration-150" />
                    <input ref="confirmPasswordInput" v-model="confirmPassword" type="password"
                        placeholder="Confirm password"
                        class="text-center p-2 border-2 text-foreground focus:text-primary border-foreground rounded-lg hover:border-accent bg-primary focus:bg-foreground transition-all duration-150" />
                </div>
                <button ref="submitButton"
                    class="mt-auto m-12 rounded-xl p-3 border-2 border-foreground hover:bg-accent focus:bg-accept transition-all duration-150"
                    type="submit">Submit</button>
            </form>
            <Footer />
        </div>
    </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import Footer from '@/components/Footer.vue';
import axios from 'axios';
import { inputError, clearInput } from '@/global/global';

export default {
    components: {
        Navbar,
        Footer
    },
    methods: {
        showErrorMessage(message) {
            this.$refs.messageContainer.classList.remove("opacity-0")
            this.$refs.messageContainer.classList.add("bg-error")
            this.$refs.messageContainer.innerText = message
        },
        hideMessage() {
            this.$refs.messageContainer.classList.add("opacity-0")
            this.$refs.messageContainer.classList.forEach((c) => {
                if (c.startsWith("bg-")) {
                    this.$refs.messageContainer.classList.remove(c)
                }
            })

        },
        clearForm() {
            this.hideMessage()

            clearInput(this.$refs.usernameInput)
            clearInput(this.$refs.emailInput)
            clearInput(this.$refs.passwordInput)
            clearInput(this.$refs.confirmPasswordInput)
        },
        submit() {
            this.clearForm()

            let username = this.$refs.usernameInput
            let email = this.$refs.emailInput
            let password = this.$refs.passwordInput
            let confirmPassword = this.$refs.confirmPasswordInput

            let empty = false

            if (this.username === undefined || this.username === "") {
                inputError(username)
                empty = true
            }

            if (this.email === undefined || this.email === "") {
                inputError(email)
                empty = true
            }

            if (this.password === undefined || this.password === "") {
                inputError(password)
                empty = true
            }

            if (this.confirmPassword === undefined || this.password === "") {
                inputError(confirmPassword)
                empty = true
            }

            if (empty) {
                this.showErrorMessage("Fill in the fields!")
                return
            }

            if (this.password != this.confirmPassword) {
                this.showErrorMessage("Passwords mismatch!")
                inputError(password)
                inputError(confirmPassword)
                return
            }

            let data = {
                "username": this.username,
                "email": this.email,
                "password": this.password
            }

            axios.post("identity/users", data)
                .then((response) => {
                    this.$router.push({name: "login"})
                })
                .catch((error) => {
                    let message = error.response.data.message

                    if (message === undefined || message === null) {
                        this.showErrorMessage("Unknown server error.")
                        return
                    }

                    this.showErrorMessage(message)

                    if (message.toLowerCase().includes("email")) {
                        inputError(email)
                    }

                    if (message.toLowerCase().includes("username")) {
                        inputError(username)
                    }

                    if (message.toLowerCase().includes("password")) {
                        inputError(password)
                    }
                })
        }
    }
}
</script>