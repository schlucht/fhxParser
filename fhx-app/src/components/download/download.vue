<template>
    <div class="row">
        <div class="col">
            <form>
                <label class="form-label" for="fhxfile">Datei auswählen</label>
                <input class="form-control" type="file" name="fhxfile" id="fhxfile"  @change="onFileLoad">
                <button class="btn btn-primary" type="submit" @click="onLoadFile">Speichern</button>
                <button class="btn btn-danger" type="cancel">Abbrechen</button>
            </form>
        </div>
        <div class="col-8" >
            <preview>{{ fileInput }}</preview>
        </div>
    </div>
</template>
<script setup>
import Preview from './preview.vue';
import { ref } from 'vue';

const fileInput = ref(null);

function onLoadFile(event) {
    event.preventDefault()
    if (fileInput.value == "") {
        return;
    }
    console.log("save File", fileInput.value.substring(0, 500) + "...")
}

function onFileLoad(event) {
    const files = event.target.files
    const fileReader = new FileReader()
    fileReader.addEventListener('load', () => {
            fileInput.value = fileReader.result
    })
    fileReader.readAsText(files[0])
}

</script>
<style lang='scss' scoped></style>