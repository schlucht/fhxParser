<template>
    <div class="container">
 <form action="" method="GET" @submit.prevent>
      <fieldset class="control-group">
        <legend><span class="icomoon-folder icon-round"></span>FHX Hochladen</legend>
        <label for="file" class="btn label-file">FHX Datei auswählen</label>        
        <input type="file" id="file" name="file" @change="uploadFile" accept=".fhx" />
        <span>{{ fileName }}</span>
      </fieldset>
      <fieldset class="radio-button">
        <legend><span class="icomoon-factory icon-round"></span>Betrieb ausswählen:</legend>
        <div v-for="plant in plants" :key="plant.plant_id">                  
          <input
            v-model="plantId"
            type="radio"
            name="plant"
            :value="plant.plant"
            :id="plant.plant_id"
            :checked="plant.plant_id === store.plant.id"
            @change="setPlant"
          />
          <label :for="'plant_' + plant.plant_id">{{ plant.plant }}</label>
        </div>
      </fieldset>
      <fieldset class="group-button">
        <legend><span class="icomoon-document"></span>Speichern</legend>        
        <button ref="save" @click="uploadText" class="btn">Speichern</button>
        <button class="btn" @click="reset" type="button">Abbrechen</button>
      </fieldset>      
    </form>
    <div class="res">
        <pre>{{ fileUpload.text }}</pre>
    </div>
  </div>
</template>
<script setup>
import { onMounted, ref } from 'vue'
import { store } from '../store/store.js';
import { readTitleFromFhx } from '../assets/js/helpers.js';
import { allPlants } from '../models/plants.js';
import { Storage } from '../models/storage.js';
import { useRouter } from 'vue-router';

const plants = ref(null)
const plantId = ref(0)
const fileName = ref('')
const save = ref(0)
const fileUpload = { text: '', name: '', plantId: 0 }
const router = useRouter();

onMounted(async () => {
    save.value.disabled = true    
      
    plants.value = await allPlants();   
});

function setPlant(event) {
    const ps = {
        id: event.target.id,
        plant: event.target.value,
    };
    const str = new Storage('plant', ps);
    store.plant = ps;
    str.save();
    router.push('/fhx');
}

function reset(event) {
  console.log(plantId.value)
}

function uploadFile(event) {
 
  const files = event.target.files || event.dataTransfer.files
  if (!files.length) return
  fileName.value = files[0].name
  const fr = new FileReader()
  fr.onload = () => {
    let result = fr.result
    fileUpload.name = fileName.value
    fileUpload.text = result
    fileUpload.plantId = plantId.value
    save.value.disabled = false
  }
  fr.readAsText(files[0])
}

function uploadText() {
    readTitleFromFhx(fileUpload.text);
    // console.log(fileUpload.text);
    // übergabe des Textes an die API
}

</script>
<style scoped>
.container {
    display: flex;
    flex-direction: row;
}
.radio-button {
  display: flex;
  flex-direction: row;
  gap: calc(var(--padding) * 2);
}
button:disabled {
  background-color: var(--light-gray);
}
</style>