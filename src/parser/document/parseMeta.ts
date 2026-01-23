import { Meta } from "../models/meta";
import { match } from "../utils";

export function parseMeta(metaString: string): Meta {
    const regex = /user="([^"]+)"\s+time=([0-9]+)(?:\/\*\s*"([^"]+)"\s*\*\/)?/;

    const matches = match(regex, metaString);
    if(!matches) throw new Error(`Invalid meta string: ${metaString}`);

    return {
        user: matches[1],
        time: parseInt(matches[2]),
        timeStr: matches[3] ?? undefined
    }   
}