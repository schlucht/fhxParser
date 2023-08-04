<template>
  <FhxPlantOrder :visible="hasPlant"></FhxPlantOrder>
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
import { usePlantStore } from './stores/plant';
import { storeToRefs } from 'pinia';
import { ref } from 'vue';

import FhxHeader from './components/FhxHeader.vue'
import FhxNavbar from './components/FhxNavbar.vue'
import FhxPlantOrder from './components/specials/FhxPlantOrder.vue';

const hasPlant = ref('');
const plantname = ref('');
const { setActualPlant, loadPlants } = usePlantStore();


l_Plants()

async function l_Plants() {
  await loadPlants()
  const plant = JSON.parse(localStorage.getItem('localPlant'));
  setActualPlant(plant);
  if (plant){
    plantname.value = plant.plant_name;  
    hasPlant.value = "false";
  } else {
    hasPlant.value = "true";
  }
  console.log(hasPlant.value)
  
}
</script>

<style scoped></style>
