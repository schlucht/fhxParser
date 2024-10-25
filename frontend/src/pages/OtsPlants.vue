<template>
    <div>
        <h1>Plants</h1>            
            <form method="@submit.prevert">
                <input type="text" v-model="plantname" id="plantname">
                <label for="plantname">Anlage Name</label>
                <button type="submit" @click=createPlant>Create</button>
            </form>
            <ul  v-for="plant in plants" :key="plant.plant_id">                
                <li>
                    <input 
                        type="checkbox" 
                        @change=activatePlant 
                        :value="plant.plant_id"
                        :checked="plant.plant_id === store.plant.id"
                        >
                    {{ plant.plant }}
                </li>
            </ul>   
    </div>
</template>
<script setup>
    import { onMounted, ref } from 'vue';
    import { store } from '../store/store.js';
    import { useRouter } from 'vue-router';

    const plants = ref([]);
    const router = useRouter();
    const plantname = ref('');

    function activatePlant(e) {
        const id = e.target.value;
        const name = e.target.parentElement.parentElement.children[0].innerText;        
        localStorage.setItem("plant", JSON.stringify({id:id, plant:name}));
        store.plant = {id:id, plant:name};
        router.go(0);
    }

    async function createPlant(ev) {
        ev.preventDefault();
        if (plantname.value === '') {
            return;
        }
        const res = await fetch(api + "/plant/save", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                plant: plantname.value
            })
        });
        const data = await res.json();
        
    }
    
    onMounted(async() => {
        console.log(api)
        const res = await fetch(api + "/plant/all");
        const data = await res.json();  
      
        plants.value = data.data.plants;        
    });
</script>
<style lang='css' scoped>
    
</style>