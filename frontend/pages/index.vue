
<template>
  <div>
    <main>
      <section id="unitslist">
        <h2>Alle Standard Units</h2>
      <UnitCard></UnitCard>
        <article>
          <h2>{{ unit }}</h2>
          <ul>
            <li  v-for="u, index in Units" :key="index">
            <ul>
              <li v-for="p, index in u.unitprocedure" :key="index">
                <h3>{{ p.name }}</h3>
                <ul>
                  <li v-for="attr, index in p.params" :key="index">
                    {{  attr.name }}:  
                    {{ attr.value.cv }} -> 
                    {{ attr.value.stringvalue }}
                    <p>{{ attr.desc }}</p>
                  </li>
                </ul>
              </li>
            </ul>
          </li>
          </ul>
          <footer></footer>
        </article>
      
      </section>
      <!--Hello-->
      <section id="recipelist">
        <h2>Produkte mit Rezepte</h2>
      </section>
    </main>
    <footer>
      <p>Copyright &copy;</p>
      <p>Schmid Lothar</p>
      <p>1.2023</p>
    </footer>
  </div>
</template>
<script setup lang="ts">
import { Units } from '../models/loadUnits'
const unit = Units[0].unitname
</script>

<style lang="scss" scoped>

</style>
