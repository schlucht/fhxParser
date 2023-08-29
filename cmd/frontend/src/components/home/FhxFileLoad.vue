<template>  
  <div class="new-file">
    <form action="" method="POST" @submit.prevent>
      <fieldset class="control-group">
        <legend><span class="icomoon-folder icon-round"></span>FHX Hochladen</legend>
        <label for="file" class="btn label-file">FHX Datei auswählen</label>
        <input type="file" id="file" name="file" @change="uploadFile" accept=".fhx" />
        <span>{{ fileName }}</span>
      </fieldset>
      <fieldset class="control-group">
        <legend><span class="icomoon-factory icon-round"></span>Betrieb:</legend>
        <!-- <p v-if="loading">Loading posts...</p>
        <p v-if="error"> {{ error.message }}</p> -->       
          <p class="plant-name" @click="selectPlant">{{ plant.plant_name }}</p>        
      </fieldset>
      <fieldset class="group-button">
        <legend><span class="icomoon-database-upload icon-round"></span>Speichern</legend>
        <button ref="save" @click="uploadText" class="btn">Speichern</button>
        <button class="btn" @click="reset" type="button">Abbrechen</button>
      </fieldset>
    </form>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { usePlantStore } from '../../stores/plant_store';
import { storeToRefs } from 'pinia';

import notie from 'notie'
import { fileDataUpload } from '../../models/fileUpload.js';
// import  FhxPlantOrder  from '../../components/home/FhxFileLoad.vue';


const { plant, showModal } = storeToRefs(usePlantStore());
const { openModal } = usePlantStore();
const save = ref(0)
// const showNewPlant = ref(getShowModal);

let plantId = plant.value.id;
let fileName = '';
const fileUpload = { text: '', name: '', plant_id: 0 }

onMounted(() => {
  // Aus und einblenden Speicher Button
  // showModal.value = true
})

function selectPlant() {
  openModal();
  console.log(showModal.value);
}

// Einlesen einer FHX-Datei
function uploadFile(event) {
  if (plantId === 0) {
    notie.alert({
      type: 'error',
      text: 'Keine Anlage ausgelesen!'
    })
    return
  } 

  const files = event.target.files || event.dataTransfer.files
  if (!files.length) return
  fileName = files[0].name
  const fr = new FileReader()
  fr.onload = () => {
    let result = fr.result
    fileUpload.name = fileName
    fileUpload.text = result
    fileUpload.plant_id = parseInt(plantId);
    save.value.disabled = false
  }
  fr.readAsText(files[0])
}

// Hochgeladene FHX Datei auf dem Server speichern
async function uploadText() {
  if (fileUpload.text === '') {
    notie.alert({
      type: 'info',
      text: 'Ausgewählte Datei nicht hochgeladen'
    })
    return
  }
  
  const stream = JSON.stringify(fileUpload) 

  // JSON String auf den Server laden
  const {ok, message} = await fileDataUpload(stream);  
  if (ok) {
      notie.alert({
      type: 'success',
      text: message,
      stay: false
    })
  } else {
      notie.alert({
      type: 'warning',
      text: message,
      stay: false
    })
  }  
}

</script>

<style scoped>
button:disabled {
  background-color: var(--light-gray);
}

.plant-name {
  width: 100%;
  text-align: center;  
  font-size: 2.5rem;
  color: var(--blue);
}
</style>
