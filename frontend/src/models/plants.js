async function allPlants() {
  
    try {
        const resp = await fetch(api + "/plant/all");
        const data = await resp.json();
        
        var plants = [];
        for (let p of data.data.plants) {
            
            plants.push(p)
        }           
        return plants    
    } catch(e) {
        console.error("allPlants", e)
    }    
}

export { allPlants}