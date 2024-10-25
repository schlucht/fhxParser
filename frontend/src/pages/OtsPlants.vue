<template>
    <div>
        <form>
            <div class="control-group">
                <label for="plantname" class="control-label">Anlage Name</label>
                <input id="plantname" name="plantname" placeholder="Anlage Name" v-model="plantName">
            </div>
            <a href="javascript:void(0)" class="btn btn-primary" @click="savePlant">Speichern</a>
        </form>
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
    import notie from 'notie';    

    const plants = ref([]);
    const router = useRouter();
    const plantName = ref("");    

    function activatePlant(e) {
        const id = e.target.value;
        const name = e.target.parentElement.parentElement.children[0].innerText;        
        localStorage.setItem("plant", JSON.stringify({id:id, plant:name}));
        store.plant = {id:id, plant:name};
        router.go(0);
    }

    async function savePlant() {
        console.log(plantName.value);
        const body = {
            plant: plantName.value
        };      

       // Anlage in der DB speichern
       try {
       const res =  await fetch(api + "/plant/save", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(body),
        });
       const data = await res.json();
       if (data.error) {
           notie.alert({
               type: 'error',
               text: data.message
           });
       } else {
           plantName.value = "";
       }
       } catch (err) {
           console.log("Fehler: " + err);
       }
    }
    
    onMounted(async() => {        
        // auslesen der Daten aus der DB
        const res = await fetch(api + "/plant/all");
        const data = await res.json();  
      
        plants.value = data.data.plants;        
    });


</script>
<style lang='css' scoped>
    
</style>