<script setup>
import FhxFileLoad from '../components/home/FhxFileLoad.vue'
import FhxOpCard from '../components/card/FhxOpCard.vue'
import { RouterView } from 'vue-router';

</script>

<template>  
  <section class="" id="file-load">
    <fhx-file-load></fhx-file-load>
  </section>
  <section class="page">
    <section class="left-list m-5-t" id="card-section">
      <fhx-op-card></fhx-op-card>
    </section>
    <section class="right-details m-5-t">
      <RouterView />
    </section>
  </section>  
</template>

<style>
.page {
  display: grid;
  grid-template-columns: 40rem 1fr;
}

</style>