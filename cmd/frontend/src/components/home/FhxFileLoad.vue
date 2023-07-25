<template>
  <div class="new-file">
    <form action="" method="GET" @submit.prevent>
      <fieldset class="control-group">
        <legend><span class="icomoon-folder icon-round"></span>FHX Hochladen</legend>
        <label for="file" class="btn label-file">FHX Datei auswählen</label>
        <input type="file" id="file" name="file" @change="uploadFile" accept=".fhx" />
        <span>{{ fileName }}</span>
      </fieldset>
      <fieldset class="radio-button">
        <legend><span class="icomoon-factory icon-round"></span>Betrieb ausswählen:</legend>
        <div v-for="plant in plants" :key="plant.id">
          <input
            v-model="plantId"
            type="radio"
            name="plant"
            :value="plant.id"
            :id="'plant_' + plant.id"
          />
          <label :for="'plant_' + plant.id">{{ plant.plant_name }}</label>
        </div>
      </fieldset>
      <fieldset class="group-button">
        <legend><span class="icomoon-document"></span>Speichern</legend>
        <button ref="save" @click="uploadText" class="btn">Speichern</button>
        <button class="btn" @click="reset" type="button">Abbrechen</button>
      </fieldset>
    </form>
  </div>
</template>
<script setup>
import { onMounted, ref } from 'vue'
import notie from 'notie'

const plants = ref(null)
const plantId = ref(0)
const fileName = ref('')
const save = ref(0)
const fileUpload = { text: '', name: '', plantId: 0 }

onMounted(() => {
  save.value.disabled = true
})

fetch(`${import.meta.env.VITE_API_URL}/all-plants`)
  .then((resp) => resp.json())
  .then((data) => (plants.value = data))
  .catch((err) => console.log(err))

function reset(event) {
  console.log(plantId.value)
}

function uploadFile(event) {
  if (plantId.value === 0) {
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
    fileUpload.plantId = plantId.value
    save.value.disabled = false
  }
  fr.readAsText(files[0])
}
function uploadText() {
  if (fileUpload.text === '') {
    notie.alert({
      type: 'info',
      text: 'Kein Text hochgeladen'
    })
    return
  }
  notie.alert({
    type: 'success',
    text: `Datei ${fileUpload.name} hochgeladen!`,
    stay: true
  })
}
</script>

<style scoped>
.radio-button {
  display: flex;
  flex-direction: row;
  gap: calc(var(--padding) * 2);
}
button:disabled {
  background-color: var(--light-gray);
}
</style>
