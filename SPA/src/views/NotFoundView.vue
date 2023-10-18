<template>
    <div class="frame h-screen bg-secondary flex items-center justify-center">
        <Navbar/>
        <div class="cont lg:w-1/2 w-full h-2/3 flex flex-col">
            <div class="h-1/3 flex items-center justify-center">
                <span class="text-foreground text-4xl">Logo</span>
            </div>
            <div class="h-1/3 flex flex-col items-center justify-center text-foreground">
                <span class="text-5xl font-bold">Page is not found</span>
            </div>
            <div class="h-1/3 flex items-center justify-center text-foreground">
                <router-link to="/" class="border-2 border-foreground p-3 rounded-xl hover:bg-accent focus:bg-accept transition-all duration-150">Return to home page</router-link>
            </div>
            <Footer/>
        </div>
    </div>
</template>

<script>
import Footer from '@/components/Footer.vue';
import Navbar from '@/components/Navbar.vue';

export default {
    components: {
        Footer,
        Navbar
    }
}
</script>