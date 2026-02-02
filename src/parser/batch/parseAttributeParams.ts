import { EnumAttr, NumericAttr, StringAttr} from "../models/batch/attributeInstance";
import { matchStrR, rxExec } from "../utils/utils";
import { Regex } from "../utils/const";

export function parseValueNumeric(line: string): NumericAttr {
    const trimmed = line.trim();
    const numAttr = {
        description: '',
        high: 0,
        low: 0,
        scalable: false,
        cv: 0,
        units: ''
    }
    const num = rxExec(Regex.numAttr.numeric, trimmed);
    console.log(num)
    if(num) {
        numAttr.description = num[1];
        numAttr.high = parseInt(num[2]);
        numAttr.low = parseInt(num[3]);
        numAttr.scalable = num[4] === 'T';
        numAttr.cv = parseInt(num[5]);
        numAttr.units = num[6];
    }

    return numAttr;
}

export function parseValueEnum(lines: string[]): EnumAttr{
    
    if(lines.length === 0) throw new Error("Kein Daten übergeben!");
    if(lines.length < 3) throw new Error("Falsche Daten übergeben!");

        let en = {            
            set: '',
            stringValue: '',
            changeable: true,
        };

        for(let l of lines) {
            const set = matchStrR(Regex.stringAttr.set, l);
            if(set.ok) {
                en.set = set.value;
            }
            const sv = matchStrR(Regex.stringAttr.sv, l);
            if(sv.ok) {
                en.stringValue = sv.value;
            }
            const ch = matchStrR(Regex.stringAttr.ch, l);
            if(ch.ok) {
                en.changeable = ch.value === 'T';
            }            
        }
        return en;
}