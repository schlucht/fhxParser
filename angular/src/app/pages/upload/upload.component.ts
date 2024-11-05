import { NgClass } from '@angular/common';
import { Component, signal } from '@angular/core';
import { Observable } from 'rxjs/internal/Observable';

@Component({
  selector: 'ots-upload',
  standalone: true,
  imports: [NgClass],
  templateUrl: './upload.component.html',
  styleUrl: './upload.component.css'
})

export class UploadComponent {
  fileContent = signal<string>('Resultat...');
  fileName = signal<string>( '...');
  onFileSelected(event: Event) {
    const fileInput = event.target as HTMLInputElement;
    if (fileInput.files && fileInput.files.length > 0) {
      const file = fileInput.files[0];
      this.fileName.set(file.name);      
      const res = this.readFileContent(file);
      res.subscribe(res =>  {
        this.fileContent.set("..." + res.slice(390, 500) + "...");
      });      
    }    
  }

  private readFileContent(file: File): Observable<string> {
    return new Observable<string>((observer) => {
      const reader = new FileReader();
      reader.onload = () => {
        observer.next(reader.result as string);
        observer.complete();
      };
      reader.onerror = (error) => {
        observer.error(error);
      };
      reader.readAsText(file);    
    });
  }

}
