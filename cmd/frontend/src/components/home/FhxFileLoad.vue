<template>
    <div class="new-file">
        <form>
            <fieldset class="control-group">
                <legend><span class="icomoon-folder icon-round"></span>FHX Hochladen</legend>
                <label for="file" class="btn label-file">FHX Datei auswählen</label>
                <input type="file" id="file" name="file">
            </fieldset>
            <fieldset class="radio-button">
                <legend><span class="icomoon-factory icon-round"></span>Betrieb ausswählen:</legend>
                <div v-for="plant in plants" :key="plant.id">
                    <input type="radio" name="plant" :value="plant.id" :id="'plant_'+plant.id">
                    <label :for="'plant_'+ plant.id">{{ plant.plant_name }}</label>
                </div>                
            </fieldset>
            <fieldset class="group-button">
                <legend><span class="icomoon-document"></span>Speichern</legend>
                <button class="btn">Speichern</button>
                <button class="btn" type="button">Abbrechen</button>
            </fieldset>
        </form>
    </div>
</template>
<script setup>
import { ref } from 'vue';

    const plants = ref(null);
    fetch(`${import.meta.env.VITE_API_URL}/all-plants`)
        .then(resp => resp.json())
        .then(data => plants.value = data)
        .catch(err => console.log(err))
</script>

<style scoped>
    .radio-button {
        display: flex;
        flex-direction: row;
        gap: calc(var(--padding)*2)
    }
</style>