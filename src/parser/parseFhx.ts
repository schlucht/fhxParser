import { FhxDocument } from "./models/fhx";
import { readBlock } from "./utils";
import { parseHeader } from "./document/parseHeader";
import { parseLocale } from "./document/parseLocal";
import { parseSchema } from "./document/parseSchema";

export function parseFhx(txtFhx: string[]): FhxDocument | null {

    const header = parseHeader(txtFhx[0], txtFhx[1]);  
    const schemaLines = readBlock("SCHEMA", txtFhx);
    const localLines = readBlock("LOCALE", txtFhx);
    const schema = parseSchema(schemaLines);
    const locale = parseLocale(localLines);

    console.log(JSON.stringify(locale));

    return null;
}