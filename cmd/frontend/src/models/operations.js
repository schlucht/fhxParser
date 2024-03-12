
/**
 * Alle Operation anhand der Anlagen ID einlesen
 * @param {int} plantId 
 */
// async function laodAlloperations(plantId) {       

//     if (!plantId) return;

//     let antwort = JSON.stringify({id: plantId})
//     console.log(antwort)
//     const requestOptions = {
//         method: 'POST',
//         headers: {
//         'Accept': 'application/text',
//         'Content-Type': 'application/text'
//         },
//         body: antwort
//     }

//     try {
//         const res = await fetch(`${import.meta.env.VITE_API_URL}/allGetOperations`, requestOptions);
//         const data = await res.json();                    
//         return data;
//     } catch (err) {
//         console.error("Error in allGetOperations: ", err)
//     }
// }
/**
 * Load all operations based on the plant ID
 * @param {int} plantId 
 */
async function loadAllOperations(plantId) {       
    if (!plantId) return;

    try {
        const res = await fetch(`${import.meta.env.VITE_API_URL}/allGetOperations`, {
            method: 'POST',
            headers: {
                'Accept': 'application/text',
                'Content-Type': 'application/text'
            },
            body: JSON.stringify({ id: plantId })
        });
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

export { loadAllOperations, loadParamsFromOPId }