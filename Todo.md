# Daten laden

- (x) Op auf die Startseite geladen
- ( ) Seite anpassen Anzeigen der OP's

## File einlesen

### Auswahl einer fhx Datei
1. Webseite mit Input Typ 'File' erstellen
2. Auswahl einer fhx Datei aus dem Dateisystem
    - Kontrolle dass nur fhx Datei eingelesen wird
3. fhx String an API übergeben /api/read_fhx
    - `ReadFhxHandler()`
---
### Speichern fhx api
1. Einlesen der Datei mit `ReadFhxHandler()`
    - Parsen der Datei `return []Fhx`
    - Verbindung mit MySQL DB
    - Suchen UP nach Name. Wenn der Datensatz gefunden wird, dass duchgehen und die Daten anpassen. Bei den Value Daten die Daten mit einem Timestamp ein neuen DS erzeugen. Keine Änderungen vornehmen.
    - Wenn kein Datensatz gefunden wird, einen neuen anlegen und die Daten in die DB speichern.



### Git Kommandos

- add - Neues Feature implementiert
- fixed - Fehler behoben
- refact - Code Refactoring
- infra - Infrastruktur angepasst
    