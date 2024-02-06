<script setup>
import FhxHome from '@/components/home/FhxHome.vue'
import FhxFileLoad from '@/components/home/FhxFileLoad.vue';
</script>

<template>  
  <section class="" id="file-load">
    <fhx-home></fhx-home>
  </section>  
  <FhxFileLoad />
</template>

<style>
.page {
  display: grid;
  grid-template-columns: 40rem 1fr;
}

</style>