<template>
    <div class="frame h-screen bg-secondary flex items-center justify-center">
        <Navbar />
        <div class="cont lg:w-1/2 w-full h-2/3 flex flex-col items-center">
            <div class="h-1/3 flex items-center justify-center text-foreground">
                <span class="text-3xl font-bold">Share</span>
            </div>
            <div ref="messageContainer" class="text-foreground text-center w-1/2 m-4 opacity-0 p-2 rounded-sm transition-all duration-150">
                <span ref="messageSpan">error message</span>
            </div>
            <form @submit.prevent="submit()" class="h-2/3 text-foreground w-2/4" v-if="!this.processing">
                <div class="h-2/3 flex flex-col space-y-4">
                    <input ref="fileInput" class="border-2 rounded-md border-foreground transition-all duration-150" type="file" />
                    <div class="flex lg:flex-row flex-col lg:space-y-0 lg:space-x-2 space-y-2">
                        <div class="flex lg:w-1/2 h-1/2 flex-col">
                            <select ref="expireTypeInput"
                                class="text-center lg:h-24 h-12 border-2 text-foreground focus:text-primary border-foreground rounded-lg hover:border-accent bg-primary focus:bg-foreground transition-all duration-150">
                                <option selected disabled hidden value="">Duration</option>
                                <option v-bind:value="expireType.Name" v-for="expireType in this.expireTypes">{{
                                    expireType.Duration === 1 ? `1 minute` : `${expireType.Duration} minutes` }}</option>
                            </select>
                        </div>
                        <div class="lg:w-1/2 h-1/2">
                            <input ref="passwordInput" v-model="password" placeholder="Password (optional)"
                                class="text-center w-full h-12 border-2 text-foreground focus:text-primary border-foreground rounded-lg hover:border-accent bg-primary focus:bg-foreground transition-all duration-150" />
                        </div>
                    </div>
                </div>
                <div class="h-1/3 flex items-center justify-center">
                    <button ref="submitButton"
                        class="m-10 rounded-xl p-3 border-2 border-foreground hover:bg-accent focus:bg-accept transition-all duration-150"
                        type="submit">Share</button>
                </div>
            </form>

            <div class="h-2/3 text-foreground w-full flex justify-center" v-if="this.processing">
                <div class="h-1/2 w-1/2">
                    <Loader/>
                </div>
            </div>
            <Footer/>
        </div>
    </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import Footer from '@/components/Footer.vue';
import Loader from '@/components/Loader.vue';
import { checkAuth, inputError, clearInput } from '@/global/global';
import axios from 'axios';

export default {
    components: {
        Navbar,
        Footer,
        Loader
    },
    data() {
        return {
            expireTypes: [],
            processing: false
        }
    },
    created() {
        checkAuth(this.$router)

        let data = {
            headers: {
                "Authorization": `Bearer ${localStorage.token}`
            }
        }

        axios.get("share/types", data)
            .then((response) => {
                this.expireTypes = response.data
            })
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
        showAcceptedMessage(message, link) {
            console.log(link)
            this.$refs.messageContainer.classList.remove("opacity-0")
            this.$refs.messageContainer.classList.add("bg-accept")

            if (link === null) {
                this.$refs.messageSpan.innerText = message
            }
            else {
                let a = document.createElement('a')
                a.href = link
                a.innerText = link

                this.$refs.messageSpan.innerText = message
                this.$refs.messageSpan.appendChild(a)
            }
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
        clearForm() {
            this.hideMessage()
            clearInput(this.$refs.fileInput)
            clearInput(this.$refs.expireTypeInput)
            clearInput(this.$refs.passwordInput)
        },
        submit() {
            this.clearForm()

            let expireType = this.$refs.expireTypeInput
            let file = this.$refs.fileInput

            let empty = false

            if (file.value === "" || file.value === null) {
                inputError(file)
                empty = true
            }

            if (expireType.value === "Default" || expireType.value === "") {
                inputError(expireType)
                empty = true
            }

            if (empty) {
                this.showErrorMessage("Fill in the fields!")
                return
            }

            if (this.password === undefined || this.password === null) {
                this.password = ""
            }

            this.processing = true

            let formData = new FormData()
            formData.append("file", file.files[0])
            formData.append("expireType", expireType.value)
            formData.append("password", this.password)

            let data = {
                headers: {
                    "Authorization": `Bearer ${localStorage.token}`
                },
                formData
            }

            axios.post("share/links", data.formData, {headers: data.headers})
                .then((response) => {
                    this.processing = false

                    let link = `${location.origin}/get/${response.data.id}`
                    let message = `Your file is successfuly shared! Link to the file: `

                    this.showAcceptedMessage(message, link)
                })
                .catch((error) => {
                    this.processing = false

                    if (error.response) {
                        if (error.response.status != undefined & error.response.status === 401) {
                            localStorage.removeItem("token")
                            checkAuth(this.$router)
                        }

                        let message = error.response.data.message

                        if (message.toLowerCase().includes("password")) {
                            inputError(password)
                        }

                        if (message.toLowerCase().includes("file")) {
                            inputError(file)
                        }

                        if (message.toLowerCase().includes("type")) {
                            inputError(expireType)
                        }

                        this.showErrorMessage(message)
                    }
                    else {
                        this.showErrorMessage("Unknown server error")
                        console.log(error.message)
                    }
                })
            
            //this.processing = false
        }
    }
}
</script>