<template>
  <div class="center-content">
    <form action="" method="GET" @submit.prevent>
      <fieldset class="control-group">
        <legend><span class="icomoon-folder icon-round"></span>FHX Hochladen</legend>
        <label for="file" class="btn label-file">FHX Datei auswählen</label>        
        <input type="file" id="file" name="file" @change="uploadFile" accept=".fhx" />
        <span>{{ fileName }}</span>
      </fieldset>
      <fieldset class="radio-button">
        <legend><span class="icomoon-factory icon-round"></span>Betrieb ausswählen:</legend>
        <div class="radio-group" v-for="plant in plants" :key="plant.plant_id">                  
          <input
            v-model="plantId"            
            type="radio"
            name="plant"
            :value="plant.plant"
            :id="plant.plant_id"
            :checked="plant.plant_id === (plantId)"
            @change="setPlant"
          /> 
          <label :for="plant.plant_id">{{ plant.plant }}</label>
        </div>
      </fieldset>
      <fieldset class="group-button">
        <legend><span class="icomoon-document"></span>Speichern</legend>        
        <button ref="save" @click="uploadText" class="btn">Speichern</button>
        <button class="btn" @click="reset" type="button">Abbrechen</button>
      </fieldset>      
    </form>
    <div class="res">
        <pre>{{ smallText }}</pre>
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
import notie from 'notie';

const plants = ref(null)
const plantId = ref(0)
const fileName = ref('')
const save = ref(0)
const smallText = ref('');
const fileUpload = { text: '', name: '', plantId: 0 }
const router = useRouter();

onMounted(async () => {
  save.value.disabled = true; 
  plants.value = await allPlants(); 
  plantId.value =  store.plant.id;
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
  fileName.value = '';
  fileUpload.text = '';
  fileUpload.name = '';
  fileUpload.plantId = 0;
  save.value.disabled = true;
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

async function uploadText() {
    smallText.value =  readTitleFromFhx(fileUpload.text);
    const body = {
      text: fileUpload.text,
      name: fileUpload.name,
      plant_id: fileUpload.plantId,
    };
   
    const res = await fetch(api + "/fhx/upload", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(body),
        });
        
    const data = await res.json();
    console.log(data);
    if (data.error) {
      notie.alert({
        type:'error',
        text: data.message,
      });
      return;
    } else {
      notie.alert({
        type:'success',
        text: data.message,
      });
    }
}

</script>
<style scoped>
.radio-group {
  display: flex;
  flex-direction: row;  
  margin: 3rem;
}
input[type=radio]{  
  appearance: none;
  display: flex;
  flex-direction: row;  
  & ~ label {
    display: flex;
    flex-direction: row;
    font-size: 2rem;
  }
    &:checked ~ label::before {
      content: '✔';
      font-size: 2rem;
      margin-right: 1rem;    
    }
}


button:disabled {
  background-color: var(--light-gray);
}
</style>