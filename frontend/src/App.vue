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
	<RouterView />
	<footer id="pageFooter">
        <div class="center-content">
            <ots-footer></ots-footer>
        </div>
    </footer>
</template>

<script setup>
  import OtsHeaderVue from './components/OtsHeader.vue';
	import OtsFooterVue from './components/OtsFooter.vue';
	import OtsTopVue from './components/OtsTop.vue';
  import { store } from './store/store.js';

  
  const actualPlant = localStorage.getItem("plant");
  if (actualPlant) {
    store.plant = JSON.parse(actualPlant);
  }
  
</script>

<style scoped>
  
</style>
