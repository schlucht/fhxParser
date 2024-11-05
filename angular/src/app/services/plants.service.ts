import { Injectable } from '@angular/core';
import { Plant } from '../models/plant';
import { Observable } from 'rxjs/internal/Observable';

const plants: Plant[] = [
  {plant_id: 1, plant: 'MZA'},
  {plant_id: 2, plant: 'E16'},
  {plant_id: 3, plant: 'Penta'},
  {plant_id: 4, plant: 'MZA Alpha'},
];

@Injectable({
  providedIn: 'root'
})
export class PlantsService {
  constructor() { }

  getPlants(): Observable<Plant[]> {
    return new Observable<Plant[]>(observer => {
      observer.next(plants);
    });    
  }
}
