<script setup>
import OtsHeader from './components/ots-header.vue';
import OtsFooter from './components/ots-footer.vue';
</script>

<template>
  <div>
    <ots-header></ots-header>
    <router-view></router-view>
    <ots-footer></ots-footer>
  </div>
</template>

<style scoped></style>
