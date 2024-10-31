import { Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { ErrorComponent } from './pages/error/error.component';
import { OperationComponent } from './pages/operation/operation.component';
import { UnitComponent } from './pages/unit/unit.component';
import { RecipeComponent } from './pages/recipe/recipe.component';
import { UploadComponent } from './pages/upload/upload.component';
import { PlantComponent } from './pages/plant/plant.component';
import { LoginComponent } from './pages/login/login.component';
import { ImpressumComponent } from './pages/impressum/impressum.component';

export const routes: Routes = [
    {
        path: "home", 
        loadComponent: () => import('./pages/home/home.component')
                                .then(mod => HomeComponent)
    },
    {
        path: "operation", 
        loadComponent: () => import('./pages/operation/operation.component')
                                .then(mod => OperationComponent)
    },
    {
        path: "unit", 
        loadComponent: () => import('./pages/unit/unit.component')
                                .then(mod => UnitComponent)
    },
    {
        path: "recipe", 
        loadComponent: () => import('./pages/recipe/recipe.component')
                                .then(mod => RecipeComponent),        
    },
    {
        path: "upload", 
        loadComponent: () => import('./pages/upload/upload.component')
                                .then(mod => UploadComponent)
    },
    {path: "plant", 
        loadComponent: () => import('./pages/plant/plant.component')
                                .then(mod => PlantComponent)
    },
    {
        path: "login", 
        loadComponent: () => import('./pages/login/login.component')
                                .then(mod => LoginComponent)
    },
    {
        path: "impressum", 
        loadComponent: () => import('./pages/impressum/impressum.component')
                                .then(mod => ImpressumComponent)
    },

    {path: "", redirectTo: 'home', pathMatch: 'full'},
    {path: "**", loadComponent: () => import('./pages/error/error.component').then(mod => ErrorComponent)},
];


