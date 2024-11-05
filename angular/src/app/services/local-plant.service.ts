import { Injectable } from '@angular/core';
import { Plant } from '../models/plant';

@Injectable({
  providedIn: 'root'
})
export class LocalPlantService {

  public readonly key = 'plant';

  save(plant: Plant) {
    localStorage.setItem(this.key, JSON.stringify(plant));
  } 
  
  read(): Plant {
    let p = localStorage.getItem(this.key);
    if (p) {
      return JSON.parse(p);
    } else {
      return new Plant(0, '');
    }
  }
}
