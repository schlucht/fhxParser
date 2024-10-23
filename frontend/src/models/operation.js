import { store } from '../store/store.js';

class Operation {
    constructor(plantId) {
        this.plantId = plantId;  
        this.operations = [];        
    }

    async loadOperation() {
        const res = await fetch(api + "/operation/" + this.plantId + "/all");
        const data = await res.json();
        this.operations = data.data;
    }
}

// function allOperations(plantId) {
//     plantId = store.plant.value.id;
//     console.log(plantId);
// }

export { Operation };