<template>
    <div class="frame bg-secondary h-screen flex flex-col items-center justify-center">
        <Navbar/>
        <div class="cont lg:w-1/2 w-full h-5/6 text-foreground flex flex-col items-center">
            <div class="h-1/4 flex items-center justify-center">
                <span class="text-3xl">Logo</span>
            </div>
            <div class="h-3/4 overflow-y-auto overflow-x-hidden w-3/4 space-y-3">
                <p>
                    Lorem ipsum dolor, sit amet consectetur adipisicing elit. Reprehenderit ea, expedita numquam optio dolores dicta dolore molestias sapiente hic placeat maiores. Vero sit, quo culpa molestias repellendus rerum nemo ut.
                </p>
                <p>
                    Lorem ipsum dolor sit amet consectetur adipisicing elit. Aperiam dolorem id necessitatibus fugiat, est rerum veritatis asperiores ratione, repudiandae dolor maxime expedita ducimus corrupti nam enim inventore ex laborum veniam?
                </p>
                <p>
                    Lorem ipsum dolor sit amet consectetur adipisicing elit. Nisi, aliquam. Perspiciatis nisi modi ipsa asperiores quibusdam cumque amet, reiciendis aspernatur adipisci, harum rerum voluptatibus sequi corrupti, perferendis hic mollitia dignissimos?
                </p>
            </div>
            <Footer/>
        </div>
    </div>
</template>

<script>
import Navbar from '@/components/Navbar.vue';
import Footer from '@/components/Footer.vue';

export default {
    components: {
        Navbar,
        Footer
    }
}
</script>