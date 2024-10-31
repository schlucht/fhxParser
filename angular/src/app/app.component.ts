import { Component, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { TopComponent } from './layouts/top/top.component';
import { HeaderComponent } from "./layouts/header/header.component";
import { FooterComponent } from './layouts/footer/footer.component';

@Component({
  selector: 'ots-root',
  standalone: true,
  imports: [RouterOutlet, TopComponent, HeaderComponent, FooterComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = signal('angular');
}
