
async function loadOperation(plantId) {
    const res = await fetch(api + "/operation/" + plantId + "/all");
    const data = await res.json();    
    return data;       
}

export { loadOperation };