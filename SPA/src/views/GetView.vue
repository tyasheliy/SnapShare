<template>
    <div class="frame h-screen bg-primary flex flex-col items-center justify-center text-foreground">
        <div class="flex flex-col h-full">
            <div class="h-1/3 flex items-center justify-center">
                <span class="text-6xl font-bold">SnapShare</span>
            </div>
            <div ref="messageContainer" class="text-center rounded-sm m-4 p-2 transition-all duration-150 opacity-0">
                <span ref="messageSpan">error message</span>
            </div>
            <div v-if="this.secured" class="h-2/3 flex flex-col space-y-4">
                <input ref="passwordInput" v-model="password" placeholder="Password" type="password"
                    class="text-center p-2 border-2 text-foreground focus:text-primary border-foreground rounded-lg hover:border-accent bg-primary focus:bg-foreground transition-all duration-150" />
                <button @click="submit()"
                    class="rounded-xl p-3 border-2 border-foreground hover:bg-accent focus:bg-accept transition-all duration-150">Send</button>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import { inputError, clearInput } from '@/global/global';

export default {
    data() {
        return {
            secured: false
        }
    },
    created() {
        axios.head(`share/links/${this.$route.params.id}`)
            .then((response) => window.location.href = `${axios.defaults.baseURL}share/links/${this.$route.params.id}`)
            .catch((error) => {
                if (error.response) {
                    if (error.response.status == 400) {
                        axios.get(`share/links/${this.$route.params.id}`)
                            .catch((error) => {
                                if (!error.response) {
                                    this.showErrorMessage("Unknown server error")
                                    return
                                }

                                let message = error.response.data.message

                                if (message.toLowerCase().includes("password")) {
                                    this.secured = true
                                }

                                this.showErrorMessage(message)
                            })
                    }
                }
                else {
                    this.showErrorMessage("Unknown server error")
                }
            })
    },
    methods: {
        showErrorMessage(message) {
            this.$refs.messageContainer.classList.remove("opacity-0")
            this.$refs.messageContainer.classList.add("bg-error")

            this.$refs.messageSpan.innerText = message
        },
        clear() {
            let container = this.$refs.messageContainer

            container.classList.add("opacity-0")
            container.classList.forEach((c) => {
                if (c.startsWith("bg-")) {
                    container.classList.remove(c)
                }
            })

            clearInput(this.$refs.passwordInput)
        },
        submit() {
            this.clear()

            let data = {
                "password": this.password
            }

            axios.post("share/entries", data)
                .then((response) => {
                    let entryId = response.data.id
                    axios.head(`share/links/${this.$route.params.id}?entry=${entryId}`)
                        .then((response) => {
                            window.location.href = `${axios.defaults.baseURL}share/links/${this.$route.params.id}?entry=${entryId}`
                        })
                        .catch((error) => {
                            if (!error.response) {
                                this.showErrorMessage("Unknown server error")
                                return
                            }
                            if (error.response.status == 400) {
                                axios.get(`share/links/${this.$route.params.id}?entry=${entryId}`)
                                    .catch((error) => {
                                        if (!error.response) {
                                            this.showErrorMessage("Uknown server error")
                                            return
                                        }

                                        let message = error.response.data.message

                                        if (message.toLowerCase().includes("password")) {
                                            inputError(this.$refs.passwordInput)
                                        }

                                        this.showErrorMessage(message)
                                        return;
                                    })
                            }
                        })
                })
                .catch((error) => {
                    if (error.response.data.message === undefined || error.response.data.message === null) {
                        this.showErrorMessage("Unknown server error")
                        return
                    }

                    this.showErrorMessage(error.response.data.message)
                })
        }
    }
}
</script>