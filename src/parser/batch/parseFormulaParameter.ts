import { FormulaParameter, FormulaParameterType } from "../models/batch/formualParameter";
import { matchStrR, readBlock, readSection } from "../utils/utils";
import { Regex } from "../utils/const";
import { parseRectangle } from "./rectangel";
import { AttributeInstance } from "../models/batch/attributeInstance";
import { parseValueEnum, parseValueNumeric, parseValueString } from "./parseAttributeParams";

export function parseFormulaParameter(lines: string[]): FormulaParameter[]{
    
    const sectionFormula = readSection(lines,'FORMULA_PARAMETER', 'ATTRIBUTE_INSTANCE'); 

    const blockFormula = readBlock('FORMULA_PARAMETER',sectionFormula);
  
    const parseFormula : FormulaParameter[] = parseFormulaBlock(blockFormula);
    
    
    return parseFormula
}

export function parseAttributeInstance(lines: string[]): AttributeInstance[]{
    const sectionAttribute = readSection(lines, 'ATTRIBUTE_INSTANCE', 'PFC_ALGORITHM' );
    const blockAttribute = readBlock('ATTRIBUTE_INSTANCE',sectionAttribute);   
    const parseAttribute : AttributeInstance[] = parseAttributeBlock(blockAttribute);    
    return parseAttribute;
}



export function parseAttributeBlock(block: string[][]): AttributeInstance[] {
    if (block.length === 0) 
        throw new Error("Kein Daten übergeben!");
    if (!block[0][0].includes('ATTRIBUTE_INSTANCE')) 
        throw new Error("Kein FormulaParameter vorhanden");

    const attributeInstances: AttributeInstance[] = [];
    for(let b of block) {
        const attribute: AttributeInstance = {
            name: "",
            value: {}
        }
        for(let i = 0; i < b.length; i++) {
            const trimmed = b[i].trim();
            if(trimmed.startsWith('ATTRIBUTE_INSTANCE')) {
                const name = matchStrR(Regex.name, b[i]);
                if(name.ok) {
                    attribute.name = name.value;
                }                
                continue;            
            }
            if(trimmed.includes('DESCRIPTION')) {
                const value = parseValueNumeric(trimmed);
                attribute.value = value;
                continue;
            }

            if(trimmed.includes('SET')) {
                const value = parseValueEnum(b.slice(i, i + 3));
                attribute.value = value;
            }

            if(trimmed.startsWith('VALUE { CV="" }')) {
                const value = parseValueString(trimmed);
                attribute.value = value;
            }
        }
        attributeInstances.push(attribute);
    }
    return attributeInstances;
}

export function parseFormulaBlock(block: string[][]): FormulaParameter[] {
    
    if (block.length === 0) 
        throw new Error("Kein Daten übergeben!");
    if (!block[0][0].includes('FORMULA_PARAMETER')) 
        throw new Error("Kein FormulaParameter vorhanden");    
    const formulaParameters: FormulaParameter[] = [];  
    for(let b of block) {
        const formular: FormulaParameter = {
            name: "",
            type: 'BATCH_PARAMETER_REAL',
            connection: 'CONSTANT',
            rectangle: { x: 0, y: 0, h: 0, w: 0 },
            group: ""
        }
        for(let l of b) {
            const trimmed = l.trim();
            if(trimmed.startsWith('FORMULA_PARAMETER')) {
                const name = matchStrR(Regex.formParam, trimmed);
                if(name.ok) {
                    formular.name = name.value;
                }
                const type = matchStrR(Regex.formParam, trimmed, 2);
                if(type.ok) {
                    formular.type = type.value as FormulaParameterType;
                }
                continue;                
            }
            if(trimmed.startsWith('CONNECTION')) {
                const conn = matchStrR(Regex.connection, trimmed);
                if(conn.ok) {
                    formular.connection = conn.value as FormulaParameter['connection'];
                }
                continue;
            }
            if(trimmed.startsWith('RECTANGLE')) {
                const rec = parseRectangle(trimmed);
                formular.rectangle = rec;
                continue;
            }
            if(trimmed.startsWith('GROUP')) {
                const group = matchStrR(Regex.group, trimmed);
                if(group.ok) {
                    formular.group = group.value;
                }
                continue;
            }
        }
        formulaParameters.push(formular);
    }    
    return formulaParameters;
}