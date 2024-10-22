async function allPlants() {
  
    try {
        const resp = await fetch("https://5101-schlucht-fhxparser-ai96x3m9vg1.ws-eu116.gitpod.io/plant/all");
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