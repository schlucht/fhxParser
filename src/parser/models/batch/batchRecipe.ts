import { BoolFlag } from "../fhx";
import { Meta } from "../meta";
import { PfcAlgorithm } from "../pfc/pfcAlgorithm";
import { AttributeInstance } from "./attributeInstance";
import { FormulaParameter } from "./formualParameter";

export interface BatchRecipe {
    name: string;
    type: 'OPERATION' | string;
    category?: string;
    meta: Meta;
  
    description?: string;
    useEquipmentTrains?: BoolFlag;
    equipmentUnitClass?: string;
    author?: string;
    abstract?: string;
    batchUnits?: string;
    batchLength?: string;
    defaultBatchSize?: number;
    minimumBatchSize?: number;
    maximumBatchSize?: number;
    productCode?: string;
    productName?: string;
    recipeApprovalInfo?: string;
    version?: string;
    formulaParameters?: FormulaParameter[];
    attributeInstances?: AttributeInstance[];
    pfcAlgorithm?: PfcAlgorithm;
  }