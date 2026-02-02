
import {promises as fs} from 'fs';
import { parseFhx } from './document/parseFhx';

async function readFhxFile(path: string) {
    try {
        const data = await fs.readFile(path);
        const text = decodeWithBom(data);
        const lines = text.split(/\r?\n/);
        parseFhx(lines);
    } catch(err) {
        console.error(err);
    }
}

function decodeWithBom(buf: Buffer){
    // UTF-8
    if (buf.length > 3 &&
        buf[0] === 0xEF &&
        buf[1] === 0xBB && 
        buf[2] === 0xBF
    ) {
        const dec = new TextDecoder('utf-8');
        return dec.decode(buf.subarray(3));
    }

    // UTF-16LE
if (buf.length >= 2 && 
    buf[0] === 0xFF && 
    buf[1] === 0xFE) {
    const dec = new TextDecoder('utf-16le');
    return dec.decode(buf.subarray(2));
  }

  //UTF-16
  if (buf.length >= 2 && buf[0] === 0xFE && buf[1] === 0xFF) {
    // Byte-Swap nach LE (kein slice)
    const be = buf.subarray(2);
    const le = Buffer.allocUnsafe(be.length);
    for (let i = 0; i < be.length; i += 2) {
      le[i] = be[i + 1];
      le[i + 1] = be[i];
    }
    const dec = new TextDecoder('utf-16le');
    return dec.decode(le);
  }
  
    const dec = new TextDecoder('utf-8');
    return dec.decode(buf);
}

export {readFhxFile}