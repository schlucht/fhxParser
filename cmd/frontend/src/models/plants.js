

// holt alle Betriebe vom Server
async function loadAllPlants() {  
    try {
        const resp = await fetch(`${import.meta.env.VITE_API_URL}/all-plants`);
        
        console.log(resp);
        const data = await resp.json();
        var plants = [];
        for (let p of data) {
            plants.push(p)
        }   
        return plants    
    } catch(e) {
        console.log("loadAllPlants:", e)
    }    
}

export { loadAllPlants }