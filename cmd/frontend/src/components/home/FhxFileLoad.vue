<template>  
  <div class="new-file">
    <form action="" method="POST" @submit.prevent>
      <fieldset class="control-group">
        <legend><span class="icomoon-folder icon-round"></span>FHX Hochladen</legend>
        <label for="file" class="btn label-file">FHX Datei auswählen</label>
        <div class="file-name">
          <input type="file" id="file" name="file" @change="uploadFile" accept=".fhx" />
          <span>{{ fileName }}</span>
        </div>
      </fieldset>
      <fieldset class="control-group">
        <legend><span class="icomoon-factory icon-round"></span>Betrieb:</legend>
        <!-- <p v-if="loading">Loading posts...</p>
        <p v-if="error"> {{ error.message }}</p> -->       
        <p class="plant-name" @click="selectPlant">{{ plant.plant_name }}</p>        
      </fieldset>
      <fieldset class="group-button">
        <legend><span class="icomoon-database-upload icon-round"></span>Speichern</legend>
        <button  @click="uploadText" class="btn" :disabled="isSaved">Speichern</button>
        <button class="btn" @click="resets" type="button" :disabled="isSaved">Abbrechen</button>
      </fieldset>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import notie from 'notie'
import { storeToRefs } from 'pinia';

import { usePlantStore } from '@/stores/plant_store';
import { fileDataUpload } from '@/models/fileUpload.js';

const { plant } = storeToRefs(usePlantStore());
const { openModal } = usePlantStore();
const isSaved = ref(true)

let plantId = plant.value.id;
const fileName = ref('');
const fileUpload = { text: '', name: '', plant_id: 0 }

function selectPlant() {
  usePlantStore.showModal = openModal();  
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
  fileName.value = files[0].name
  const fr = new FileReader()
  fr.onload = () => {
    let result = fr.result
    fileUpload.name = fileName.value
    fileUpload.text = result
    fileUpload.plant_id = parseInt(plantId);
    isSaved.value = false
  }
  fr.readAsText(files[0])
}

// Zurücksetzen und löschen der geladenen Datei
function resets() {
  fileName.value = ''
  fileDataUpload.text =  fileName
  fileDataUpload.plant_id = 0
  isSaved.value = true
}

// Hochgeladene FHX Datei auf dem Server speichern
async function uploadText() {
  if (fileUpload.text === '') {
    notie.alert({
      type: 'info',
      text: 'Ausgewählte Datei nicht hochgeladen'
    })
    fileName.value = ''
    fileDataUpload.text =  fileName
    fileDataUpload.plant_id = 0
    return
  }
  isSaved.value = true  
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
  fileName.value = ''  
}

</script>

<style scoped>

form {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: calc(var(--padding)*4);
  padding: 1rem 0rem 2rem 0rem;
  border-bottom: 5px solid var(--light-gray);
}

button:disabled {
  background-color: var(--light-gray);
}
.control-group {
  border: none;
  display: flex;
  flex-direction: row;
}

.plant-name {
  width: 20rem;
  text-align: center;  
  font-size: 2.5rem;
  color: var(--blue);
  cursor: pointer;
}

.file-name {
  margin-top: 1rem;
}
</style>
