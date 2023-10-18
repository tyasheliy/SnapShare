<template>
    <div>
        <Navbar/>
        <div class="frame h-screen bg-secondary flex items-center justify-center">
            <div class="cont lg:w-1/2 w-screen h-2/3 flex flex-col">
                <div class="h-1/4 flex items-center text-foreground justify-center w-full">
                    <span class="text-3xl font-bold">{{ this.userData.userName }}</span>
                </div>
                <div class="h-3/4 w-full text-foreground flex lg:flex-row flex-col">
                    <div class="lg:w-1/2 flex justify-center">
                        <div class="flex flex-col space-y-2 text-center">
                            <span>{{ this.userData.email }}</span>
                            <span>{{ this.userData.isEmailConfirmed ? "Email is confirmed" : "Email is not confirmed"  }}</span>
                            **form**
                        </div>
                    </div>
                    <div class="lg:w-1/2">
                        <form @submit.prevent="">

                        </form>
                    </div>
                </div>
                <Footer/>
            </div>
        </div>
    </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import Footer from '@/components/Footer.vue';
import { checkAuth } from '../global/global.js'
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
        changePassword() {

        },
        changeEmail() {

        },
        logout() {

        }
    }
}
</script>