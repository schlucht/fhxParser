export function match(regex: RegExp, str: string) {
    return str.match(regex);
}

export function matchStr(regex: RegExp, str: string, group = 1): string | null {
    const m = str.match(regex);
    return m ? m[group] : null;
}

export function matchNum(regex: RegExp, str: string, group = 1): number | null {
    const m = str.match(regex);
    return m ? Number(m[group]) : null;
}

export function readBlock(blockName: string, lines: string[]): string[] {

    const block:string[] = []
    let startBlock = false;
    let curlies = 0;

    for(const line of lines) {
        const trimmed = line.trim();
        if(!startBlock) {
            if(trimmed.startsWith(blockName)){                
                startBlock = true;            
            }            
        } else {
            if(trimmed.includes("{")) {                
                curlies++;
                continue
            }
            if(trimmed.includes("}")) {
                curlies--;
                if(curlies === 0) {
                    // break;
                    return block;
                }
            }
            block.push(trimmed.trim());
        }
    }
    return block;
}