<template>
    <div class="new-file">
        <form>
            <fieldset class="control-group">
                <legend><span class="icomoon-folder icon-round"></span>FHX Hochladen</legend>
                <label for="file" class="btn label-file">FHX Datei auswählen</label>
                <input type="file" id="file" name="file">
            </fieldset>
            <fieldset>
                <legend><span class="icomoon-factory icon-round"></span>Betrieb ausswählen</legend>
                <input type="radio" name="plant" value="1" id="plant_d29">
                <label for="plant_d29">D29</label>
                <input type="radio" name="plant" value="0" id="plant_e13">
                <label for="plant_e13">E13</label>
                <input type="radio" name="plant" value="0" id="plant_c07">
                <label for="plant_c07">C07</label>
                <input type="radio" name="plant" value="0" id="plant_b12">
                <label for="plant_b12">B12</label>
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

</script>
<style scoped>
    
</style>