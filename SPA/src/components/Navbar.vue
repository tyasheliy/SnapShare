<template>
<nav class="z-40">
    <button @click="switchMenu()" class="block fixed top-0 right-0 lg:hidden bg-primary z-50 border-foreground border-2 rounded-md mt-7 mr-4">
        <BurgerIcon/>
    </button>
    <div class="fixed top-0 right-0 w-32 scale-0 lg:scale-100 h-screen" @mouseenter="switchMenu()">

    </div>
    <div id="nav-collapse" @mouseleave="switchMenu()" class="fixed top-0 right-0 w-screen lg:w-32 flex flex-col h-screen bg-primary translate-x-full transition-all duration-300">
        <div class="text-center flex items-center justify-center h-24">
            <router-link to="/" class="font-inter text-foreground text-3xl">Logo</router-link>
        </div>
        <div class="h-1/2 flex flex-col justify-center items-center space-y-4">
            <router-link to="/share" class="text-foreground hover:text-xl focus:text-accept hover:text-accent transition-all duration-150">Share</router-link>
            <router-link to="/about" class="text-foreground hover:text-xl focus:text-accept hover:text-accent transition-all duration-150">About</router-link>
        </div>
        <div class="mt-auto h-24 flex items-center justify-center m-16">
            <router-link id="loginLink" to="/login" class="text-foreground hover:text-xl focus:text-accept hover:text-accent transition-all duration-150 whitespace-nowrap hidden">Log in</router-link>
            <router-link id="profileLink" to="/profile" class="text-foreground hover:text-xl focus:text-accept hover:text-accent transition-all duration-150 whitespace-nowrap hidden">Profile</router-link>
        </div>
    </div>
</nav>
</template>

<script>
import BurgerIcon from '@/icons/BurgerIcon.vue';

export default {
    name: "Navbar",
    components: {
        BurgerIcon
    },
    data() {
        return {
            menuOpened: false
        }
    },
    mounted() {
        if (localStorage.token !== undefined) {
            document.querySelector("#profileLink").classList.remove("hidden")
        }
        else {
            document.querySelector("#loginLink").classList.remove("hidden")
        }
    },
    methods: {
        switchMenu() {
            let collapse = document.querySelector("#nav-collapse")
            if (collapse === null) {
                return
            }

            switch (this.menuOpened){
                case false:
                    collapse.classList.remove("translate-x-full")
                    break
                case true:
                    collapse.classList.add("translate-x-full")
                    break
            }

            this.menuOpened = !this.menuOpened
        }
    }
}
</script>