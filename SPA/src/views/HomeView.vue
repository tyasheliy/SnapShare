<template>
  <Navbar />
  <div class="h-screen bg-zinc-50 dark:bg-zinc-800">
      <button @click="switchMode()">Switch display mode</button>
  </div>
  <Footer />
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import Footer from '@/components/Footer.vue';

export default {
  name: "HomeView",
  components: {
    Navbar,
    Footer
  },
  methods: {
    switchMode(){
      if (localStorage.theme === ""){
        localStorage.theme = "light"
      }

      switch (localStorage.theme) {
        case "light":
          localStorage.theme = "dark"
          document.documentElement.classList.add("dark")
          break
        case "dark":
          localStorage.theme = "light"
          document.documentElement.classList.remove("dark")
          break
      }
    }
  }
}
</script>