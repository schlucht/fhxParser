<template>  
  <RouterView/>
</template>

<script setup>
  import { store } from './store/store.js';
  
  const actualPlant = localStorage.getItem("plant");
  if (actualPlant) {
    store.plant = JSON.parse(actualPlant);
  }
  
</script>

<style scoped>
  
</style>
