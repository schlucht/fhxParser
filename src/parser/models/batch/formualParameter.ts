import { Rect } from "../fhx";

export type FormulaParameterType = 'BATCH_PARAMETER_REAL' | 'ENUMERATION_VALUE' | 'UNICODE_STRING';

export interface FormulaParameter {
  name: string;
  type: FormulaParameterType;
  connection?: 'INPUT' | 'OUTPUT' | 'INOUT' | 'CONSTANT';
  rectangle?: Rect;
  group?: string;
}