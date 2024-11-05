import { NgClass } from '@angular/common';
import { Component, OnInit, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { Plant } from '../../models/plant';
import { PlantsService } from '../../services/plants.service';
import { PlantItemComponent } from './plant-item/plant-item.component';

@Component({
  selector: 'ots-plant',
  standalone: true,
  imports: [NgClass, FormsModule, PlantItemComponent],
  templateUrl: './plant.component.html',
  styleUrl: './plant.component.css',
})
export class PlantComponent implements OnInit {
  inputText: string = '';
  plantData = '';

  plantService = inject(PlantsService);
  plants:Plant[] = [];

  onSubmit() {
    console.log(this.inputText);
    this.inputText = '';
  }

  readPlant(p: Plant) {
    console.log("edit", p)
    this.inputText= p.plant;
  }

  saveLocal() {
    
  }
  ngOnInit() {
    this.plantService.getPlants().subscribe((plants:Plant[]) => this.plants = plants);
  }

}
