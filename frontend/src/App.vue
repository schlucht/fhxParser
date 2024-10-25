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
  <main class="center-content">
	  <RouterView />
  </main>
	<footer id="pageFooter">
    <div class="center-content">
        <ots-footer></ots-footer>
    </div>
  </footer>
</template>

<script setup>
  import OtsHeader from './components/OtsHeader.vue';
	import OtsFooter from './components/OtsFooter.vue';
	import OtsTop from './components/OtsTop.vue';
  import { store } from './store/store.js';

  
  const actualPlant = localStorage.getItem("plant");
  if (actualPlant) {
    store.plant = JSON.parse(actualPlant);
  }
 console.log(store.plant) 
</script>

<style scoped>
  
</style>
