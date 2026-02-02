
import { Meta } from "../models/meta";
import { Schema } from "../models/document/schema";
import { Version } from "../models/document/version";
import { matchNumR, matchStrR } from "../utils/utils";
import { parseMeta } from "./parseMeta";
import { Regex } from "../utils/const";

export function parseSchema(schemaLines: string[]): Schema {    
    const version: Version = {};
    const meta: Meta =  schemaLines.length > 1 
        ? parseMeta(schemaLines[1])
        : {user: '', time: 0};

    for(const line of schemaLines) {
        const trimmed = line.trim();
        if(trimmed.length === 0) continue;
        
        const versRaw = matchNumR(Regex.version, line);
        if(versRaw.ok) version.versionRaw = versRaw.value;
        
        const versStr = matchStrR(Regex.versionStr, line, 2);
        if (versStr.ok) version.versionStr = versStr.value;

        const maj = matchNumR(Regex.major, line);
        if(maj.ok) version.major = maj.value;

        const min = matchNumR(Regex.minor, line);
        if(min.ok) version.minor = min.value;

        const maint = matchNumR(Regex.maintenance, line);
        if(maint.ok) version.maintenance = maint.value

        const build = matchNumR(Regex.build, line);
        if(build.ok) version.build = build.value;

        const buildId = matchStrR(Regex.buildId, line);
        if(buildId.ok) version.buildId = buildId.value;
        
        const versionStr = matchStrR(Regex.versionStr, line, 2);
        if(versionStr.ok) version.versionStr = versionStr.value;
        
        const onlineUpg = matchStrR(Regex.upgrade, line);
        if(onlineUpg.ok) version.onlineUpgrade = onlineUpg.value === 'F' ? false : true;
    }
    return {meta, version}
}