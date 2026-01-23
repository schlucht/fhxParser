
import { Locale } from "../models/document/locale";
import { Meta } from "../models/meta";
import { matchStr } from "../utils";
import { parseMeta } from "./parseMeta";

export function parseLocale(localeLines: string[]): Locale {
    let meta: Meta = parseMeta(localeLines[0]);
    let locale = matchStr(/LOCALE="([^"]+)"/, localeLines[1]) || '';

    return {meta, locale};
}