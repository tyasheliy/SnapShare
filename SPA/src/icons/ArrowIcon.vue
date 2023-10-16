<template>
    <svg id="arrow" data-name="1-Arrow Up" xmlns="http://www.w3.org/2000/svg" class="fill-neutral-400 dark:fill-neutral-200" viewBox="0 0 32 32"><path d="m26.71 10.29-10-10a1 1 0 0 0-1.41 0l-10 10 1.41 1.41L15 3.41V32h2V3.41l8.29 8.29z"/></svg>
</template>

<script>
export default {
    name: "ArrowIcon",
    props: ["direction"],
    mounted() {
        if (this.direction === null) {
            return
        }

        let icon = document.getElementById("arrow")

        switch (this.direction){
            case "left":
                icon.classList.add("rotate-90")
                break
            case "right":
                icon.classList.add("-rotate-90")
                break
        }
    }
}
</script>
