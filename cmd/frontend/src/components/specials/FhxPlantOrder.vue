<template>
    <div v-if="usePlantStore.showModal" class="background">
        <div class="messagebox">
            <h2 class="messagebox-title">Betrieb auswählen</h2>
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
import { onMounted, ref } from 'vue';
import { loadAllPlants } from '../../models/plants';
import { usePlantStore } from '../../stores/plant_store';
import {storeToRefs } from 'pinia';
   
    const props = defineProps({
        visible: Boolean
    });
    
    const plants = ref(null);  
   

    // const { visible } = toRefs(props);    

    // Speichern des Betreiben sind dem Store
    function savePlant(e) {
        const id = parseInt(e.target.value);
        const plant_name = e.target.options[e.target.selectedIndex].text;
        if (id == 0) {
            usePlantStore.openModal();
            return;     
        }
        // Auswahl des Betriebes in den Store schreiben
        const item = localStorage.getItem('localPlant');
        if (item) {
            localStorage.removeItem('localPlant');
        }
        localStorage.setItem('localPlant', JSON.stringify({ id, plant_name })); 
        
        setTimeout(() => {
            usePlantStore.closeModal();
        }, 2000);
        
    }  

    onMounted(async() => {
        // alle Betrieb aus der Datenbank laden
        plants.value = await loadAllPlants();
    });

</script>
<style scoped>
    .background {
        position: absolute;
        height: 100vh;
        width: 100vw;
        background-color: rgba(155,155,155,0.8);
        display: flex;
        justify-content: center;
        align-items: center;
        z-index: 1000;
    }
    .messagebox {            
        width: 50rem;
        background-color: var(--white);
        text-align: center;
    }
    .messagebox-title {
        font-size: 3rem;
        padding: 1rem;
        background-color: var(--blue);
        color: var(--white);
    }
    .messagebox-form {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;               
    }
    label[for="select-plant"] {
        color: var(--mid-blue);
    }
    select {
        font-size: 2.5rem;
    }
</style>