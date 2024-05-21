import { OtsAlertMessage } from "./OtsAlert.js";

let globalPlant = {}

document.addEventListener('DOMContentLoaded', () => {
    
    const plantControl = document.getElementById('plant');
    const w= document.getElementById('window')
    const btnWindowClose = document.getElementById('btnWindowClose')
    const btnSavePlant = document.getElementById('btnSavePlant')
    const plantItems = document.getElementById('plantitems')

    btnWindowClose.addEventListener('click', closeWindow)
    btnSavePlant.addEventListener('click', savePlant)

    /**
 * Fetches all plants from the '/plants/allPlants' endpoint and populates the 'plantItems' select element with the retrieved data.
 *
 * @return {Promise<void>} - A promise that resolves when the plants are successfully fetched and the 'plantItems' select element is populated.
 */
    const plants = async () => {
        try {
            const res = await fetch('/plants/allPlants')
            const data = await res.json()
            for (let plant of data) {               
                plantItems.innerHTML += `<option value="${plant.plant_id}">${plant.plant}</option>`
            }
        } catch (error) {
            console.error("Plant load in main.js", error)
        }
    }   
    let plant = window.localStorage.getItem('plant')      
    if (!plant) {
        w.style.display = 'flex'
        plantControl.innerHTML = '<leer>'
        plants()
    } else {
        const p = JSON.parse(plant)
        plantControl.innerHTML = p.plant
        w.style.display = 'none'
    }
    
    /**
 * Closes the window by hiding it.
 *
 * @return {void} No return value.
 */
    function closeWindow() {
        console.log("Hallo")
        w.style.display = 'none'
    }

/**
 * Saves the selected plant to local storage and updates the plant control element.
 *
 * @return {void} No return value.
 */
    function savePlant() {       
       globalPlant = {
        plant_id: plantItems.value, 
        plant: plantItems[plantItems.selectedIndex].textContent
    }
       window.localStorage.setItem('plant', JSON.stringify(globalPlant))
       plantControl.innerHTML = globalPlant.plant
       w.style.display = 'none'
    }
    
    
})