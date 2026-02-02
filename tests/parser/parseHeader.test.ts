
import { describe, it, expect } from "vitest";
import { parseHeader } from "../../src/parser/document/parseHeader";

describe("readHeader", () => {
  it("parst einen einfachen Rezept-Header korrekt", () => {
    const input = [
      '/* Version: 10.3.1.3657.xr */',
        '/* "05-Aug-2020 09:24:47" */',      
    ];

    
    const result = parseHeader(input[0], input[1]);

    expect(result.version).toBe("Version: 10.3.1.3657.xr");
    expect(result.dateStr).toBe('"05-Aug-2020 09:24:47"');

  });
});
