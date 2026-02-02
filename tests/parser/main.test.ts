import { describe, test, vi, expect, } from "vitest";

describe('parse/main', () => {
    test('create() ruft eine fhxDatei mit dem Pfad auf', async () => {
        
        vi.mock('../../src/parser/readFile', () => {
            return  { readFhxFile: vi.fn(), }
        });
        const { default: fhx } = await import('../../src/parser');
        const { readFhxFile } = await import('../../src/parser/readFile');
        const path = '../../files/OP_DRUCK.fhx';
        fhx.create(path);
        expect(readFhxFile).toHaveBeenCalledWith(path);
        expect(readFhxFile).toHaveBeenCalledTimes(1);
    });    
});