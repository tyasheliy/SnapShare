<template>
  <router-view/>
</template>

<script>
function checkAuth(){
  if (localStorage.token === undefined) {
            let promise = new Promise((res, rej) => setTimeout(() => {
                res()
            }, 200))
            promise.then(() => redirectWithUnauthError())
  }
}

function redirectWithUnauthError() {
  console.log("Token has expired or unauthenticated")

  let messageContainer = document.querySelector("#messageContainer")
  let messageSpan = document.querySelector("#messageSpan")

  if (messageContainer === undefined || messageSpan === undefined) {
    return
  }
            
  messageContainer.classList.add("bg-error")
  messageContainer.classList.remove("opacity-0")

  messageSpan.innerText = ""

  this.$router.push({name: "login"})
}
</script>

<style>
  @import url('https://fonts.googleapis.com/css2?family=Inter&display=swap');

  .frame {
    width: 100vw;
    overflow-x: hidden;
  }

  .cont {
    background: #1E1E24;
  }

  body {
    -ms-overflow-style: none;
    scrollbar-width: none;
  }

  body::-webkit-scrollbar {
    display: none;
  }
</style>
