
/**
 * Alle Operation anhand der Anlagen ID einlesen
 * @param {int} plantId 
 */
async function laodAlloperations(plantId) {       

    let antwort = JSON.stringify({plantId})
    
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
        const data = res.json();
        return data;
    } catch (err) {
        console.error("Error in FileUpload: ", err)
    }
}

export { laodAlloperations }