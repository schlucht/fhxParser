<script setup>
import FhxFileLoad from '@/components/home/FhxFileLoad.vue';
</script>

<template>  
  <section class="page" id="file-load">    
    <FhxFileLoad />
  </section>  
</template>

<style>

.page {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    min-height: 100vh;
    padding: calc(var(--padding) * 3) 0;
}

</style>