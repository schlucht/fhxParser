<template>
  <article class="card-list">
    <h2 class="card-title">
      <span class="icomoon-library icon-round"></span>
      {{ plant.plant_name }} Operationen
      <span class="asterik asterik-orange">{{ ops.Count }}</span>
    </h2>
    <section class="card-body">      
      <ul class="list">
        <fhx-op-card-item 
          v-for="o in ops.Operations" 
          :key="o.name"           
          >
          <router-link 
            :to="{name: 'detail', params: {id: o.id}}"              
            :alt="o.name">{{ o.name }}
          </router-link>
        </fhx-op-card-item>          
      </ul>
    </section>
  </article>
</template>
<script setup>

import {storeToRefs } from 'pinia';
import { ref } from 'vue'
import { laodAlloperations, loadParamsFromOPId } from '../../models/operations.js'
import { usePlantStore } from '@/stores/plant_store';

import FhxOpCardItem from './FhxOpCardItem.vue';

const { plant } = storeToRefs(usePlantStore());
const ops = ref({})
const op = ref({})

async function loadOps() {
  const data = await laodAlloperations(plant.value.id)
  console.log(plant.value.id)
  if (data.content === "{}"){
    if (!data.ok) {
      console.log(data.message)
    }
  } else {
    ops.value = JSON.parse(data.content) 
  }
  
};

async function idOP(e) {
  const opid = e.target.dataset['id'];
  const data = await loadParamsFromOPId(+opid)  
  op.value = JSON.parse(data.content)
  console.log(op.value)
};

loadOps()

</script>
<style scoped>
.cards {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

.card-list {
  display: flex;
  flex-direction: column;
  flex-wrap: wrap;
  justify-content: space-between;
  margin: var(--padding);
  border-radius: calc(var(--border-radius) * 2);
  background-color: var(--white);
  box-shadow: 3px, 3px, 3px, rgba(0, 0, 0, 1);
  min-width: 35rem;
  max-width: 35rem;

  .card-title {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: var(--padding);
    background-color: var(--light-blue);
    border-top-left-radius: calc(var(--border-radius) * 2);
    border-top-right-radius: calc(var(--border-radius) * 2);
    color: var(--white);
    padding: calc(var(--padding));
    .asterik {
      width: 50px;
    }
  }
}
.card-body {
  margin: calc(var(--padding) * 3);
}
.list {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  padding: var(--padding);
  border-bottom: 1px var(--back-crl) solid;
  gap: 5px;
}
</style>
