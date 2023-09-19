<template>
  <FhxPlantOrder :show= showModal></FhxPlantOrder>
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

import FhxHeader from '@/components/FhxHeader.vue'
import FhxNavbar from '@/components/FhxNavbar.vue'
import FhxPlantOrder from '@/components/specials/FhxPlantOrder.vue';

import { usePlantStore } from '@/stores/plant_store.js';
import { storeToRefs } from 'pinia';

const { showModal } = storeToRefs(usePlantStore());
const { loadPlant } = usePlantStore();

// Ladet den Betrieb aus dem Store
loadPlant();
console.log('App: ', showModal.value)
</script>

<style scoped>
 
</style>
