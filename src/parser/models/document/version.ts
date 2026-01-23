import { BoolFlag } from "../fhx";

export interface Version {
    versionRaw?: number;     // VERSION=...
    versionRawDateStr?: string;
    major?: number;
    minor?: number;
    maintenance?: number;
    build?: number;
    buildId?: string;
    versionStr?: string;
    onlineUpgrade?: BoolFlag;
  }