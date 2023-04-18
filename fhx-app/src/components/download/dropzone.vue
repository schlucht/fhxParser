<template>
    <div class="dropzone"
        @dragenter.prevent="toggleActive"
        @dragleave.prevent="toggleActive"
        @dragover.prevent
        @drop.prevent="toggleActive"
        :class="{ 'active-dropzone': active }">
        <span>Drog or Drop File</span>
        <span>OR</span>
        <label for="dropzoneFile">Select File</label>
        <input type="file" name="dorpzoneFile" id="">
    </div>
</template>
<script setup>
import { ref } from 'vue';
const active = ref(false)
const toggleActive = () => {
    active.value = !active.value;
};

</script>
<style lang="scss"  scoped>
@import '../../assets/styles/style.scss';

// https://github.com/johnkomarnicki/vue-3-dropzone
.dropzone {
    width: 400px;
    height: 200px;
    display: flex;
    @include flex-col-center;
    row-gap: 16px;
    border: dashed $green 1px;

    label {
        padding: 8px 12px;
        color: $white;
        background-color: $green;
        transition: 0.3s ease all;
    }

    input {
        display: none;
    }
}

.active-dropzone {
    color: $blue;
    border-color: $orange;

    label {
        background-color: $light-gray;
        color: $green;
    }
}
</style>