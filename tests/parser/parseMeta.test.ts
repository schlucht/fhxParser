import { describe, it, expect } from "vitest";
import { parseMeta } from "../../src/parser/document/parseMeta";

describe('parseMeta', () => {
    it('parsed die Metadaten', () => {

        const input = 'user="FLAMBRIGGE" time=1596611920/* "05-Aug-2020 09:18:40" */'
        const result = parseMeta(input);

        expect(result.user).toBe('FLAMBRIGGE');
        expect(result.time).toBe(1596611920);
        expect(result.timeStr).toBe('05-Aug-2020 09:18:40');
    });
});
