console.log('Start APP')

// Globale gespeicherte Variablen
 let plant =  {
    storeId: "plant",
    plantId: "",
    plantName: "keine",
}

const plants = []

const plantH3 = document.querySelector('#plant')

readLocalStorage()

//Laden der Anlage aus dem Localstorage
function readLocalStorage() {
    const store = JSON.parse(localStorage.getItem(plant.storeId))

    if (store) {
        plant.plantId = store.plantId
        plant.plantName = store.plantName      
    } else {
        localStorage.clear()
        const newStore = JSON.stringify(plant)
        localStorage.setItem('plant', newStore)
    }
  
    plantH3.innerHTML =`ANLAGE:  ${plant.plantName}`
}

function saveActivePlant(newPlant) {
    localStorage.clear();
    plant = newPlant;
    const newStore = JSON.stringify(plant);
    localStorage.setItem('plant', newStore);
    plantH3.innerHTML = `ANLAGE: ${plant.plantName}`;
}

// Extrahiert den Titel aus einem Text aus
function readTitleName(text, substring) {
    const lines = text.split('\n')
    const regexp = new RegExp(`${substring}"(?<f>.*)" `)
    let res = "none"
    if (lines.length > 0) {
        lines.forEach(l => {
            if (l.indexOf(substring) > -1) {
                console.log(regexp)
                const matches = l.match(regexp)
                if (matches) {
                    res = matches.groups.f
                }
            } 
        })
    }
    return res
}

