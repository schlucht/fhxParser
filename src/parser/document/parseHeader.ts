import {Header} from "../models/fhx";



 export function parseHeader(v: string, d: string): Header  {
    
    const version = v.substring(3, v.length - 2).trim();
    const dateStr = d.substring(3, d.length - 2).trim();    

    return {version, dateStr};
}











