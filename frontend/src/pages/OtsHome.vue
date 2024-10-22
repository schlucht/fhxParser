<template>
    <div id="top">
        <div class="center-content">
            <ots-top></ots-top>
        </div>
    </div>
    <header id="pageHeader">
        <div class="center-content">
            <ots-header></ots-header>
        </div>
    </header>
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
    <footer id="pageFooter">
        <div class="center-content">
            <ots-footer></ots-footer>
        </div>
    </footer>
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