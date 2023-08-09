<template>
  <FhxPlantOrder :visible="!(plant.plant_name !== 'leer')"></FhxPlantOrder>
  <fhx-header></fhx-header>
  <div class="container">
    <fhx-navbar></fhx-navbar>
    <main>      
      <RouterView />
    </main>
  </div>
</template>
<script setup>
import { RouterView } from 'vue-router'


import FhxHeader from './components/FhxHeader.vue'
import FhxNavbar from './components/FhxNavbar.vue'
import FhxPlantOrder from './components/specials/FhxPlantOrder.vue';
import { usePlantStore } from './stores/plant_store.js';
import { storeToRefs } from 'pinia';

const { plant } = storeToRefs(usePlantStore());
const { loadPlant } = usePlantStore();
// Ladet den Betrieb aus dem Store
loadPlant();

</script>

<style scoped>
 
</style>
