import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Plant } from '../../../models/plant';

@Component({
  selector: 'ots-plant-item',
  standalone: true,
  imports: [],
  templateUrl: './plant-item.component.html',
  styleUrl: './plant-item.component.css'
})
export class PlantItemComponent {
  @Input() plant:Plant = {plant_id: 0, plant: ''};
  @Output() change = new EventEmitter<Plant>();
  @Output() select = new EventEmitter<string>();
  
  onChange(p:Plant) {
    this.change.emit(p);
  }
  
  onSelectPlant(name: string) {
    this.select.emit(name);
  }
}
