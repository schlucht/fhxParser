import { Plant } from './plant.js';
document.addEventListener('DOMContentLoaded', () => {  

console.log('Start APP');
const plantH3 = document.querySelector('#plantName');
const plant = new Plant();

function loadPlant() {
    console.log(plant.print());
    if (!plant.hasPlant()) {
        document.location.href = '/plant';
        return;
    } else {
        plantH3.innerHTML = plant.plantName;
    }
}

// Die Ausgewählte Anlage in der Lokalstorage speichern.
// const saveButton = document.getElementById('save-plant-store')
const rdB = document.querySelectorAll("input[name='newPlant']")

// saveButton.addEventListener('click', (e) => {
//     rdB.forEach(element => {            
//         if (element.checked) {           
//             plant.plantId = element.id,
//             plant.plantName = element.value,                
//             localStorage.setItem(plant.storeId, JSON.stringify(plant))
//             readLocalStorage()
//         }        
//     });
// })

const plants = [];


readLocalStorage()

//Laden der Anlage aus dem Localstorage
function readLocalStorage() {
    const store = JSON.parse(localStorage.getItem(plant.storeId));
    if (store) {
        plant.plantId = store.plantId;
        plant.plantName = store.plantName;    
    } 
    saveActivePlant(plant);
    return plant;
}

function saveActivePlant(newPlant) {
    localStorage.clear();   
    localStorage.setItem('plant', JSON.stringify(newPlant));
    plantH3.innerHTML = `ANLAGE: ${newPlant.plantName}`;
    activateNavigation(newPlant.plantId);
}

function activateNavigation(plantid) {
    const nav = document.querySelectorAll(".nav-left li a");
    console.dir(nav);
    nav.forEach((n) => {
        let attr = n.getAttribute('href');
        n.setAttribute('href', attr + '/' + plantid);
    });
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
});
