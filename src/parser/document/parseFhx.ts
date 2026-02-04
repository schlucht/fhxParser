import { FhxDocument } from "../models/fhx";
import { Blocks } from "../utils/const";
import { readBlock } from "../utils/utils";
import { parseHeader } from "./parseHeader";
import { parseLocale } from "./parseLocal";
import { parseSchema } from "./parseSchema";
import { parseRecipe } from "../batch/parseBatchRecipe";

export function parseFhx(txtFhx: string[]): FhxDocument | null {

    const header = parseHeader(txtFhx[0], txtFhx[1]);  
    const schemaLines = readBlock(Blocks.schema, txtFhx);
    const localLines = readBlock(Blocks.local, txtFhx);
    const schema = parseSchema(schemaLines[0]);
    const locale = parseLocale(localLines[0]);
    const recipes = parseRecipe(txtFhx);

    

    return {header, schema, locale, recipes};
}