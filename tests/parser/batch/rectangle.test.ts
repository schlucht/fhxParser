import { describe, it, expect } from "vitest";
import { parsePosition, parseRectangle } from "../../../src/parser/batch/rectangel";

describe("batch/rectangle", () => {
    const input = 'RECTANGLE= { X=-50 Y=-50 H=1 W=1 }';
    const inputPos = 'POSITION= { X=90 Y=100 }';

     it.each([
        {
          title: "wenn Daten übergeben werden",
          input: '',
          error: "Kein Daten übergeben!"
        },
        {
          title: "wirft Fehler wenn kein RECTANGLE existiert",
          input: "CONNECTION=INPUT",
          error: "Kein Rectangle vorhanden"
        }
      ])("$title", ({ input, error }) => {
        expect(() => parseRectangle(input as any)).toThrow(error);
      });

    it("korrekt parsen", () => {
        const result = parseRectangle(input);       

        // Beispielprüfungen für den ersten Block
        expect(result).toMatchObject({
            x: -50,
            y: -50,
            h: 1,  
            w: 1
        });        
    });

    it("parsen Position", () => {
      const result = parsePosition(inputPos);
      expect(result).toMatchObject({
        x: 90,
        y: 100
      });
    });
    
});