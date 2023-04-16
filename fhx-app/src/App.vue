<script setup>
import OtsHeader from './components/ots-header.vue';
import OtsFooter from './components/ots-footer.vue';
</script>

<template>
  <div>
    <ots-header></ots-header>
    <main class="container p-5">
      <router-view></router-view>
    </main>
    <ots-footer></ots-footer>
  </div>
</template>

<style scoped></style>
