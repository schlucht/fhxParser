<template>
    <div v-if="visible" class="background">
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
                <button class="btn">OK</button>
            </form>
        </div>
    </div>
</template>
<script setup>
import { onMounted, ref, toRefs } from 'vue';
import { loadAllPlants } from '../../models/plants';

    const isVisible = ref(true);
    const props = defineProps({
        visible: Boolean
    });
    
    const plants = ref(null);    

    const { visible } = toRefs(props);
    isVisible.value = visible.value;

    // Speichern des Betreiben sind dem Store
    function savePlant(e) {
        const id = e.target.value;
        const plant_name = e.target.options[e.target.selectedIndex].text;
        if (id == 0) {
            isVisible.value =true ;
            return;     
        }
        // Auswahl des Betriebes in den Store schreiben
        const item = localStorage.getItem('localPlant');
        if (item) {
            localStorage.removeItem('localPlant');
        }
        localStorage.setItem('localPlant', JSON.stringify({ id, plant_name }));        
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
        display: flex;
        flex-direction: column;
        height: 20rem;
        width: 50rem;
        background-color: var(--white);
        text-align: center;
    }
    .messagebox-title {
        font-size: 3rem;
        padding-top: 2rem;
        background-color: var(--blue);
        color: var(--white);
    }
    .messagebox-form {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-around;
        height: 100%;
        width: 100%;        
    }
    label[for="select-plant"] {
        color: var(--mid-blue);
    }
</style>