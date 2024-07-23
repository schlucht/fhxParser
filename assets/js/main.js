console.log('Start APP')

// Globale gespeicherte Variablen
 let plant =  {
    storeId: "plant",
    plantId: "",
    plantName: "keine",
}

const plants = []

const plantH3 = document.querySelector('#plant')
const noPlant = document.getElementById('no-plant')

readLocalStorage()

//Laden der Anlage aus dem Localstorage
function readLocalStorage() {
    const store = JSON.parse(localStorage.getItem(plant.storeId))    
    if (store) {
        plant.plantId = store.plantId
        plant.plantName = store.plantName
        if (plant.plantId = "") {
            noPlant.style.display = 'flex'
        } else {
            noPlant.style.display = 'none'
        }
    } else {
        localStorage.clear()
        const newStore = JSON.stringify(plant)
        localStorage.setItem('plant', newStore)
    }
    plantH3.innerHTML =`ANLAGE:  ${plant.plantName}`
}


