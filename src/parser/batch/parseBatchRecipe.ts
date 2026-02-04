import { parseMeta } from "../document/parseMeta";
import { BatchRecipe } from "../models/batch/batchRecipe";
import { Blocks, Regex } from "../utils/const";
import { matchNumR, matchStrR, readBlock } from "../utils/utils";
import { parseAttributeInstance, parseFormulaParameter } from "./parseFormulaParameter";
import { parsePFC } from "./pfcAlgorithm";

export function parseRecipe(txtFhx: string[]): BatchRecipe[]{
    
    const recipe = readBlock(Blocks.recipe, txtFhx);
    const batchRecipe = readHeader(recipe[0]);
    const formualParameter = parseFormulaParameter(txtFhx);
    const attributeInstance = parseAttributeInstance(txtFhx);
    const pfcAlgo = parsePFC(txtFhx);

    batchRecipe.formulaParameters = formualParameter || undefined;
    batchRecipe.attributeInstances = attributeInstance || undefined;
    batchRecipe.pfcAlgorithm = pfcAlgo || undefined;    
    return [batchRecipe];
}

function readHeader(recipe: string[]): BatchRecipe {
    const batch: BatchRecipe = {
        meta: {
            user: "",
            time: 0,
            timeStr: "",
        },
        name: "",
        type: "",
        category: "",        
    }

    const bname = matchStrR(Regex.name, recipe[0]);
    if(bname.ok) batch.name = bname.value;

    const btype = matchStrR(Regex.type, recipe[0]);
    if(btype.ok) batch.type = btype.value;

    const bcat = matchStrR(Regex.category, recipe[0]);
    if(bcat.ok) batch.category = bcat.value;
    
    const meta = parseMeta(recipe[1]);
    batch.meta = meta;
    
    for(const line of recipe) {      
        if (line.includes('VERSION')) break;

        const desc = matchStrR(Regex.description, line);
        if(desc.ok) batch.description = desc.value;
        
        const uEquip = matchStrR(Regex.useEquipmentTrains, line);
        if(uEquip.ok) batch.useEquipmentTrains = uEquip.value === 'F' ? true : false;

        const equipUnit = matchStrR(Regex.equipmentUnitClass, line);
        if(equipUnit.ok) batch.equipmentUnitClass = equipUnit.value;
        
        const auth = matchStrR(Regex.author, line);
        if(auth.ok) batch.author = auth.value;
        
        const abstr = matchStrR(Regex.abstract, line);
        if(abstr.ok) batch.abstract = abstr.value;
        
        const bUnits = matchStrR(Regex.batchUnits, line);
        if(bUnits.ok) batch.batchUnits = bUnits.value;

        const bLength = matchStrR(Regex.batchLength, line);
        if(bLength.ok) batch.batchLength = bLength.value;
        
        const bSize = matchNumR(Regex.defaultBatchSize, line);
        if(bSize.ok) batch.defaultBatchSize = bSize.value;

        const minSize = matchNumR(Regex.minimumBatchSize, line);
        if(minSize.ok) batch.minimumBatchSize = minSize.value;
        
       const maxSize = matchNumR(Regex.maximumBatchSize, line);
        if(maxSize.ok) batch.maximumBatchSize = maxSize.value;

        const pCode = matchStrR(Regex.productCode, line);
        if(pCode.ok) batch.productCode = pCode.value;
      
        const pName = matchStrR(Regex.productName, line);
        if(pName.ok) batch.productName = pName.value;
        
        const rAprov = matchStrR(Regex.recipeApproval, line);
        if(rAprov.ok) batch.recipeApprovalInfo = rAprov.value;
        
        const vRecipe = matchStrR(Regex.versionRecipe, line);
        if(vRecipe.ok) batch.version = vRecipe.value;        
       
    }

    return batch;
}

