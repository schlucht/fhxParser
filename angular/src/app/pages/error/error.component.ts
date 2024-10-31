import { Component, signal } from '@angular/core';

@Component({
  selector: 'ots-error',
  standalone: true,
  imports: [],
  templateUrl: './error.component.html',
  styleUrl: './error.component.css'
})
export class ErrorComponent {
  link = signal(location.pathname);

  ngOnInit() {
    this.link.set(location.pathname);   
  }
  
}
