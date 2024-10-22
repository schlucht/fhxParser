async function allPlants() {
  
    try {
        const resp = await fetch("hhtps//5101-idx-fhxparsergit-1729582237031.cluster-23wp6v3w4jhzmwncf7crloq3kw.cloudworkstations.dev/plant/all");
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