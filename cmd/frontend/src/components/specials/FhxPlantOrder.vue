<template>
<Transition name="modal-animate">
    <div v-show="show" class="background">
        <div class="messagebox">
            <h2 class="messagebox-title"><span class="icomoon-factory icon-round"></span>Betrieb auswählen</h2>
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
</Transition>
</template>
<script setup>
import { onMounted, ref } from 'vue';
import { loadAllPlants } from '@/models/plants';
import { usePlantStore } from '@/stores/plant_store';
import notie from 'notie'
   
    const props = defineProps({
        show: Boolean
    });
    
    const plants = ref(null);  
    const { closeModal } = usePlantStore();    

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
        let plant = {id, plant_name}
        setTimeout(() => {            
            closeModal(plant);
        }, 500);        
    }  

    onMounted(async() => {
        // alle Betrieb aus der Datenbank laden
        const data = await loadAllPlants()
        
        if (!data.ok) {
            notie.alert({
                type: 'error',
                text: data.message,
            })  
        } else {
            plants.value = JSON.parse(data.content);
        }
    
    });

</script>
<style scoped>
    .modal-animate-enter-active,
    .modal-animate-leave-active {
        transition: opacity 0.3s cubic-bezier(0.52, 0.02, 0.19, 1.02);
    }
    .modal-animate-enter-from {
        opacity: 0;
    }
    .modal-animate-leave-from {
        opacity: 1;
    }
    .modal-animate-leave-to {
        opacity: 0;
    }
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
        display: grid;    
        grid-template-rows: 6rem auto;       
        width: 50rem; 
        height: 20rem; 
        background-color: var(--white);
        text-align: center;
        border-radius: 1rem;
        box-shadow: 4px 2px 5px 20px rgba(255,255,255,.1);
    }
    .messagebox-title {
        font-size: 3rem;
        padding: 1rem 1rem 1rem 1rem;
        background-color: var(--blue);
        color: var(--white);
        & span{
            margin: 2rem 2rem 0 0;
        }  
    }
    .messagebox-form {        
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: space-around;        
        
        & label[for="select-plant"] {
            color: var(--mid-blue);
        }
        & select {
            font-size: 2.5rem;
        }               
    }
  
</style>