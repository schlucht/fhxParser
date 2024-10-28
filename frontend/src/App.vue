<template>  
 	<div class="top w3-container w3-theme-d2">
      <ots-top></ots-top>    
  </div>
  <header class="header w3-container">
      <ots-header></ots-header>
  </header>
  <main id="main">
	  <RouterView />
  </main>
	<footer class="footer w3-container w3-theme-d2">
    <ots-footer></ots-footer>    
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
  .top {
    height: 6rem
  }
  .footer {
    height: 10rem;
  }
</style>
