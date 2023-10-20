<template>
  <article class="card-list">
    <h2 class="card-title">
      <span class="icomoon-library icon-round"></span>
      {{ plant.plant_name }} Operationen
    </h2>
    <section class="card-body">
      <ul class="lists">
        <li class="list">
          <span class="asterik asterik-orange">R</span>
          <div>
            <a href="#">Anzahl</a>            
            <p>Anzahl: {{ ops.Count || 0 }}</p>
          </div>
        </li>
        <li class="list">
          <span class="asterik asterik-orange">U</span>
          <div>
            <a href="#">Unitpozeduren</a>
            <p>Letzte Änderung: 25.03.2021</p>
            <p>Anzhal: 21</p>
          </div>
        </li>
        <li class="list">
          <span class="asterik asterik-orange">O</span>
          <div>
            <a href="#">Operationen</a>
            <p>Letzte Änderung: 25.03.2021</p>
            <p>Anzahl: 36</p>
          </div>
        </li>
      </ul>
      <ul>
        <li v-for="o in ops.Operations" :key="o.name">{{ o.name }}</li>
      </ul>
    </section>
  </article>
</template>
<script setup>
import {storeToRefs } from 'pinia';
import { ref } from 'vue'
import { laodAlloperations } from '../../models/operations.js'
import { usePlantStore } from '@/stores/plant_store';

const { plant } = storeToRefs(usePlantStore());
const ops = ref({})

async function loadOps() {
  const data = await laodAlloperations(plant.value.id)
  ops.value = JSON.parse(data.content)  
}

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
  justify-content: space-between;
  margin: var(--padding);
  border-radius: calc(var(--border-radius) * 2);
  background-color: var(--white);
  box-shadow: 3px, 3px, 3px, rgba(0, 0, 0, 1);
  min-width: 35rem;
  max-width: 35rem;
}

.card-list .card-title {
  display: flex;
  align-items: center;
  gap: var(--padding);
  background-color: var(--light-blue);
  border-top-left-radius: calc(var(--border-radius) * 2);
  border-top-right-radius: calc(var(--border-radius) * 2);
  color: var(--white);
  padding: calc(var(--padding));
}
.card-body {
  margin: calc(var(--padding) * 3);
}
.list {
  display: flex;
  align-items: center;
  padding: var(--padding);
  border-bottom: 1px var(--back-crl) solid;
}
.list div {
  margin-left: calc(var(--padding) * 2);
}
.list div p {
  font-size: 1.2rem;
  color: var(--orange);
  margin-top: var(--padding);
}

.card-footer {
  background-color: var(--white);
  padding: var(--padding);
  border-bottom-left-radius: calc(var(--border-radius) * 2);
  border-bottom-right-radius: calc(var(--border-radius) * 2);
  border-top: solid 1px var(--black);
}
.card-footer span {
  font-weight: 700;
}
</style>
