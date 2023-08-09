
import { defineStore } from 'pinia'

const usePlantStore = defineStore({
    id: 'plant',
    state: () => ({
        plant: {id: 0, plant_name:"leer"},
        loading: false,
        error: null ,
    }),  
    getters: {
      getPlant() {     
        if (this.plant === null){
          this.plant = {id: 0, plant_name:"leer"};
        }   
        return this.plant;
      }
    },
    actions: {
      // Den Betrieb aus dem Store lesen
        loadPlant() {
          const item = localStorage.getItem('localPlant');
          if (item) {
            this.plant = JSON.parse(item);
          }
        },        
    }

   
})

export { usePlantStore }
