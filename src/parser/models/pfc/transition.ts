import { BoolFlag, Pos } from "../fhx";

export interface Transition {
    name: string;
    position?: Pos;
    termination: BoolFlag;
    expression: string;
  }
  