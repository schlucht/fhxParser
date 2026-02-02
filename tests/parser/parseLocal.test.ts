import { describe, it } from "vitest";
import { parseLocale } from "../../src/parser/document/parseLocal";
import { expect } from "vitest";

describe('parseLocale', () => {
    it('Local section parsen', () => {
        const input = [
          'LOCALE',
          'user="FLAMBRIGGE" time=1596611920/* "05-Aug-2020 09:18:40" */',
          'LOCALE="English_United States.1252"'
        ]
        
        const result = parseLocale(input);
        // console.log(result)
        
        expect(result.meta.user).toBe('FLAMBRIGGE');
        expect(result.meta.time).toBe(1596611920);
        expect(result.meta.timeStr).toBe('05-Aug-2020 09:18:40')
        expect(result.locale).toBe('English_United States.1252')

    });
});