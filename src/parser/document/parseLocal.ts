
import { Locale } from "../models/document/locale";
import { Meta } from "../models/meta";
import { Regex } from "../utils/const";
import { matchStrR } from "../utils/utils";
import { parseMeta } from "./parseMeta";

export function parseLocale(localeLines: string[]): Locale {
    const locale: Locale = {meta: {user: '', time: 0}, locale: ''};

    const meta: Meta = parseMeta(localeLines[1]);   
    if(meta) locale.meta = meta;

    const loc = matchStrR(Regex.local, localeLines[2]);    
    if(loc.ok) locale.locale = loc.value;

    return locale;
}