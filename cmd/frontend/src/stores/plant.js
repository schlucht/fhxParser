import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { loadAllPlants } from '../api/plants'

const usePlantStore = defineStore({
    id: 'plants',
    state: () => ({
        plants: [],
        loading: false,
        error: null
    }),
    actions: {
        async loadPlants() {
            this.plants = [];
            this.loading = true;
            try {                
                this.plants = await loadAllPlants()
                console.log(this.plants)
            } catch (e) {
                this.error = e;
            } finally {
                this.loading = false
            }
        }
    }
})

export { usePlantStore }
