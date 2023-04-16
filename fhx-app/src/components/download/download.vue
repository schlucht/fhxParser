<template>
    <div class="row">
        <div class="col">
            <form>
                <label class="form-label" for="fhxfile">Datei auswählen</label>
                <input class="form-control" type="file" name="fhxfile" id="fhxfile">
                <button class="btn btn-primary" type="submit">Speichern</button>
                <button class="btn btn-danger" type="cancel">Abbrechen</button>
            </form>
        </div>
        <div class="col-8">
            <preview></preview>
        </div>
    </div>
</template>
<script setup>
import Preview from './preview.vue';


</script>
<style lang='scss' scoped></style>