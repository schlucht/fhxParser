
<template>
  <div>
    <main>
      <section id="unitslist">        
        <UnitCard v-for="u of ups" :key="u.name" :ups="u"></UnitCard>         
      </section>      
      <section id="recipelist">
        <h2>Produkte mit Rezepte</h2>
      </section>
    </main>   
  </div>
</template>
<script setup lang="ts">
import { ups } from '../models/loadUnits'
const unit = ups
</script>

<style lang="scss" scoped>
  #unitslist {
    display: flex;
    gap: 15px;
  }
</style>
