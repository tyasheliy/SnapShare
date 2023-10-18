export const checkAuth = function(router) {
    if (localStorage.token === undefined) {
        router.push({ name: "login" })
        let promise = new Promise((res, rej) => setTimeout(() => {
            res()
        }, 1))
        promise.then(() => showUnauthLoginMessage())
    }
}

function showUnauthLoginMessage() {
    console.log("Token has expired or unauthenticated")

    let messageContainer = document.querySelector("#messageContainer")
    let messageSpan = document.querySelector("#messageSpan")

    if (messageContainer === undefined || messageSpan === undefined) {
        return
    }

    messageContainer.classList.add("bg-error")
    messageContainer.classList.remove("opacity-0")

    messageSpan.innerText = "Unauth"
}