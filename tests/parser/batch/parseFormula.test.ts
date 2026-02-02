
import { describe, it, expect } from "vitest";
import { parseFormulaBlock } from "../../../src/parser/batch/parseFormulaParameter";

describe("parseFormulaBlock – DDT Tests", () => {

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
    expect(() => parseFormulaBlock(input as any)).toThrow(error);
  });

  //
  // 🔥 2. Erfolgsfälle
  //
  it("parsed 3 FormulaParameter korrekt", () => {
    const input = [
      [
        'FORMULA_PARAMETER NAME="FP_FSB_PSH" TYPE=BATCH_PARAMETER_REAL',
        'CONNECTION=INPUT',
        'RECTANGLE= { X=-50 Y=-50 H=1 W=1 }',
        'GROUP="Operating"'
      ],
      [
        'FORMULA_PARAMETER NAME="FP_FSB_PSL" TYPE=BATCH_PARAMETER_REAL',
        'CONNECTION=INPUT',
        'RECTANGLE= { X=-50 Y=-50 H=1 W=1 }',
        'GROUP="Operating"'
      ],
      [
        'FORMULA_PARAMETER NAME="FP_OPTION" TYPE=ENUMERATION_VALUE',
        'CONNECTION=INPUT',
        'RECTANGLE= { X=-50 Y=-50 H=1 W=1 }',
        'GROUP="Operating"'
      ]
    ];

    const result = parseFormulaBlock(input);

    expect(result.length).toBe(3);

    // Beispielprüfungen für den ersten Block
    expect(result[0]).toMatchObject({
      name: "FP_FSB_PSH",
      type: "BATCH_PARAMETER_REAL",
      connection: "INPUT",
      rectangle: { x: -50, y: -50, h: 1, w: 1 },
      group: "Operating"
    });

    expect(result[2].type).toBe("ENUMERATION_VALUE");
  });
});

