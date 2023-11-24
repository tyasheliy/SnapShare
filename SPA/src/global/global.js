export const checkAuth = function(router) {
    if (localStorage.token === undefined) {
        router.push({ name: "login" })
        let promise = new Promise((res, rej) => setTimeout(() => {
            res()
        }, 1))
        promise.then(() => showUnauthLoginMessage())
    }
}

export const inputError = function(input) {
    input.classList.forEach((c) => {
        if (c.startsWith("border-") & c != "border-2") {
            input.classList.remove(c)
        }
    })
    input.classList.add("border-error")
}

export const clearInput = function(input) {
    input.classList.forEach((c) => {
        input.classList.forEach((c) => {
            if (c.startsWith("border-") & c != "border-2") {
                input.classList.remove(c)
            }
        })
        input.classList.add("border-foreground")
    })
}

function showUnauthLoginMessage() {
    console.log("Token has expired or unauthenticated")

    let messageContainer = document.querySelector("#messageContainer")
    let messageSpan = document.querySelector("#messageSpan")
    let submitButton = document.querySelector("#submitButton")

    if (messageContainer === undefined || messageSpan === undefined) {
        return
    }

    messageContainer.classList.add("bg-error")
    messageContainer.classList.remove("opacity-0")

    submitButton.classList.add("mt-4")

    messageSpan.innerText = "You are unauthenticated or your session has expired"
}