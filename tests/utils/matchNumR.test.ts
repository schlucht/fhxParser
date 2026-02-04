
import { describe, it, expect } from "vitest";
import { MatchResult, matchNumR } from "../../src/parser/utils/utils";
import { Regex } from "../../src/parser/utils/const";

describe('util/matchNumR', () => {
    const line = 'user="FLAMBRIGGE" time=1596611920/* "05-Aug-2020 09:18:40" */'
    it('testen wenn nicht übergeben wird', () => {                
        const result = matchNumR(Regex.meta, '' );
        expect(result.ok).toBe(false);
    });

    it('test wenn user vorhanden ist', () => {
        const result = matchNumR(Regex.time, 'time=13245681' ) as MatchResult<number>;        
        expect(result.ok).toBe(true); 
        if(result.ok) expect(result.value).toBe(13245681);       
    });

    it('test wenn kein String zurückgegeben wird', () => {
        const result = matchNumR(Regex.user, 'xxx')
        expect(result.ok).toBe(false);
    });

    it('test wenn Nummer gefunden wird', () => {
        const result = matchNumR(Regex.time, "user=123456");
        expect(result.ok).toBe(false);

    });

    it('test gruppe nicht vorhanden', () => {
        const result = matchNumR(Regex.user, line, 10);
        expect(result.ok).toBe(false);
    });

});