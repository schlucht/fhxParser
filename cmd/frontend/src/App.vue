<script setup>
import { RouterView } from 'vue-router'
import FhxHeader from './components/FhxHeader.vue'
import FhxNavbar from './components/FhxNavbar.vue'
</script>

<template>
  <fhx-header></fhx-header>
  <div class="container">
    <fhx-navbar></fhx-navbar>
    <main>
      <RouterView />
    </main>
  </div>
</template>

<style scoped></style>
