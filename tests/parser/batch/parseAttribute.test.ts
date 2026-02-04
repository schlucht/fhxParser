
import { describe, it, expect } from "vitest";
import { parseAttributeBlock } from "../../../src/parser/batch/parseFormulaParameter";
import { parseValueEnum } from "../../../src/parser/batch/parseAttributeParams";

describe("parseAttributeBlock – DDT Tests", () => {

  //
  // 🔥 1. Fehlerfälle
  //
  it.each([
    {
      title: "wirft Fehler bei leerem Block",
      input: [],
      error: "Kein Daten übergeben!"
    },
    {
      title: "wirft Fehler wenn kein FORMULA_PARAMETER existiert",
      input: [["CONNECTION=INPUT"]],
      error: "Kein FormulaParameter vorhanden"
    }
  ])("$title", ({ input, error }) => {
    expect(() => parseAttributeBlock(input as any)).toThrow(error);
  });

  //
  // 🔥 2. Erfolgsfälle
  //
  it("parsed 3 FormulaParameter korrekt", () => {
    const input1 = [
        [
            'ATTRIBUTE_INSTANCE NAME="FP_PAHH"',
            'VALUE { DESCRIPTION="" HIGH=20000 LOW=-1000 SCALABLE=F CV=1200 UNITS="mbar" }'
          ],
          [
            'ATTRIBUTE_INSTANCE NAME="FP_PAL"',
            'VALUE { DESCRIPTION="" HIGH=20000 LOW=-1000 SCALABLE=F CV=-200 UNITS="mbar" }'
        ],
    ];
    const input2 = [
        [
            'ATTRIBUTE_INSTANCE NAME="FP_PV_KENNLINIE"',
            'VALUE',
            'SET="L_EIN_AUS"',
            'STRING_VALUE="AUS"',
            'CHANGEABLE=F'
          ],
    ];

    const input3 = [
        [
            'ATTRIBUTE_INSTANCE NAME="FP_BESCHREI_TEXT"',
            'VALUE { CV="" }',  
        ]
    ];

    const result1 = parseAttributeBlock(input1);
    const result2 = parseAttributeBlock(input2);
    const result3 = parseAttributeBlock(input3);    

    // Beispielprüfungen für den ersten Block
    const res = result1[0]   
    expect(res).toMatchObject({
      name: "FP_PAHH",   
      value: {
        description: "",
        high: 20000,
        low: -1000,
        scalable: false,
        cv: 1200,
        units: "mbar"
      }   
    });    
    expect(result2[0]).toMatchObject({
      name: "FP_PV_KENNLINIE",   
      value: {
        set: 'L_EIN_AUS',
        stringValue: 'AUS',
        changeable: false,
      }   
    });    
    expect(result3[0]).toMatchObject({
      name: "FP_BESCHREI_TEXT",   
      value: {
        cv: ''
      }   
    });    
  });
});

describe("enum Attribute", () => {

  it('Keine Daten übergeben!', () => {
    expect(() => parseValueEnum([])).toThrow("Kein Daten übergeben!");
  });
  
  it('Falsche Daten übergeben!', () => {    
    expect(() => parseValueEnum([
      'ATTRIBUTE_INSTANCE NAME="FP_OPTION"',
      'VALUE'])).toThrow("Falsche Daten übergeben!");
  }); 

  it('korrekte Daten', () => {
    const input =  [
      'ATTRIBUTE_INSTANCE NAME="FP_OPTION"',
      'VALUE',
      'SET="LGF_DRUCK_OPT"',
      'STRING_VALUE="ABLUFT1_REGLER"',
      'CHANGEABLE=F'
    ];

    const result = parseValueEnum(input);
    expect(result).toMatchObject({            
      set: 'LGF_DRUCK_OPT',
      stringValue: 'ABLUFT1_REGLER',
      changeable: false,
    });
  });

});