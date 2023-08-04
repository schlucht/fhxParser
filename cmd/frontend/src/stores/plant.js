
import { defineStore } from 'pinia'
import { loadAllPlants } from '../api/plants'
import { computed } from 'vue';

const usePlantStore = defineStore({
    id: 'plants',
    state: () => ({
        plants: [],
        loading: false,
        error: null ,
        actualPlant: null,
        
    }),
    getters: {
        getActualPlant: (state) => {
            return computed(() => state.actualPlant);            
        }
    },
    actions: {
        async loadPlants() {
            this.plants = [];
            this.loading = true;
            try {                
                this.plants = await loadAllPlants()                      
            } catch (e) {
                this.error = e;
            } finally {
                this.loading = false
            }
        },
        setActualPlant(plant) {
           this.actualPlant = plant;
        }
    }

   
})

export { usePlantStore }
