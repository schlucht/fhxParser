async function laodAllRecipes(plantId) {
    let antwort = JSON.stringify({id: plantId})
    console.log(antwort)
    const requestOptions = {
        method: 'POST',
        headers: {
        'Accept': 'application/text',
        'Content-Type': 'application/text'
        },
        body: antwort
    }
    try {
        const res = await fetch(`${import.meta.env.VITE_API_URL}/allGetRecipes`, requestOptions);
        const data = await res.json();        
        return data;
    } catch (err) {
        console.error("Error in allGetRecipes: ", err)
        return {
            ok: false,            
            message: "Fehler in laodAllRecipes: " + e,
            content: {},
            id: 1,
        }
    }
}

export { laodAllRecipes }