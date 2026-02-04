
export type MatchResult<T> = { ok: true; value: T } | { ok: false };

export function rxExec(regex: RegExp, str: string): RegExpExecArray | null {
    const flags = regex.flags.replace('g', '');
    const safe = new RegExp(regex.source, flags);
    return safe.exec(str);
}
  
export function matchStrR(regex: RegExp, str: string, group = 1): MatchResult<string> {
  const m = rxExec(regex, str);
  if (!m || group < 0 || group >= m.length) return { ok: false };  
  const v = m[group];
  return typeof v === 'string' ? { ok: true, value: v } : { ok: false };
}

export function matchNumR(regex: RegExp, str: string, group = 1): MatchResult<number> {
  const r = matchStrR(regex, str, group);
  if (!r.ok) return r;
  const n = Number(r.value.trim());
  return Number.isNaN(n) ? { ok: false } : { ok: true, value: n };
}


export function readBlock(blockName: string, lines: string[]): string[][] {
    let block:string[] = []
    const blocks:string[][] = [];

    let startBlock = false;
    let curlies = 0;

    for(const line of lines) {
        const trimmed = line.trim();
        if(!startBlock) {
            if(trimmed.startsWith(blockName)){                
                startBlock = true;
                block.push(trimmed.trim());            
            }            
        } else {
            if(trimmed.startsWith("{")) {                
                curlies++;
                continue;
            }
            if(trimmed.startsWith("}")) {
                curlies--;
                if(curlies === 0) {
                    startBlock = false;
                    blocks.push(block);
                    block = [];
                }
                continue;
            }
            block.push(trimmed.trim());
        }
    }
    return blocks;
}

export function readSection(lines: string[], start: string, end: string = "" ): string[] {
    const sectionFormula: string[] = [];
    let isForumla = false;
    for(const line of lines) {
        const trimming = line.trim();
        if(!isForumla) {
            if(line.includes(start)) {
                isForumla = true;
                sectionFormula.push(trimming);
            }
        } else {
            if (end != "") {
                if(line.includes(end)) break;
            }
            sectionFormula.push(trimming);
        }
    }
    return sectionFormula;
}
