import {Header} from "../models/fhx";



 export function parseHeader(v: string, d: string): Header  {
    
    const version = v.substring(3, v.length - 3).trim();
    const dateStr = d.substring(3, d.length - 3).trim();    

    return {version, dateStr};
}











