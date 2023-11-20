
/**
 * Alle Operation anhand der Anlagen ID einlesen
 * @param {int} plantId 
 */
async function laodAlloperations(plantId) {       

    let antwort = JSON.stringify({id: plantId})
    
    const requestOptions = {
        method: 'POST',
        headers: {
        'Accept': 'application/text',
        'Content-Type': 'application/text'
        },
        body: antwort
    }

    try {
        const res = await fetch(`${import.meta.env.VITE_API_URL}/allGetOperations`, requestOptions);
        const data = await res.json();                
        return data;
    } catch (err) {
        console.error("Error in allGetOperations: ", err)
    }
}

/**
 * 
 * @param {number} opId Id der Operation
 * @returns Parameter einer Id
 * 
 */
async function loadParamsFromOPId(opId) {
    let antwort = JSON.stringify({id: opId});    
    const requestOptions = {
        method: 'POST',
        headers: {
        'Accept': 'application/text',
        'Content-Type': 'application/text'
        },
        body: antwort
    }
    try {
        const res = await fetch(`${import.meta.env.VITE_API_URL}/getParamsFromOPId`, requestOptions);
        const data = await res.json();        
        return data;
    } catch (err) {
        console.error("Error in getParamsFromOPId: ", err)
    }
}

export { laodAlloperations, loadParamsFromOPId }