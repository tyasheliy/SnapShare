<template>
  <div class="flex flex-col">
    <Navbar ref="nav" />
    <div class="frame h-screen bg-secondary flex flex-col items-center justify-center">
      <div class="cont lg:w-1/2 w-full h-2/3 flex flex-col items-center">
        <div class="flex items-center justify-center h-1/3">
          <span class="text-foreground text-3xl">Logo</span>
        </div>
        <div class="text-center text-foreground flex flex-col justify-center space-y-2 w-2/3 h-1/3">
          <span class="text-2xl font-semibold">Welcome to SnapShare!</span>
          <p>
            SnapShare makes your files accessable to anyone in the internet in a few clicks!
          </p>
        </div>
        <div class="text-center text-foreground h-1/3 flex items-end">
          <button id="getStartedButton" @click="getStarted()" class="m-14 rounded-xl p-3 border-2 border-foreground hover:bg-accent focus:bg-accept transition-all duration-150">Get started</button>
        </div>
        <Footer/>
      </div>
    </div>
  </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import Footer from '@/components/Footer.vue';

export default {
  name: "HomeView",
  components: {
    Navbar,
    Footer,
  },
  methods: {
    getStarted(){
      this.$refs.nav.switchMenu()
    }
  }
}
</script>

<style scoped>
#frame {
  height: 100vh;
}
</style>