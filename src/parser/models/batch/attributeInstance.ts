import { BoolFlag } from '../fhx';

export interface AttributeInstance {
	name: string;
	value: NumericAttr | EnumAttr | StringAttr;
}

export interface NumericAttr {	
	description?: string;
	high?: number;
	low?: number;
	scalable?: BoolFlag;
	cv?: number; // Current Value
	units?: string;
}
export interface EnumAttr {	
	set?: string;
	stringValue?: string;
	changeable?: BoolFlag;
}
export interface StringAttr {	
	cv?: string;
}
