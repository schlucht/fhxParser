import { Pos, Rect } from "../models/fhx";

export function parseRectangle(rec: string): Rect {  

    if(!rec) throw new Error("Kein Daten übergeben!");
    if(!rec.includes("RECTANGLE")) throw new Error("Kein Rectangle vorhanden");

    const RECTANGLE_REGEX =
    /RECTANGLE\s*=\s*\{\s*X\s*=\s*(-?\d+)\s+Y\s*=\s*(-?\d+)\s+H\s*=\s*(-?\d+)\s+W\s*=\s*(-?\d+)\s*\}/i;
    const rect: Rect = { x: 0, y: 0, h: 0, w: 0 };
    const m = RECTANGLE_REGEX.exec(rec.trim());
    if (!m) return rect

    const [, x, y, h, w] = m;
    return {
        x: Number(x),
        y: Number(y),
        h: Number(h),
        w: Number(w),    
    };
}

export function parsePosition(pos: string): Pos {    
    if(!pos) throw new Error("Kein Daten übergeben!");
    if(!pos.includes("POSITION")) throw new Error("Kein Position vorhanden");

    const POSREGEX = /POSITION\s*=\s*\{\s*X\s*=\s*(-?\d+)\s+Y\s*=\s*(-?\d+)\s+}/i;
    const posi: Pos = { x: 0, y: 0 };
    const m = POSREGEX.exec(pos.trim());
    if (!m) return posi

    const [, x, y] = m;
    return {
        x: Number(x),
        y: Number(y),           
    };
}