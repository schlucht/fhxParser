<template>
   
    <main id="pageContent">
        <div class="center-content">
            <div class="home">
                <h2>Delta-V Documentation</h2>
                <hr>
                <p>
                    Parameter und Dokumentationen von Delta-V Rezepten.
                </p>
                <div class="content-image">
                    <img src="../assets/images/deltav_main.jpg" alt="Delta-V">
                </div>
            </div>
        </div>
    </main>
    
</template>
<script setup>
import OtsFooter from '../components/OtsFooter.vue';
import OtsHeader from '../components/OtsHeader.vue';
import OtsTop from '../components/OtsTop.vue';
    
</script>
<style lang='css' scoped>
.home {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: space-around;
}

</style>