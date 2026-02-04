import { EnumAttr, NumericAttr, StringAttr} from "../models/batch/attributeInstance";
import { matchStrR, rxExec } from "../utils/utils";
import { Regex } from "../utils/const";

// Parst den Numeric Wert eines Attributes:
// ATTRIBUTE_INSTANCE NAME="FP_PAHH"
//   {
//     VALUE { DESCRIPTION="" HIGH=20000 LOW=-1000 SCALABLE=F CV=1200 UNITS="mbar" }
//   }
// return: { description: '', high: 20000, low: -1000, scalable: false, cv: 1200}
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

// Parst enum Values aus Attribute
// ATTRIBUTE_INSTANCE NAME="FP_PV_KENNLINIE"
//   {
//     VALUE
//     {
//       SET="L_EIN_AUS"
//       STRING_VALUE="AUS"
//       CHANGEABLE=F
//     }
//   }
// return: { set: 'L_EIN_AUS', stringValue: 'AUS', changeable: false }
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

// Parst den String Wert eines Attributes:
// ATTRIBUTE_INSTANCE NAME="FP_BESCHREI_TEXT"
// {
//   VALUE { CV="" }
// }
// return: { cv: '' }
export function parseValueString(line: string): StringAttr {
    const trimmed = line.trim();
    const str = {
        cv: ''
    }
    const cv = matchStrR(Regex.cv, trimmed);

    if(cv.ok) {
        str.cv = cv.value;
    }
    return str;
}