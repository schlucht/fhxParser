import { BoolFlag } from "../fhx";

export interface AttributeInstance {
    name: string;
    value: AttributeValue;
  }

  type AttributeValue =
    | {
        // numerisch skalierte Werte
        description?: string;
        high?: number;
        low?: number;
        scalable?: BoolFlag;
        cv?: number;              // Current Value
        units?: string;
      }
    | {
        // Enum/String-Werte
        set?: string;
        stringValue?: string;
        changeable?: BoolFlag;
      }
    | {
        // einfache String CV
        cv?: string;
      };