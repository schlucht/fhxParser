import { Rect } from "../fhx";
import { StepParameter } from "./stepParameter";

export interface Step {
    name: string;
    definition: string;
    description?: string;
    rectangle?: Rect;
    keyParameter?: string;
    stepParameters?: StepParameter[];
  }