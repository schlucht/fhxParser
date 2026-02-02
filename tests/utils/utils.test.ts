import { describe, expect, it } from "vitest";
import { Regex } from "../../src/parser/utils/const";
import { rxExec } from "../../src/parser/utils/utils";

describe('utils/rxExec', () => {
    it('liefert einen treffer user', () => {
        const line = 'user="FLAMBRIGGE" time=1596611920/* "05-Aug-2020 09:18:40" */'
        
        const result = rxExec(Regex.user, line);

        const [, dq, sq, bare] = result!;
        const value = dq ?? sq ?? bare;
        expect(value).toBe('FLAMBRIGGE');
    });

    it('liefert einen treffer time', () => {
        const reg = new RegExp(Regex.user.source, Regex.user.flags + 'g');

        reg.lastIndex = 50;

        const line = 'xxx user="FLAMBRIGGE" yyy';
        const result = rxExec(reg, line);

        expect(result).not.toBeNull();
        const [, dq, sq, bare] = result!;
        const value = dq ?? sq ?? bare;
        expect(value).toBe('FLAMBRIGGE');

        expect(reg.lastIndex).toBe(50);    
   
    })
    
    
});