
import { Meta } from "../models/meta";
import { Schema } from "../models/document/schema";
import { Version } from "../models/document/version";
import { matchNum, matchStr } from "../utils";
import { parseMeta } from "./parseMeta";

export function parseSchema(schemaLines: string[]): Schema {    
    const version: Version = {};
    let meta: Meta = parseMeta(schemaLines[0]);
    for(const line of schemaLines) {
        let m = matchNum(/VERSION=([0-9]+)(?:\/\*\s*"([^"]+)"\s*\*\/)?/, line) || undefined;
        if(m) version.versionRaw = m
        let b = matchStr(/VERSION=([0-9]+)(?:\/\*\s*"([^"]+)"\s*\*\/)?/, line, 2) || undefined;
        if (b) version.versionRawDateStr = b;
        m = matchNum(/MAJOR_VERSION=([0-9]+)/, line) || undefined;
        if (m) version.major = m;
        m = matchNum(/MINOR_VERSION=([0-9]+)/, line) || undefined;
        if (m) version.minor = m
        m = matchNum(/MAINTENANCE_VERSION=([0-9]+)/, line) || undefined;
        if(m) version.maintenance = m;
        m = matchNum(/BUILD_VERSION=([0-9]+)/, line) || undefined;
        if(m) version.build = m;
        b = matchStr(/BUILD_ID="([^"]+)"/, line) || undefined;
        if(b) version.buildId = b;
        b = matchStr(/VERSION_STR="([^"]+)"/, line) || undefined;
        if(b) version.versionStr = b;
        version.onlineUpgrade = (matchStr(/ONLINE_UPGRADE=([A-Z]+)/, line) === 'F' ? true : false) || undefined;       
        
    }
    return {meta, version}
}