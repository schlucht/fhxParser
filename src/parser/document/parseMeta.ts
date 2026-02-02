import { Meta } from "../models/meta";
import { Regex } from "../utils/const";
import { matchNumR, matchStrR } from "../utils/utils";

export function parseMeta(metaString: string): Meta {
    const meta: Meta = {user: '', time: 0};
    
    const user = matchStrR(Regex.meta, metaString);
    if(user.ok) meta.user = user.value

    const time = matchNumR(Regex.meta, metaString, 2);
    if(time.ok) meta.time = time.value;

    const timeStr = matchStrR(Regex.meta, metaString, 3);
    if(timeStr.ok) meta.timeStr = timeStr.value;

    return meta;
}
