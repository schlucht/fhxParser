<template>
    <div v-if="isVisible" class="background">
        <div class="messagebox">
            <h2 class="messagebox-title">Hier ist der Titel</h2>
            <form class="messagebox-form">
                <label for="select-plant">Bitte den Betrieb auswählen</label>
                <select id="select-plant" @change="savePlant">
                    <option value="0">&lt;keine&gt;</option>
                    <option v-for="plant in plants" 
                    :key="plant.id" 
                    :value="plant.id">{{ plant.plant_name }}</option>
                </select>
            </form>
        </div>
    </div>
</template>
<script setup>
import { ref, toRefs } from 'vue';
import { usePlantStore } from '../../stores/plant';
import { storeToRefs } from 'pinia';

const { plants } = storeToRefs(usePlantStore());
    const isVisible = ref('true');
    const props = defineProps({
        visible: String
    })

    const { visible } = toRefs(props);
    isVisible.value = visible.value;

    function savePlant(e) {
        const id = e.target.value;
        const plant_name = e.target.options[e.target.selectedIndex].text;
        if (id == 0) {
            isVisible.value ='true' ;
            return;     
        }
        const item = localStorage.getItem('localPlant');
        if (item) {
            localStorage.removeItem('localPlant');
        }
        localStorage.setItem('localPlant', JSON.stringify({ id, plant_name }));
        isVisible.value = 'false'; 
    }    
    
</script>
<style scoped>
    .background {
        position: absolute;
        height: 100vh;
        width: 100vw;
        background-color: red;
        display: flex;
        justify-content: center;
        align-items: center;
        z-index: 1000;
    }
    .messagebox {
        height: 25rem;
        width: 55rem;
        background-color: var(--white);
    }
</style>