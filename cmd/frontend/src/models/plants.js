

// holt alle Betriebe vom Server
async function loadAllPlants() {  
    try {        
        const resp = await fetch(`${import.meta.env.VITE_API_URL}/all-plants`);        
        const data = await resp.json();     
        
        return data  
    } catch(e) {
        console.error("Fehler in loadAllPlants:", e)
        return {
            ok: false,            
            message: "Fehler in loadAllPlants: " + e,
            content: {},
            id: 1,
        }
    }    
}

export { loadAllPlants }