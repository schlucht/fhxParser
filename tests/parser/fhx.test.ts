import { mock } from "node:test";
import { describe, expect, test, vi } from "vitest";

describe('src/main', () => {
    test('ruft fhx.create() mit Pfad auf', async () => {
        const path = '/home/user/fhxParser/files/OP_DRUCK.fhx';

        vi.mock('../../src/parser', () => {
            return {
                default: {create: vi.fn()}
            };
        });
        await import ('../../src/main');
        const  {default: mockedFhx}  = await import ('../../src/parser');
        
        expect(mockedFhx.create).toHaveBeenCalledWith('/home/user/fhxParser/files/OP_DRUCK.fhx');
        expect(mockedFhx.create).toHaveBeenCalledTimes(1);
    });
});
