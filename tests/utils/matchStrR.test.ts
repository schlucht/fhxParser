
import { describe, it, expect } from 'vitest';
import { matchStrR } from '../../src/parser/utils/utils';
import { Regex } from '../../src/parser/utils/const';

describe('util/matchStrR', () => {
  const line = 'user="FLAMBRIGGE" time=1596611920/* "05-Aug-2020 09:18:40" */';

  it('ok:false bei leerem Input', () => {
    const result = matchStrR(Regex.user, '');
    expect(result.ok).toBe(false);
  });

  it('ok:true und value bei user vorhanden', () => {
    const result = matchStrR(Regex.user, line);
    expect(result.ok).toBe(true);
    if (result.ok) expect(result.value).toBe('FLAMBRIGGE');
  });

  it('ok:false wenn kein Treffer', () => {
    const result = matchStrR(Regex.user, 'xxx');
    expect(result.ok).toBe(false);
  });

  it('ok:false wenn falsches Pattern (time auf user-Line)', () => {
    const result = matchStrR(Regex.time, 'user=123456');
    expect(result.ok).toBe(false);
  });

  it('ok:false wenn Gruppe außerhalb des Bereichs', () => {
    const result = matchStrR(Regex.user, line, 10);
    expect(result.ok).toBe(false);
  });

  it('liefert Full Match bei group=0', () => {
    const res = matchStrR(Regex.user, 'user="FLAMBRIGGE"', 0);
    expect(res.ok).toBe(true);
    if (res.ok) expect(res.value).toMatch(/^user\s*=/i);
  }); 

  it('ignoriert lastIndex bei globalen Regexen (rxExec entfernt g)', () => {
    const reg = new RegExp(Regex.user.source, Regex.user.flags + 'g');
    reg.lastIndex = 100;
    const res = matchStrR(reg, 'xxx user="FLAMBRIGGE" yyy');
    expect(res.ok).toBe(true);
    if (res.ok) expect(res.value).toBe('FLAMBRIGGE');
    expect(reg.lastIndex).toBe(100);
  });

  // Nur wenn Regex.user Alternativ-Gruppen nutzt:
  it('ok:false, wenn gewählte Gruppe nicht gematcht hat (undefined)', () => {
    const res = matchStrR(Regex.user, 'user="FLAMBRIGGE"', 2); // SQ-Gruppe leer
    expect(res.ok).toBe(false);
  });

  // Optional: Multiline/Unicode
  it('matcht in mehrzeiligem Input', () => {
    const text = `aaa
bbb
user="FLAMBRIGGE"
ccc`;
    const res = matchStrR(Regex.user, text);
    expect(res.ok).toBe(true);
    if (res.ok) expect(res.value).toBe('FLAMBRIGGE');
  });

  it('funktioniert mit Unicode im Wert', () => {
    const res = matchStrR(Regex.user, 'user="Jörg-Świętopełk"');
    expect(res.ok).toBe(true);
    if (res.ok) expect(res.value).toBe('Jörg-Świętopełk');
  });
});
