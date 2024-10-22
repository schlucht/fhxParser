<template>
    <div>
        <h1>Plants</h1>            
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

    function activatePlant(e) {
        const id = e.target.value;
        const name = e.target.parentElement.parentElement.children[0].innerText;        
        localStorage.setItem("plant", JSON.stringify({id:id, plant:name}));
        store.plant = {id:id, plant:name};
        router.go(0);
    }
    
    onMounted(async() => {
        const res = await fetch("https://5101-schlucht-fhxparser-ai96x3m9vg1.ws-eu116.gitpod.io/plant/all");
        const data = await res.json();  
      
        plants.value = data.data.plants;        
    });
</script>
<style lang='css' scoped>
    
</style>