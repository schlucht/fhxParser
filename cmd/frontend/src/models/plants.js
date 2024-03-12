

// holt alle Betriebe vom Server
async function loadAllPlants() {  
    try {
        const resp = await fetch(`${import.meta.env.VITE_API_URL}/all-plants`);        
        const data = await resp.json();     
        
        return data  
    } catch(e) {
        console.log("loadAllPlants:", e)
    }    
}

export { loadAllPlants }