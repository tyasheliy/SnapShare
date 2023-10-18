<template>
    <div>
        <Navbar/>
        <div class="frame bg-secondary h-screen flex flex-col justify-center items-center">
            <div class="cont lg:w-1/2 w-full h-2/3 flex flex-col items-center">
                <div class="h-1/3 flex items-center">
                    <span class="text-foreground text-4xl font-bold">Log in</span>
                </div>
                <div class="h-2/3">
                    <form @submit.prevent="submit()" class="flex flex-col space-y-2">
                        <div ref="messageContainer" class="h-1/3 text-foreground text-center rounded-sm opacity-0 transition-all duration-300">
                            <span ref="messageSpan" class="inline-block w-2/3 p-1">
                                error message
                            </span>
                        </div>
                        <div class="h-1/3 flex flex-col space-y-3">
                            <input ref="usernameInput" placeholder="Username" class="text-center p-2 border-2 text-foreground focus:text-primary border-foreground rounded-lg hover:border-accent bg-primary focus:bg-foreground transition-all duration-150" v-model="username" type="text"/>
                            <input ref="passwordInput" placeholder="Password" class="text-center p-2 border-2 text-foreground focus:text-primary border-foreground rounded-lg hover:border-accent bg-primary focus:bg-foreground transition-all duration-150" v-model="password" type="password"/>
                        </div>
                        <div class="h-1/3 text-foreground text-center">
                            <input class="m-12 rounded-xl p-3 border-2 border-foreground hover:bg-accent focus:bg-accept transition-all duration-150" type="submit"/>
                        </div>
                        
                    </form>
                </div>
                <Footer/>
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
        showErrorMessage(message){
            this.$refs.messageSpan.innerText = message
            this.$refs.messageContainer.classList.add("bg-error")
            this.$refs.messageContainer.classList.remove("opacity-0")
        },
        hideMessageContainer() {
            this.$refs.messageContainer.classList.add("opacity-0")
            this.$refs.messageContainer.classList.forEach(c => {
                if (c.startsWith("bg-")) {
                    this.$refs.messageContainer.classList.remove(c)
                }
            });
        },
        submit() {
            let empty = false

            if (this.username === undefined) {
                this.$refs.usernameInput.classList.add("border-error")
                empty = true
            }

            if (this.password === undefined) {
                this.$refs.passwordInput.classList.add("border-error")
                empty = true
            }

            if (empty) {
                this.showErrorMessage("Fill in the fields!")
                return
            }

            let body = {
                "username": this.username,
                "password": this.password
            }
            axios.post("http://localhost/identity/auth", body)
                .then((response) => {
                    
                })
                .catch((error) => {
                    console.log(error)
                })
        }
    }
}
</script>