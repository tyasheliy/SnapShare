<template>
    <div>
        <Navbar/>
        <div class="frame h-screen bg-secondary flex items-center justify-center">
            <div class="cont lg:w-1/2 w-screen h-4/5 flex flex-col items-center">
                <div class="h-1/4 flex items-center text-foreground justify-center w-full">
                    <span class="text-3xl font-bold">Profile</span>
                </div>
                <div class="h-3/4 w-full text-foreground flex lg:flex-row flex-col">
                    <div class="lg:w-1/2 lg:h-full h-1/2 flex justify-center">
                        <div class="flex flex-col text-center items-center">
                            <span class="m-4 text-xl font-semibold">User info</span>
                            <div class="flex flex-col items-center space-y-2">
                                <span>{{ this.userData.userName }}</span>
                                <span>{{ this.userData.email }}</span>
                                <span>{{ this.userData.isEmailConfirmed ? "Email is confirmed" : "Email is not confirmed"  }}</span>
                            </div>
                            <button @click="logout()" class="m-8 w-1/2 rounded-xl p-3 border-2 border-foreground hover:bg-accent focus:bg-accept transition-all duration-150">Logout</button>
                        </div>
                    </div>
                    <div class="lg:w-1/2 lg:h-full h-1/2 flex flex-col items-center">
                        <span class="text-xl font-semibold m-4">Change password</span>
                        <form class="w-full h-full flex flex-col items-center" @submit.prevent="changePassword()">
                            <div class="h-2/3 flex flex-col items-center space-y-4">
                                <input ref="oldPasswordInput" v-model="oldPassword" placeholder="Old Password" type="password" class="text-center p-2 border-2 text-foreground focus:text-primary border-foreground rounded-lg hover:border-accent bg-primary focus:bg-foreground transition-all duration-150"/>
                                <input ref="newPasswordInput" v-model="newPassword" placeholder="New Password" type="password" class="text-center p-2 border-2 text-foreground focus:text-primary border-foreground rounded-lg hover:border-accent bg-primary focus:bg-foreground transition-all duration-150"/>
                                <button type="submit" class="rounded-xl p-3 border-2 border-foreground hover:bg-accent focus:bg-accept transition-all duration-150">Change</button>
                            </div>
                        </form>
                    </div>
                </div>
                <div ref="messageContainer" class="max-w-md flex items-center justify-center text-foreground rounded-sm transition-all duration-150 opacity-0">
                    <span ref="messageSpan" class="p-2">error message</span>
                </div>
                <Footer/>
            </div>
        </div>
    </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import Footer from '@/components/Footer.vue';
import { checkAuth, inputError, clearInput } from '../global/global.js'
import axios from 'axios';

export default {
    components: {
        Navbar,
        Footer
    },
    data() {
        return {
            userData: {
                userName: "",
                email: "",
                isEmailConfirmed: false
            }
        }
    },
    created() {
        checkAuth(this.$router)

        let data = {headers: {"Authorization": `Bearer ${localStorage.token}`}}
        axios.get("identity/users?intention=data", data)
            .then((response) => this.userData = response.data)
            .catch((error) => {
                localStorage.removeItem("token")
                checkAuth(this.$router)
            })
    },
    methods: {
        showErrorMessage(message) {
            this.$refs.messageContainer.classList.remove("opacity-0")
            this.$refs.messageContainer.classList.add("bg-error")

            this.$refs.messageSpan.innerText = message
        },
        showAcceptedMessage(message) {
            this.$refs.messageContainer.classList.remove("opacity-0")
            this.$refs.messageContainer.classList.add("bg-accept")

            this.$refs.messageSpan.innerText = message
        },
        hideMessage() {
            let container = this.$refs.messageContainer

            container.classList.add("opacity-0")

            container.classList.forEach((c) => {
                if (c.startsWith("bg-")) {
                    container.classList.remove(c)
                }
            })
        },
        clearForChangingPasswords() {
            this.hideMessage()
            clearInput(this.$refs.oldPasswordInput)
            clearInput(this.$refs.newPasswordInput)
        },
        changePassword() {
            this.clearForChangingPasswords()

            let oldPassword = this.$refs.oldPasswordInput
            let newPassword = this.$refs.newPasswordInput

            let empty = false

            if (this.oldPassword === undefined || this.oldPassword === "") {
                inputError(oldPassword)
                empty = true
            }

            if (this.newPassword === undefined || this.newPassword === "") {
                inputError(newPassword)
                empty = true
            }

            if (empty) {
                this.showErrorMessage("Fill in the fields!")
                return
            }

            if (this.oldPassword === this.newPassword) {
                this.showErrorMessage("Passwords match!")
                inputError(oldPassword)
                inputError(newPassword)
                return
            }

            let data = {
                "username": this.userData.userName,
                "oldPassword": this.oldPassword,
                "password": this.newPassword
            }

            axios.patch("identity/users", data)
                .then((response) => {
                    this.showAcceptedMessage("Password changed successfully")
                })
                .catch((error) => {
                    let message = "Unknown server error"

                    if (error.response.data.message != undefined || error.response.data.message != null) {
                        message = error.response.data.message
                    }

                    if (error.response.data.title != undefined || error.response.data.title != null) {
                        message = error.response.data.title
                    }

                    if (message.toLowerCase().includes("password")) {
                        inputError(this.$refs.oldPasswordInput)
                    }
                    
                    this.showErrorMessage(message)
                })
        },
        logout() {
            localStorage.removeItem("token")
            this.$router.push({name: "login"})
        },
    }
}
</script>